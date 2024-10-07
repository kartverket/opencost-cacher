package config

import (
	"fmt"
	"net/url"
	"os"
	"strings"
)

func getOpencostMap() (map[string]string, error) {
	clusters := make(map[string]string)

	for _, env := range os.Environ() {
		parts := strings.SplitN(env, "=", 2)
		key, value := parts[0], parts[1]

		if strings.HasPrefix(key, "OPENCOST_URL_") {
			clusterName := strings.TrimPrefix(key, "OPENCOST_URL_")
			clusterName = strings.ToLower(clusterName)
			trimmedURL := strings.TrimSpace(value)

			if trimmedURL == "" {
				fmt.Printf("Empty URL for cluster '%s'\n", clusterName)
				continue
			}

			parsedURL, err := url.ParseRequestURI(trimmedURL)
			if err != nil {
				return nil, fmt.Errorf("invalid URL '%s' for cluster '%s': %v", trimmedURL, clusterName, err)
			}

			if parsedURL.Scheme != "http" && parsedURL.Scheme != "https" {
				return nil, fmt.Errorf("unsupported URL scheme '%s' in URL '%s' for cluster '%s'", parsedURL.Scheme, trimmedURL, clusterName)
			}

			clusters[clusterName] = trimmedURL
		}
	}

	if len(clusters) == 0 {
		return nil, fmt.Errorf("no OpenCost URLs set")
	}

	return clusters, nil
}
