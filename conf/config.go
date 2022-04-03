package conf

import "github.com/caarlos0/env/v6"

type AppConfig struct {
	Port string `env:"PORT" envDefault:"8080"`
	//DB CONFIG
	DBHost   string `env:"DB_HOST" envDefault:"112.137.129.247"`
	DBPort   string `env:"DB_PORT" envDefault:"10032"`
	DBUser   string `env:"DB_USER" envDefault:"docker"`
	DBPass   string `env:"DB_PASS" envDefault:"docker"`
	DBName   string `env:"DB_NAME" envDefault:"auction"`
	DBSchema string `env:"DB_SCHEMA" envDefault:"public"`
	EnableDB string `env:"ENABLE_DB" envDefault:"true"`
	Version  string `env:"VERSION" envDefault:"v0.0.1"`
}

var config AppConfig

func SetEnv() {
	_ = env.Parse(&config)
}

func LoadEnv() AppConfig {
	return config
}

var Message = `{
  "publisher": "NFT_Auction",
  "onlySignOn": "https://uchain.duckdns.org"
}`
