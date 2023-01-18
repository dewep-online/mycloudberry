package firewall

import (
	"context"
	"fmt"
	"net"
	"strconv"
	"strings"

	"github.com/dewep-online/mycloudberry/app/pkg/env"

	"github.com/deweppro/go-utils/shell"
)

type (
	UFW struct {
		rules []Rule
		shell *shell.Shell
	}
)

func newUfw(rules []Rule) *UFW {
	sh := shell.New()
	env.SetupDefaultLang(sh)

	return &UFW{
		shell: sh,
		rules: rules,
	}
}

func (v *UFW) Up() error {
	cmds := make([]string, 0, len(v.rules)+1)
	cmds = append(cmds, "ufw enable", "ufw default deny incoming", "ufw default allow outgoing")
	for _, rule := range v.rules {
		if rule.IsValid() {
			cmds = append(cmds, "ufw "+rule.String())
		}
	}
	return v.shell.CallPackageContext(context.TODO(), cmds...)
}

func (v *UFW) Down() error {
	return nil
}

func (v *UFW) Status() (string, error) {
	b, err := v.shell.Call(context.TODO(), "ufw status verbose")
	fmt.Println(string(b))
	return "", err
}

func (v *UFW) loadStatus() ([]byte, error) {
	return v.shell.Call(context.TODO(), "ufw status verbose")
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

const (
	networkTypeTCP  = "tcp"
	networkTypeUDP  = "udp"
	networkTypeBoth = "any"
)

var networkType = map[string]string{
	"tcp":  networkTypeTCP,
	"tcp4": networkTypeTCP,
	"tcp6": networkTypeTCP,
	"udp":  networkTypeUDP,
	"udp4": networkTypeUDP,
	"udp6": networkTypeUDP,
}

func validNetworkType(v string) string {
	switch v {
	case networkTypeTCP, networkTypeUDP:
		return "/" + v
	default:
		return "/" + networkTypeBoth
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

const (
	accessTypeAllowIn  = "allow in"
	accessTypeAllowOut = "allow out"
	accessTypeDenyIn   = "deny in"
	accessTypeDenyOut  = "deny out"
)

func validAccessType(v string) string {
	v = strings.ToLower(v)
	switch v {
	case "allow in", "allow out",
		"deny in", "deny out":
		return v
	default:
		return accessTypeAllowIn
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type Rule struct {
	ID      int64  `json:"id"`
	Type    string `json:"type"`
	Network string `json:"network"`
	Port    int64  `json:"port"`
	Range   int64  `json:"range"`
	IP      string `json:"ip"`
}

func (v Rule) String() string {
	if !v.IsValid() {
		return ""
	}
	return fmt.Sprintf(
		"%s %s%d%s%s",
		validAccessType(v.Type),
		func() string {
			if len(v.IP) == 0 {
				return ""
			}
			return fmt.Sprintf("from %s to any port ", v.IP)
		}(),
		v.Port,
		func() string {
			if v.Range <= 0 {
				return ""
			}
			return fmt.Sprintf(":%d", v.Range)
		}(),
		validNetworkType(v.Network),
	)
}

func (v Rule) IsValid() bool {
	if v.Port <= 0 {
		return false
	}
	if v.Range < 0 {
		return false
	}
	if len(v.IP) > 0 {
		if strings.Contains(v.IP, "/") {
			if ip, _, err := net.ParseCIDR(v.IP); err != nil || ip.IsLoopback() || ip.IsUnspecified() {
				return false
			}
		} else {
			if ip := net.ParseIP(v.IP); ip == nil || ip.IsLoopback() || ip.IsUnspecified() {
				return false
			}
		}
	}
	return true
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

const (
	prLeftSquareBracket  = byte('[')
	prRightSquareBracket = byte(']')
	prSpace              = byte(' ')
	prNewLine            = byte('\n')
	prColon              = byte(':')
	prSlash              = byte('/')
)

const (
	prStepInit int = iota
	prStepID
	prStepPort
	prStepRange
	prStepNetType
	prStepAction
	prStepIP
)

func ParseRules(v []byte) []Rule {
	result := make([]Rule, 0)
	rule := &Rule{}
	step := prStepInit
	buf := make([]byte, 0, 1000)

	var b byte
	for i := 0; i < len(v); i++ {
		b = v[i]

		switch step {
		case prStepInit:
			if b == prLeftSquareBracket {
				step = prStepID
				buf = buf[:0]
				continue
			}

		case prStepID:
			if b == prSpace {
				continue
			}
			if b == prRightSquareBracket {
				step = prStepPort
				i++
				rule.ID = str2int(string(buf))
				buf = buf[:0]
				continue
			}
			buf = append(buf, b)

		case prStepPort:
			if b == prSpace {
				step = prStepAction
				rule.Port = str2int(string(buf))
				buf = buf[:0]
				continue
			}
			if b == prColon {
				step = prStepRange
				rule.Port = str2int(string(buf))
				buf = buf[:0]
				continue
			}
			if b == prSlash {
				step = prStepNetType
				rule.Port = str2int(string(buf))
				buf = buf[:0]
				continue
			}
			buf = append(buf, b)

		case prStepRange:
			if b == prSpace {
				step = prStepAction
				rule.Range = str2int(string(buf))
				buf = buf[:0]
				continue
			}
			if b == prSlash {
				step = prStepNetType
				rule.Range = str2int(string(buf))
				buf = buf[:0]
				continue
			}
			buf = append(buf, b)

		case prStepNetType:
			if b == prSpace {
				step = prStepAction
				rule.Network = string(buf)
				buf = buf[:0]
				continue
			}
			buf = append(buf, b)

		case prStepAction:
			if b == prSpace && v[i+1] == prSpace && len(buf) > 0 {
				if string(buf) == "(v6)" {
					buf = buf[:0]
					continue
				}
				step = prStepIP
				rule.Type = strings.TrimSpace(string(buf))
				buf = buf[:0]
				continue
			}
			if b == prSpace && v[i+1] == prSpace {
				continue
			}
			buf = append(buf, b)

		case prStepIP:
			if b == prNewLine {
				step = prStepInit
				rule.IP = strings.TrimSpace(string(buf))
				if strings.Contains(rule.IP, "Anywhere") {
					rule.IP = ""
				}
				if rule.IsValid() {
					result = append(result, *rule)
				}
				rule = &Rule{}
				buf = buf[:0]
				continue
			}
			buf = append(buf, b)
		}
	}
	return result
}

func str2int(v string) int64 {
	pp, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		return 0
	}
	return pp
}
