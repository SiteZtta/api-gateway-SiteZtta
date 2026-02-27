package config

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Env        string        `mapstructure:"env"`
	HttpServer HttpServer    `mapstructure:"http_server"`
	Clients    ClientsConfig `mapstructure:"clients"`
	AppSecret  string        `mapstructure:"app_secret"`
}

type HttpServer struct {
	Host        string        `mapstructure:"host"`
	Port        int           `mapstructure:"port"`
	Timeout     time.Duration `mapstructure:"timeout"`
	IdleTimeout time.Duration `mapstructure:"idle_timeout"`
}

type GrpcClient struct {
	Host         string        `mapstructure:"host"`
	Port         int           `mapstructure:"port"`
	Timeout      time.Duration `mapstructure:"timeout"`
	RetriesCount int           `mapstructure:"retries_count"`
	Insecure     bool          `mapstructure:"insecure"`
}

type ClientsConfig struct {
	AuthService GrpcClient `mapstructure:"auth_service"`
}

func MustLoad() Config {
	path := fetchCfgDirPath()
	if path == "" {
		panic("config path is empty")
	}
	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic(fmt.Errorf("config file not found on path %s: %w", path, err))
	}
	// Setting viper
	viper.SetConfigFile(path)
	// Env variables
	viper.AutomaticEnv()
	viper.SetEnvPrefix("APP")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_")) // ex: APP_HTTP_SERVER_PORT -> PORT
	// Reading config
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error reading config file: %w", err))
	}
	cfg := Config{}
	if err := viper.Unmarshal(&cfg); err != nil {
		panic(fmt.Errorf("fatal error unmarshaling config: %w", err))
	}
	return cfg
}

func fetchCfgDirPath() string {
	var path string
	// --cfg="./config/local.yaml"
	flag.StringVar(&path, "cfg", "", "path to cfg dir")
	flag.Parse()
	if path == "" {
		panic("cfg path is empty")
	}
	return path
}
