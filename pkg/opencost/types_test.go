package opencost

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestLoadJSON(t *testing.T) {
	data, err := os.ReadFile("../../test_files/opencost_response.json")
	if err != nil {
		t.Fatalf("Failed to read test_data.json: %v", err)
	}

	var response Response
	err = json.Unmarshal(data, &response)

	assert.Nil(t, err)
	assert.Equal(t, 4, len(response.Data[0]))
}
