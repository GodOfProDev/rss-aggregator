package router

import (
	"github.com/godofprodev/rss-aggregator/internal/utils"
	"net/http"
)

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithJSON(w, 200, struct{}{})
}

func handlerErr(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithError(w, 400, "Something went wrong")
}
