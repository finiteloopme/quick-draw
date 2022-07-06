package main

import (
	"net/http"

	log "github.com/finiteloopme/goutils/pkg/log"
	"github.com/kelseyhightower/envconfig"
)

type HTTPConfig struct {
	Host string `default:"0.0.0.0"`
	Port string `default:"8080"`
}

func main() {
	var config HTTPConfig
	envconfig.Process("gcp", &config)
	http.Handle("/", http.FileServer(http.Dir("./")))
	listenAt := config.Host + ":" + config.Port
	log.Info("Server listening at: " + listenAt)
	http.ListenAndServe(listenAt, nil)
}
