package config

import (
	"encoding/json"
	"fmt"
	"github.com/godofprodev/rss-aggregator/internal/database"
	"github.com/godofprodev/rss-aggregator/internal/utils"
	"github.com/google/uuid"
	"net/http"
	"sync"
	"time"
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

func (apiCfg *ApiConfig) HandleCreateUser(w http.ResponseWriter, r *http.Request) {

	type parameters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("Error parsing JSON:", err))
		return
	}
	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("Couldn't create user:", err))
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, user)
}