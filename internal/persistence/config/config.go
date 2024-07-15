package config

type Config struct {
	PostgresDB       string `env:"POSTGRES_DB,required=true"`
	PostgresHost     string `env:"POSTGRES_HOST,required=true"`
	PostgresPort     string `env:"POSTGRES_PORT,required=true"`
	PostgresUser     string `env:"POSTGRES_USER,required=true"`
	PostgresPassword string `env:"POSTGRES_PASSWORD,required=true"`
}
