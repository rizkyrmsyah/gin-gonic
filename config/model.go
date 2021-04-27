package config

type Config struct {
	AppPort int
	AppEnv  string
	Debug   bool
	DB      *DB
}

type DB struct {
	Driver   string
	Host     string
	Port     int
	Username string
	Password string
	Name     string
}
