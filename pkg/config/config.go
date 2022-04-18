package config

import (
	"errors"
	"log"

	"github.com/spf13/viper"
)

// Config struct contains configuration parameters.
type Config struct {
	ServerConfig ServerConfig
	JWTConfig    JWTConfig
	DBConfig     DBConfig
}

// ServerConfig struct contains configuration parameters in terms of server.
type ServerConfig struct {
	AppVersion       string
	Mode             string
	Status           string
	RoutePrefix      string
	Debug            bool
	Port             string
	TimeoutSecs      int64
	ReadTimeoutSecs  int64
	WriteTimeoutSecs int64
}

// JWTConfig struct contains configuration parameters in terms of JWT token.
type JWTConfig struct {
	SessionTime int
	SecretKey   string
}

// DBConfig struct contains configuration parameters in terms of database connection.
type DBConfig struct {
	DataSourceName  string
	Name            string
	MigrationFolder string
	MaxOpen         int
	MaxIdle         int
	MaxLifetime     int
}

// LoadConfig takes configuation file path and returns necessary configuation
// parameters as Config struct.
func LoadConfig(filename string) (*Config, error) {
	v := viper.New()

	v.SetConfigName(filename)
	v.AddConfigPath(".")
	v.AutomaticEnv()
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("config file not found")
		}
		return nil, err
	}

	var c Config
	err := v.Unmarshal(&c)
	if err != nil {
		log.Printf("unable to decode into struct, %v", err)
		return nil, err
	}

	return &c, nil

}
