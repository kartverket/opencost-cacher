package opencost

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func getURL(baseUrl string, window string, reportType string) (string, error) {
	return fmt.Sprintf("%s/model/allocation/compute?window=%s&aggregate=%s&includeIdle=true", baseUrl, window, reportType), nil
}

func GetReport(baseUrl string, window string, reportType string) (*Response, error) {
	url, err := getURL(baseUrl, window, reportType)
	if err != nil {
		return nil, err
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, err
	}

	var response Response
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}

	return &response, nil
}
