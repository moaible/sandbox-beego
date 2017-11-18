package api

import (
	"github.com/moaible/sandbox-beego/api"
)

type ExampleApi struct {
	Configuration *api.Configuration
}

func newConfiguration() *api.Configuration {
	cfg := &api.Configuration{
		BasePath:      "http://petstore.swagger.io/v2",
		DefaultHeader: make(map[string]string),
		APIKey:        make(map[string]string),
		APIKeyPrefix:  make(map[string]string),
		UserAgent:     "Swagger-Codegen/1.0.0/go",
		Client:     &api.Client{},
	}

	cfg.Client.Config = cfg
	return cfg
}
