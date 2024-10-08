package config

import (
	"fmt"
	"github.com/spf13/viper"
	"strings"
)

type DatabaseConfig struct {
	Username string `mapstructure:"DATABASE_USERNAME"`
	Password string `mapstructure:"DATABASE_PASSWORD"`
	Host     string `mapstructure:"DATABASE_HOST"`
	Port     string `mapstructure:"DATABASE_PORT"`
	Database string `mapstructure:"DATABASE_DB"`
	Ssl      string `mapstructure:"DATABASE_SSL"`
}

func getDatabaseConfig() (string, error) {
	v := viper.New()
	var dbConfig DatabaseConfig
	v.SetDefault("DATABASE_PORT", "5432")
	v.SetDefault("DATABASE_SSL", "disable")

	v.BindEnv("DATABASE_USERNAME")
	v.BindEnv("DATABASE_PASSWORD")
	v.BindEnv("DATABASE_HOST")
	v.BindEnv("DATABASE_PORT")
	v.BindEnv("DATABASE_DB")
	v.BindEnv("DATABASE_SSL")

	err := v.Unmarshal(&dbConfig)
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
