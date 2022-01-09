package config

type MainConfig struct {
	UserName string
	UserKey  string
	Host     struct {
		HostName string
		HostPort string
	}
	DataBase struct {
		Host string
		User string
		Pass string
		Data string
	}
	SecKey      string
	FrontURL    string
	StaticPath  string
	StaticPath2 string
	ACAO        string
	ACAH        string
	LogFile     string
}

var AppMainConfig MainConfig