package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

var (
	C = new(Config)
)

func unmarshal(rawVal interface{}) bool {
	if err := viper.Unmarshal(rawVal); err != nil {
		fmt.Printf("config file loaded failed. %v\n", err)
		return false
	}
	return true
}

func Load(fpath string) bool {
	_, err := os.Stat(fpath)
	if err != nil {
		return false
	}

	viper.SetConfigFile(fpath)
	viper.SetConfigType("toml")

	err = viper.ReadInConfig()
	if err != nil {
		fmt.Printf("config file %s read failed: %v\n", fpath, err)
		return false
	}

	if !unmarshal(&C) {
		return false
	}
	fmt.Printf("config %s load ok !\n", fpath)
	return true
}

type Config struct {
	Core     Core
	Log      Log
	Token    Token
	Redis    Redis
	Postgres Postgres
}

type Core struct {
	Application string
	Host        string
	Port        string
	Enviroment  string
}

type Log struct {
	Level string
}

type Token struct {
	SecretKey string
}

type Redis struct {
	Address  string
	Password string
	DB       int
}

type Postgres struct {
	Host      string
	Port      int
	User      string
	Password  string
	Database  string
	SSLMode   string
	Migration bool
}

func (c Postgres) DSN() string {
	if os.Getenv("DATABASE_URL") != "" {
		return os.Getenv("DATABASE_URL")
	}
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		c.User, c.Password, c.Host, c.Port, c.Database, c.SSLMode)
}
