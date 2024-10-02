package config

type Server struct {
	HTTP Address
}

type Address struct {
	Address string
}

func GetServerHTTPAddress() string {
	return cfg.Server.HTTP.Address
}