package config

import (
	"github.com/godofprodev/rss-aggregator/internal/database"
	"github.com/godofprodev/rss-aggregator/internal/utils"
	"net/http"
	"sync"
)

type ApiConfig struct {
	DB *database.Queries
}

var apiInstance *ApiConfig
var apiOnce sync.Once

func Init(queries *database.Queries) {
	apiOnce.Do(func() {
		apiInstance = &ApiConfig{DB: queries}
	})
}

func GetApiConfig() *ApiConfig {
	return apiInstance
}

func (api *ApiConfig) HandleCreateUser(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithJSON(w, http.StatusOK, struct {
		Username string `json:"username"`
	}{Username: "GodOfPro"})
}
