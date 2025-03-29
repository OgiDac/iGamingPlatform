package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	_ "github.com/OgiDac/iGamingPlatform/app/docs"
	"github.com/OgiDac/iGamingPlatform/config"
	"github.com/OgiDac/iGamingPlatform/router"
	"github.com/OgiDac/iGamingPlatform/utils"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title           iGaming Platform API
// @version         1.0
// @description     API documentation for the iGaming Platform
// @host            localhost:8081
// @BasePath        /
func main() {

	app := config.App()
	env := app.Env
	db := app.MySql
	defer app.CloseDatabaseConnection()

	utils.MigrateDB(db)

	timeout := time.Duration(env.ContextTimeout) * time.Second
	r := mux.NewRouter()

	router.Setup(env, timeout, db, r)

	// r.HandleFunc("/health", router.HealthCheck)
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	srv := &http.Server{
		Addr:         env.ServerAddress,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			fmt.Println(err)
		}
	}()
	fmt.Println("Server started")

	// Shutdown

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	srv.Shutdown(ctx)
	fmt.Println("shutting down")
	os.Exit(0)
}
