package config

type Postgres struct {
	URL string
}

func GetPostgresURL() string {
	if cfg.Postgres.URL == "" && GetEnv() == EnvTest {
		return "postgres://postgres:postgres@localhost:5432/url_db"
	}
	return cfg.Postgres.URL
}
