package main

import (
	env "github.com/godofprodev/rss-aggregator/internal/config"
	"github.com/godofprodev/rss-aggregator/internal/router"
	_ "github.com/lib/pq"
)

func main() {
	apiRouter := router.NewRouter()

	err := env.Load()
	if err != nil {
		panic(err)
	}

	apiRouter.Start()
}
