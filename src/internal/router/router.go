package router

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

type RedirectRouter struct {
	RouteConfig RouteConfig
}

func NewRedirectRouter() (*RedirectRouter, error) {
	workingDirectory, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	if strings.HasSuffix(workingDirectory, "cmd") {
		workingDirectory = filepath.Join(workingDirectory, "..")
	}
	configPath := filepath.Join(workingDirectory, "config", "url-shortner.yaml")
	fileContent, err := os.ReadFile(configPath)

	if err != nil {
		return nil, err
	}
	routes := []Route{}
	err = yaml.Unmarshal(fileContent, &routes)

	routerConfig := func(routes []Route) RouteConfig {
		redirects := make(map[string]Route)
		for _, route := range routes {
			redirects[route.Path] = route
		}
		return RouteConfig{Redirects: redirects}
	}(routes)

	if err != nil {
		return nil, err
	}

	return &RedirectRouter{RouteConfig: routerConfig}, nil
}

func (redirectRouter *RedirectRouter) HandleRedirects(mux *http.ServeMux) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if route, ok := redirectRouter.RouteConfig.Redirects[r.URL.Path]; ok {
			http.Redirect(w, r, route.URL, http.StatusFound)
			return
		}
		mux.ServeHTTP(w, r)
	}
}

type RouteConfig struct {
	Redirects map[string]Route `yaml:"redirects"`
}

type Route struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}
