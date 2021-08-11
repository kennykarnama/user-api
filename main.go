package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"user-api/config"
	"user-api/domain/api/user"
)

type server struct {
	http.Server
}

func main() {
	cfg := config.Get()

	hostName, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	logrus.SetFormatter(&logrus.JSONFormatter{})
	standardFields := logrus.Fields{
		"appname":  "user-api",
		"hostname": hostName,
	}

	v := validator.New()
	userHandler := user.NewHandler(nil, v)

	logrus.WithFields(standardFields).Infof("HTTP served on port: %v", cfg.RestPort)
	httpServer := &server{
		Server: http.Server{
			Addr: ":" + cfg.RestPort,
		},
	}
	r := mux.NewRouter()
	r.Handle("/api/v1/user", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(userHandler.RegisterUser)))

	httpServer.Handler = r

	if err := httpServer.ListenAndServe(); err != nil {
		logrus.WithFields(standardFields).Fatalf("unable to serve. err: %v", err)
	}
}
