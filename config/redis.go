package config

type RedisMasterNameConfigs struct {
	Name string
}

type RedisConfigs struct {
	URI    string
	Master RedisMasterNameConfigs
}

func GetRedisURI() string {
	if cfg.Redis.URI == "" && GetEnv() == EnvTest {
		return "redis://:123456@localhost:6379"
	}
	return cfg.Redis.URI
}

func GetRedisMasterName() string {
	return cfg.Redis.Master.Name
}