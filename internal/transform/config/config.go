package config

type Config struct {
	path  string //路径
	token string //token
	ih    string //图床
}

func (c *Config) Path() string {
	return c.path
}

func (c *Config) SetPath(path string) {
	c.path = path
}

func (c *Config) Token() string {
	return c.token
}

func (c *Config) SetToken(token string) {
	c.token = token
}

func (c *Config) Ih() string {
	return c.ih
}

func (c *Config) SetIh(ih string) {
	c.ih = ih
}

func NewConfig(path string, token string, ih string) *Config {
	return &Config{path: path, token: token, ih: ih}
}
