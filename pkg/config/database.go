package config

import (
	"fmt"
	"github.com/spf13/viper"
	"strings"
)

type DatabaseConfig struct {
	Username   string `mapstructure:"DATABASE_USERNAME"`
	Password   string `mapstructure:"DATABASE_PASSWORD"`
	Host       string `mapstructure:"DATABASE_HOST"`
	Port       string `mapstructure:"DATABASE_PORT"`
	Database   string `mapstructure:"DATABASE_DB"`
	Ssl        string `mapstructure:"DATABASE_SSL"`
	CaCert     string `mapstructure:"DATABASE_CA_CERT_PATH"`
	ClientCert string `mapstructure:"DATABASE_CLIENT_CERT_PATH"`
	ClientKey  string `mapstructure:"DATABASE_CLIENT_KEY_PATH"`
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
	v.BindEnv("DATABASE_CA_CERT_PATH")
	v.BindEnv("DATABASE_CLIENT_CERT_PATH")
	v.BindEnv("DATABASE_CLIENT_KEY_PATH")

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

	params := []string{
		fmt.Sprintf("host=%s", dbConfig.Host),
		fmt.Sprintf("user=%s", dbConfig.Username),
		fmt.Sprintf("password=%s", dbConfig.Password),
		fmt.Sprintf("dbname=%s", dbConfig.Database),
		fmt.Sprintf("port=%s", dbConfig.Port),
		fmt.Sprintf("sslmode=%s", dbConfig.Ssl),
	}

	if dbConfig.CaCert != "" {
		params = append(params, fmt.Sprintf("sslrootcert=%s", dbConfig.CaCert))
	}
	if dbConfig.ClientCert != "" {
		params = append(params, fmt.Sprintf("sslcert=%s", dbConfig.ClientCert))
	}
	if dbConfig.ClientKey != "" {
		params = append(params, fmt.Sprintf("sslkey=%s", dbConfig.ClientKey))
	}

	dsn := strings.Join(params, " ")

	return dsn, nil
}
