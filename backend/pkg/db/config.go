package db

type Config struct {
	Addr     string `env:"DATABASE_ADDR" envDefault:"localhost:5432"`
	User     string `env:"DATABASE_USER" envDefault:"backend"`
	Password string `env:"DATABASE_PASSWORD" envDefault:"password"`
	Database string `env:"DATABASE_NAME" envDefault:"backend"`
}
