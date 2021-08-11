package main

import (
	"context"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"user-api/config"
	"user-api/domain/api/user"
	"user-api/domain/repository/user/mysql"
	userService "user-api/domain/service/user"
	"user-api/util/dbconn"
)

type server struct {
	http.Server
}

func main() {

	ctx := context.Background()
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

	fmt.Println(cfg.ServiceName)
	db := dbconn.InitGorm(cfg.ServiceName)
	userMysqlRepository := mysql.NewMysqlRepository(db)
	userService := userService.NewService(userMysqlRepository)

	v := validator.New()
	userHandler := user.NewHandler(ctx, userService, v)

	logrus.WithFields(standardFields).Infof("HTTP served on port: %v", cfg.RestPort)
	httpServer := &server{
		Server: http.Server{
			Addr: ":" + cfg.RestPort,
		},
	}
	r := mux.NewRouter()
	r.Handle("/api/v1/user", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(userHandler.RegisterUser))).Methods("POST")

	httpServer.Handler = r

	if err := httpServer.ListenAndServe(); err != nil {
		logrus.WithFields(standardFields).Fatalf("unable to serve. err: %v", err)
	}
}
