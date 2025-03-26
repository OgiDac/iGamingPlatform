package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/OgiDac/iGamingPlatform/config"
	"github.com/OgiDac/iGamingPlatform/router"
	"github.com/OgiDac/iGamingPlatform/utils"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {

	app := config.App()

	db := app.MySql
	defer app.CloseDatabaseConnection()

	utils.MigrateDB(db)

	timeout := 15 * time.Second
	r := mux.NewRouter()

	router.Setup(timeout, db, r)

	srv := &http.Server{
		Addr:         ":8081",
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
