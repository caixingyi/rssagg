package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetAPIKey(headers http.Header) (string, error) {
	apiKey := headers.Get("Authorization")
	if apiKey == "" {
		return "", errors.New("missing API key")
	}
	values := strings.Split(apiKey, " ")
	if len(values) != 2 || values[0] != "ApiKey" {
		return "", errors.New("invalid API key")
	}
	return values[1], nil
}
