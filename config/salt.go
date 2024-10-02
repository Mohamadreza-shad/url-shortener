package config

type Salt struct {
	Key string
}

func SaltKey() string {
	return cfg.Salt.Key
}
