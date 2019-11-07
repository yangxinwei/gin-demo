package config

type Config struct {
}

func Get() (*Config,error) {

	return &Config{},nil
}
