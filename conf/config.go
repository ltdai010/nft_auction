package conf

type AppConfig struct {
	Port string `env:"PORT" envDefault:"8080"`
	//DB CONFIG
	DBHost   string `env:"DB_HOST" envDefault:"dbmasternode.stg.int.finan.cc"`
	DBPort   string `env:"DB_PORT" envDefault:"5432"`
	DBUser   string `env:"DB_USER" envDefault:"finan_dev_user"`
	DBPass   string `env:"DB_PASS" envDefault:"Oo5Tah0re5eexoif"`
	DBName   string `env:"DB_NAME" envDefault:"finan_dev_ms_notification_management"`
	DBSchema string `env:"DB_SCHEMA" envDefault:"public"`
	EnableDB string `env:"ENABLE_DB" envDefault:"true"`
	Version  string `env:"VERSION" envDefault:"v0.0.1"`
}
