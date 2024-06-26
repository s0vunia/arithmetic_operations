package config

import (
	"github.com/spf13/viper"
	"log"
	"time"
)

type Config struct {
	HTTPServer  HTTPServer  `mapstructure:"http_server"`
	Storage     Storage     `mapstructure:"storage"`
	Operation   Operation   `mapstructure:"operation"`
	Agent       Agent       `mapstructure:"agents"`
	AuthService AuthService `mapstructure:"auth_service"`
}

type HTTPServer struct {
	Address     string        `mapstructure:"address"`
	Timeout     time.Duration `mapstructure:"timeout"`
	IdleTimeout time.Duration `mapstructure:"idle_timeout"`
}

type Storage struct {
	URL string `mapstructure:"url"`
}

type Operation struct {
	DurationInMilliSecondAddition       int `mapstructure:"addition"`
	DurationInMilliSecondSubtraction    int `mapstructure:"subtraction"`
	DurationInMilliSecondMultiplication int `mapstructure:"multiplication"`
	DurationInMilliSecondDivision       int `mapstructure:"division"`
	CountOperation                      int `mapstructure:"count_operation"`
}

type Agent struct {
	CountOfAgents int `mapstructure:"count_of_agents"`
}

type AuthService struct {
	Cost     int           `mapstructure:"cost"`
	Secret   string        `mapstructure:"secret"`
	TokenTTL time.Duration `mapstructure:"token_ttl"`
}

func Load() *Config {
	v := viper.New()

	v.AddConfigPath(".")
	v.SetConfigName("config")
	v.SetConfigType("json")

	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("couldn't load config: %s", err)
	}

	var cfg Config

	if err := v.Unmarshal(&cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	checkDuration(&cfg)

	return &cfg
}

func checkDuration(cfg *Config) {
	if cfg.Operation.DurationInMilliSecondAddition < 0 {
		log.Fatalf("duration of addition operation is lower than 0")
	}

	if cfg.Operation.DurationInMilliSecondSubtraction < 0 {
		log.Fatalf("duration of subtraction operation is lower than 0")
	}

	if cfg.Operation.DurationInMilliSecondMultiplication < 0 {
		log.Fatalf("duration of multiplication operation is lower than 0")
	}

	if cfg.Operation.DurationInMilliSecondDivision < 0 {
		log.Fatalf("duration of division operation is lower than 0")
	}
}
