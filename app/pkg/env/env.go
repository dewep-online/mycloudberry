package env

type EnvironSetter interface {
	SetEnv(key, value string)
}

func SetupDefaultLang(e EnvironSetter) {
	e.SetEnv("LANG", "en_US.UTF-8")
	e.SetEnv("LANGUAGE", "en_US")
}
