package main

import (
	"database/sql"
	"github.com/godofprodev/rss-aggregator/internal/api"
	env "github.com/godofprodev/rss-aggregator/internal/config"
	"github.com/godofprodev/rss-aggregator/internal/database"
	"github.com/godofprodev/rss-aggregator/internal/router"
	_ "github.com/lib/pq"
)

func main() {

	err := env.Load()
	if err != nil {
		panic(err)
	}

	dbUrl := env.GetEnv().DbUrl

	conn, err := sql.Open("postgres", dbUrl)
	if err != nil {
		panic(err)
	}

	api.Init(database.New(conn))

	apiRouter := router.NewRouter()

	err = env.Load()
	if err != nil {
		panic(err)
	}

	apiRouter.Start()
}
