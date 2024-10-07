package config

import (
	"fmt"
	"github.com/spf13/viper"
	"strings"
)

type DatabaseConfig struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Database string `mapstructure:"database"`
	Ssl      string `mapstructure:"ssl"`
}

func getDatabaseConfig() (string, error) {
	v := viper.New()

	v.SetEnvPrefix("DATABASE")

	v.SetDefault("port", "5432")
	v.SetDefault("ssl", "disable")

	v.AutomaticEnv()

	var dbConfig DatabaseConfig

	err := viper.Unmarshal(&dbConfig)
	if err != nil {
		return "", fmt.Errorf("unable to decode configuration into struct: %v", err)
	}

	missingFields := []string{}
	if dbConfig.Host == "" {
		missingFields = append(missingFields, "DATABASE_HOST")
	}
	if dbConfig.Username == "" {
		missingFields = append(missingFields, "DATABASE_USERNAME")
	}
	if dbConfig.Password == "" {
		missingFields = append(missingFields, "DATABASE_PASSWORD")
	}
	if dbConfig.Database == "" {
		missingFields = append(missingFields, "DATABASE_DATABASE")
	}

	if len(missingFields) > 0 {
		return "", fmt.Errorf("missing required environment variables: %s", strings.Join(missingFields, ", "))
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		dbConfig.Host, dbConfig.Username, dbConfig.Password, dbConfig.Database, dbConfig.Port, dbConfig.Ssl)

	return dsn, nil
}
