package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"kiss2u/api"
	"kiss2u/cache"
	"kiss2u/config"
	"kiss2u/kiss"

	"github.com/gorilla/mux"
)

func main() {
	parse := config.NewParser()
	opts, err := parse.ParseEnvironmentVariables()
	if err != nil {
		panic(err)
	}

	store := cache.NewStorage()

	lilaclog := kiss.NewLilacLog(store, opts.LilacLog())
	lilacrepo := kiss.NewLilacRepo(store, opts.LilacRepo(), opts.RepoName())

	err = lilaclog.Migrate()
	if err != nil {
		panic(err)
	}
	lilacrepo.Sync()

	fmt.Println("Start Service")

	go lilaclog.WatchJSON()
	if err != nil {
		fmt.Println(err)
	}

	router := mux.NewRouter()
	api.Serve(router, store, &kiss.Kiss{
		LilacLog:  lilaclog,
		LilacRepo: lilacrepo,
	})

	httpServer := &http.Server{
		Addr:         opts.ListenAddr(),
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  60 * time.Second,
		Handler:      router,
	}

	go func() {
		if err := httpServer.ListenAndServe(); err != http.ErrServerClosed {
			fmt.Println(err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	signal.Notify(stop, syscall.SIGTERM)

	<-stop
	fmt.Println("Shutting down the process...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if httpServer != nil {
		httpServer.Shutdown(ctx)
	}

	fmt.Println("Process gracefully stopped")
}
