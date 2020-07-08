package main

import (
	"flag"
	"fmt"
	"net/http"

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
		//if err := lilacrepo.GetUsers(); err != nil {
		//	fmt.Println(err)
		//}
		//if err := lilacrepo.GetSubName(); err != nil {
		//	fmt.Println(err)
		//}

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
	http.ListenAndServe(opts.ListenAddr(), router)
}
