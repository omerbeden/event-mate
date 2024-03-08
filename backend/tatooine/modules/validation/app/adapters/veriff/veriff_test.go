package veriff

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVeriff_GenerateSession(t *testing.T) {

}

func TestNewVeriffAdapterWithValidCredentials(t *testing.T) {
	// Mock the environment variables
	os.Setenv("VERIFF_API_KEY", "validApiKey")
	os.Setenv("VERIFF_SECRET_KEY", "validSecretKey")
	os.Setenv("VERIFF_BASE_URL", "validBaseUrl")

	// Call the function under test
	adapter := NewVeriffAdapter()

	// Assert that the returned instance has the correct credentials
	assert.Equal(t, "validApiKey", adapter.creds.ApiKey)
	assert.Equal(t, "validSecretKey", adapter.creds.SecretKey)
	assert.Equal(t, "validBaseUrl", adapter.creds.BaseUrl)
}
