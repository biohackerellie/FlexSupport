package config

type Env string

func (e Env) String() string {
	return string(e)
}

const (
	PROD Env = "production"
	DEV  Env = "development"
)

type Config struct {
	LogLevel       string `mapstructure:"LOG_LEVEL"`
	VerboseLogging bool   `mapstructure:"VERBOSE_LOGGING"`
	Environment    Env    `mapstructure:"ENVIRONMENT"`
	Domain         string `mapstructure:"DOMAIN"`
	DatabaseUrl    string `mapstructure:"DATABASE_URL"`
}

func New(getenv func(string, string) string) *Config {
	cfg := &Config{
		LogLevel:       getenv("LOG_LEVEL", "debug"),
		VerboseLogging: getenv("VERBOSE_LOGGING", "false") == "true",
		Environment:    Env(getenv("ENVIRONMENT", "development")),
		Domain:         getenv("DOMAIN", "http://localhost:8080"),
		DatabaseUrl:    getenv("DATABASE_URL", "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"),
	}
	return cfg
}
