package config

type Config struct {
	Port         int
	Host		 string
}

var AppConfig = Config{
	Port:         8080,
	Host:		  "127.0.0.1",
}

func GetHostAndPort() (string, int) {
    return AppConfig.Host, AppConfig.Port
}