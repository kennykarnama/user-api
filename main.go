package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"user-api/config"
	"github.com/sirupsen/logrus"
)

type server struct {
	http.Server
}

func main() {
	cfg := config.Get()

	logrus.SetFormatter(&logrus.JSONFormatter{})

	hostName, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	standardFields := logrus.Fields{
		"appname": "user-api",
		"hostname": hostName,
	}

	logrus.WithFields(standardFields).Infof("HTTP served on port: %v", cfg.RestPort)
	httpServer := NewServer(cfg)

	if err := httpServer.ListenAndServe(); err != nil {
		logrus.WithFields(standardFields).Fatalf("unable to serve. err: %v", err)
	}
}

func NewServer(cfg config.Config) *server {
	s := &server{
		Server: http.Server{
			Addr: ":" + cfg.RestPort,
		},
	}
	r := mux.NewRouter()
	s.Handler = r
	return s
}