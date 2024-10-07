package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	OpenCostURLs   map[string]string
	FullSync       bool
	DatabaseConfig string
	LocalDB        bool
}

func LoadConfig() (*Config, error) {
	viper.BindEnv("FULL_SYNC")
	fullSync := viper.GetBool("FULL_SYNC")

	if fullSync {
		fmt.Println("Full sync enabled, upserting everything")
	}

	viper.BindEnv("LOCALDB")
	localdb := viper.GetBool("LOCALDB")

	openCostUrls, err := getOpencostMap()
	if err != nil {
		return nil, fmt.Errorf("error loading OpenCost URLs: %v", err)
	}

	dsn := ""
	if !localdb {
		dsn, err = getDatabaseConfig()
		if err != nil {
			return nil, fmt.Errorf("error loading database config: %v", err)
		}
	}

	config := &Config{
		OpenCostURLs:   openCostUrls,
		FullSync:       fullSync,
		DatabaseConfig: dsn,
		LocalDB:        localdb,
	}

	return config, nil
}
