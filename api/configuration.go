package api

import (
	"encoding/base64"
	"net/http"
	"time"
)

type Configuration struct {
	Username      string            `json:"userName,omitempty"`
	Password      string            `json:"password,omitempty"`
	APIKeyPrefix  map[string]string `json:"APIKeyPrefix,omitempty"`
	APIKey        map[string]string `json:"APIKey,omitempty"`
	Debug         bool              `json:"debug,omitempty"`
	DebugFile     string            `json:"debugFile,omitempty"`
	OAuthToken    string            `json:"oAuthToken,omitempty"`
	BasePath      string            `json:"basePath,omitempty"`
	Host          string            `json:"host,omitempty"`
	Scheme        string            `json:"scheme,omitempty"`
	AccessToken   string            `json:"accessToken,omitempty"`
	DefaultHeader map[string]string `json:"defaultHeader,omitempty"`
	UserAgent     string            `json:"userAgent,omitempty"`
	Client        *Client
	Transport     *http.Transport
	Timeout       *time.Duration    `json:"timeout,omitempty"`
}

func (c *Configuration) GetBasicAuthEncodedString() string {
	return base64.StdEncoding.EncodeToString([]byte(c.Username + ":" + c.Password))
}

func (c *Configuration) AddDefaultHeader(key string, value string) {
	c.DefaultHeader[key] = value
}

func (c *Configuration) GetAPIKeyWithPrefix(APIKeyIdentifier string) string {
	if c.APIKeyPrefix[APIKeyIdentifier] != "" {
		return c.APIKeyPrefix[APIKeyIdentifier] + " " + c.APIKey[APIKeyIdentifier]
	}
	return c.APIKey[APIKeyIdentifier]
}
