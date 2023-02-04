package jwt

type Config struct {
	SecretKey string `env:"JWT_SECRET_KEY" envDefault:"secret"`
}
