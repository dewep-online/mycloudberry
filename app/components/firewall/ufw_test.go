package firewall_test

import (
	"fmt"
	"testing"

	"github.com/dewep-online/mycloudberry/app/components/firewall"
)

func TestUnit_RuleString(t *testing.T) {
	tests := []struct {
		fields firewall.Rule
		want   string
	}{
		{fields: firewall.Rule{}, want: ""},
		{fields: firewall.Rule{Port: 80}, want: "allow in 80/any"},
		{fields: firewall.Rule{Type: "allow in", Network: "tcp", Port: 80}, want: "allow in 80/tcp"},
		{fields: firewall.Rule{Type: "allow out", Network: "tcp", Port: 80, Range: 81}, want: "allow out 80:81/tcp"},
		{fields: firewall.Rule{Port: 80, IP: "128.0.0.1"}, want: "allow in from 128.0.0.1 to any port 80/any"},
		{fields: firewall.Rule{Port: 80, IP: "128.0.0.1/24"}, want: "allow in from 128.0.0.1/24 to any port 80/any"},
		{fields: firewall.Rule{Port: 80, IP: "127.0.0.1/124"}, want: ""},
		{fields: firewall.Rule{Port: 80, IP: "127.0.0.1.124"}, want: ""},
		{fields: firewall.Rule{Port: 80, IP: "127.0.0.1"}, want: ""},
		{fields: firewall.Rule{Port: 80, IP: "127.0.0.1/24"}, want: ""},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
			if got := tt.fields.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnit_ParseRules(t *testing.T) {
	data := `
Состояние: активен

     В                          Действие    Из
     -                          --------    --
[ 1] 8080                       ALLOW IN    Anywhere                  
[ 2] 22/tcp                     ALLOW IN    192.168.0.2               
[ 3] 80:90/tcp                  ALLOW IN    192.168.0.2               
[ 4] 80:90/tcp                  ALLOW IN    192.168.0.0/24            
[ 5] 8080 (v6)                  ALLOW IN    Anywhere (v6)             
[ 6] 80:90/tcp                  ALLOW IN    2606:4700:4700::1001 

`
	list := firewall.ParseRules([]byte(data))
	actual := ""
	for _, rule := range list {
		actual += rule.String() + "\n"
	}

	expected := `allow in 8080/any
allow in from 192.168.0.2 to any port 22/tcp
allow in from 192.168.0.2 to any port 80:90/tcp
allow in from 192.168.0.0/24 to any port 80:90/tcp
allow in 8080/any
allow in from 2606:4700:4700::1001 to any port 80:90/tcp
`
	if expected != actual {
		t.Fail()
	}
}
