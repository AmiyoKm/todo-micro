package configs

type Config struct {
	Port int
	Env  string
	Db   DbConfig
	JWT  JWTConfig
}

type DbConfig struct {
	Host        string
	Port        string
	Name        string
	User        string
	Password    string
	SslMode     string
	MaxConnOpen int
	MaxIdleConn int
	MaxIdleTime string
}

type JWTConfig struct {
	Secret          string
	ExpirationHours int
}
