package hub

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewOptionsFormNew(t *testing.T) {
	testEnv := map[string]string{
		"ACME_CERT_DIR":           "/tmp",
		"ACME_HOSTS":              "example.com,example.org",
		"ADDR":                    "127.0.0.1:8080",
		"ALLOW_ANONYMOUS":         "1",
		"CERT_FILE":               "foo",
		"CORS_ALLOWED_ORIGINS":    "*",
		"DEBUG":                   "1",
		"DEMO":                    "1",
		"KEY_FILE":                "bar",
		"PUBLISHER_JWT_KEY":       "foo",
		"SUBSCRIBER_JWT_KEY":      "bar",
		"PUBLISH_ALLOWED_ORIGINS": "http://127.0.0.1:8080",
	}
	for k, v := range testEnv {
		os.Setenv(k, v)
		defer os.Unsetenv(k)
	}

	options, err := NewOptionsFromEnv()
	assert.Equal(t, &Options{
		true,
		[]byte("foo"),
		[]byte("bar"),
		true,
		[]string{"*"},
		[]string{"http://127.0.0.1:8080"},
		"127.0.0.1:8080",
		[]string{"example.com", "example.org"},
		"/tmp",
		"foo",
		"bar",
		true,
	}, options)
	assert.Nil(t, err)
}

func TestMissingEnv(t *testing.T) {
	_, err := NewOptionsFromEnv()
	assert.EqualError(t, err, "The following environment variable must be defined: [PUBLISHER_JWT_KEY SUBSCRIBER_JWT_KEY]")
}
