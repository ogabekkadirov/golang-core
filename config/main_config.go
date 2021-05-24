package config

// import (
// 	"strings"
// 	"time"

// 	"github.com/spf13/viper"
// )

const (
	// pagination
	PAGINATION_LIMIT = 100
	PAGINATION_SORT  = "created_at asc"
	// pagination
)

// type (
// 	Config struct {
// 		HTTP HTTPConfig
// 	}
// 	HTTPConfig struct {
// 		Host               string `mapstructure:"host"`
// 		Port               string `mapstructure:"port"`
// 		MaxHeaderMegabytes int    `mapstructure:"maxHeaderBytes"`
// 		ReadTimeout  time.Duration `mapstructure:"readTimeout"`
// 		WriteTimeout time.Duration `mapstructure:"writeTimeout"`
// 	}
// )

// func Init(path string)(*Config,error){

// 	populateDefaults() // set default settings

// 	if err := parseConfigFile(path); err != nil {
// 		return nil, err
// 	}

// 	var cfg Config

// 	if err := unmarshal(&cfg); err != nil {
// 		return nil, err
// 	}

// 	return &cfg, nil;

// }

// func unmarshal(cfg *Config) error{

// 	if err := viper.UnmarshalKey("http", &cfg.HTTP); err != nil {
// 		return err
// 	}
// 	return nil
// }

// func populateDefaults(){
// }

// func parseConfigFile(filepath string) error{

// 	path := strings.Split(filepath, "/")

// 	viper.AddConfigPath(path[0]) // folder
// 	viper.SetConfigName(path[1]) // config file name

// 	return viper.ReadInConfig()
// }