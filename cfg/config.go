package cfg

type Config struct {
	Env        string     `mapstructure:"env"`
	HttpServer HttpServer `mapstructure:"http_server"`
}

type HttpServer struct {
	Port        int    `mapstructure:"port"`
	Timeout     string `mapstructure:"timeout"`
	IdleTimeout string `mapstructure:"idle_timeout"`
}

func MustLoad() Config {
	cfg := Config{}
	return cfg
}
