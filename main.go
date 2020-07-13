package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"kiss2u/api"
	"kiss2u/config"
	"kiss2u/kiss"
	"kiss2u/storage"

	"github.com/gorilla/mux"
	"github.com/syndtr/goleveldb/leveldb"
)

const (
	flagMigrateHelp = "Run data migrations"
)

func main() {
	var (
		flagMigrate bool
	)

	flag.BoolVar(&flagMigrate, "migrate", false, flagMigrateHelp)
	flag.Parse()

	parse := config.NewParser()
	opts, err := parse.ParseEnvironmentVariables()
	if err != nil {
		panic(err)
	}

	db, err := leveldb.OpenFile(opts.DatabaseURL(), nil)
	if err != nil {
		panic(err)
	}

	store := storage.NewStorage(db)
	lilaclog := kiss.NewLilacLog(store, opts.LilacLog())
	lilacrepo := kiss.NewLilacRepo(store, opts.LilacRepo())

	if flagMigrate {
		err = lilaclog.Migrate()
		if err != nil {
			panic(err)
		}

		lilacrepo.Sync()

		fmt.Println("=== migrate end ===")
		return
	}

	fmt.Println("++++++++++")

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

	db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if httpServer != nil {
		httpServer.Shutdown(ctx)
	}

	fmt.Println("Process gracefully stopped")
}
