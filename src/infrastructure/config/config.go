package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	ServiceName   string `envconfig:"SERVICE_NAME" default:"dock-rest-api"`
	Env           string `envconfig:"ENV" default:"development"`
	Port          int    `envconfig:"PORT" default:"8080" required:"true"`
	DBType        string `envconfig:"DB_TYPE" default:"postgres"`
	DSN           string `envconfig:"DSN" default:"dbname=dock sslmode=disable user=admin password=admin host=localhost"`
	DBTypeTest    string `envconfig:"DB_TYPE_TEST" default:"sqlite3"`
	DSNTest       string `envconfig:"DSN_TEST" default:":memory:"`
	AutoMigrateDB bool   `envconfig:"AUTO_MIGRATE_DB" default:"true"`
	Debug         bool   `envconfig:"DEBUG" default:"true"`
}

func New() Config {
	cfg := Config{}
	envconfig.MustProcess("", &cfg)
	return cfg
}
