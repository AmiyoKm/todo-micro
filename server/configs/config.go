package configs

import "time"

type Config struct {
	Port  int
	Env   string
	Db    DbConfig
	Redis RedisConfig
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

type RedisConfig struct {
	Host       string
	Port       string
	Password   string
	DB         int
	TLS        string
	Expiration time.Duration
}
