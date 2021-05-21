package config

import "time"

const (
	// pagination
	PAGINATION_LIMIT = 100
	PAGINATION_SORT  = "created_at asc"
	// pagination
)

type (
	Config struct {
		HTTP HTTPConfig
	}
	HTTPConfig struct {
		Host               string `mapstructure:"host"`
		Port               string `mapstructure:"port"`
		MaxHeaderMegabytes int    `mapstructure:"maxHeaderBytes"`
		ReadTimeout  time.Duration `mapstructure:"readTimeout"`
		WriteTimeout time.Duration `mapstructure:"writeTimeout"`
	}
)