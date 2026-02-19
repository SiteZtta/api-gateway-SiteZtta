package cfg

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Env        string     `mapstructure:"env"`
	HttpServer HttpServer `mapstructure:"http_server"`
}

type HttpServer struct {
	Address     string        `mapstructure:"address"`
	Port        int           `mapstructure:"port"`
	Timeout     time.Duration `mapstructure:"timeout"`
	IdleTimeout time.Duration `mapstructure:"idle_timeout"`
}

func MustLoad(cname string) Config {
	path := fetchCfgDirPath()
	defer func() {
		if err := recover(); err != nil {
			panic(fmt.Errorf("fatal error config file: %s", err))
		}
	}()
	if path == "" {
		panic(fmt.Errorf("config path is empty"))
	}
	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic(fmt.Errorf("config file not found: %s", path))
	}
	// Setting viper
	viper.AddConfigPath(path)
	viper.SetConfigName(cname)
	viper.SetConfigType("yaml")
	// Env variables
	viper.AutomaticEnv()
	viper.SetEnvPrefix("APP")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_")) // ex: APP_HTTP_SERVER_PORT -> PORT
	// Reading config
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error reading config file: %s", err))
	}
	cfg := Config{}
	if err := viper.Unmarshal(&cfg); err != nil {
		panic(fmt.Errorf("fatal error unmarshaling config: %s", err))
	}
	return cfg
}

func fetchCfgDirPath() string {
	var path string
	// --cfg="./cfg"
	flag.StringVar(&path, "cfg", "", "path to cfg dir")
	flag.Parse()
	if path == "" {
		panic("cfg path is empty")
	}
	return path
}
