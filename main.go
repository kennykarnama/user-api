package main

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"user-api/config"
	"user-api/domain/api/user"
	"user-api/domain/api/userauth"
	"user-api/domain/repository/user/mysql"
	"user-api/domain/repository/userauth/redis"
	userService "user-api/domain/service/user"
	userAuthService "user-api/domain/service/userauth"
	"user-api/util"
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
		"appname":  cfg.ServiceName,
		"hostname": hostName,
	}

	db := dbconn.InitGorm(cfg.ServiceName)
	redisPool := dbconn.Init(cfg.ServiceName)

	redisWrapper := util.NewRedisWrapper(redisPool)

	userMysqlRepository := mysql.NewMysqlRepository(db)
	userAuthRedisRepository := redis.NewRepository(redisWrapper)

	userService := userService.NewService(userMysqlRepository)
	userAuthService := userAuthService.NewService(cfg, userService, userAuthRedisRepository)

	v := validator.New()
	userHandler := user.NewHandler(ctx, userService, v)
	userAuthHandler := userauth.NewHandler(ctx, v, userAuthService)

	httpServer := &server{
		Server: http.Server{
			Addr: ":" + cfg.RestPort,
		},
	}
	r := mux.NewRouter()
	r.Handle("/api/v1/user", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(userHandler.RegisterUser))).Methods("POST")
	r.Handle("/api/v1/user/auth", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(userAuthHandler.Login))).Methods("POST")
	r.Handle("/api/v1/user/auth/token/validate", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(userAuthHandler.ValidateToken))).Methods("POST")
	r.Handle("/api/v1/user/auth/logout", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(userAuthHandler.Logout))).Methods("POST")
	r.Handle("/api/v1/user/auth/token/refresh", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(userAuthHandler.RefreshToken))).Methods("POST")

	httpServer.Handler = r

	logrus.WithFields(standardFields).Infof("HTTP served on port: %v", cfg.RestPort)

	if err := httpServer.ListenAndServe(); err != nil {
		logrus.WithFields(standardFields).Fatalf("unable to serve. err: %v", err)
	}
}
