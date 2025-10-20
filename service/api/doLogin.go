package api

import (
	"bubbleLink/service/api/reqcontext"
	"crypto/rand"
	"encoding/json"
	"math/big"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//generate the auth token

func generateAPIKey() (string, error) {
	const apiKeyLength = 16
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var apiKey []byte

	for i := 0; i < apiKeyLength; i++ {
		index, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "error key", err
		}
		apiKey = append(apiKey, charset[index.Int64()])
	}

	return string(apiKey), nil
}

func (rt *_router) doLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	//analyze the request body
	var requestBody struct {
		Name struct { // name has some special conditions in the front end, so we put it inside a struct
			FormatedName string `json:"FormatedName"`
		}
		Country string `json:"Country"`
		City    string `json:"City"`
	}

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request body"})
		return
	}

	country := requestBody.Country
	city := requestBody.City

	// assign the username
	username := requestBody.Name.FormatedName
	if len(username) < 3 || len(username) > 16 {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(map[string]string{"error": "Username must be between 3 and 16 characters"})
		return
	}

	//check if the user is registring or logging in
	exists, err := rt.db.CheckIfUserExists(username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(map[string]string{"error": "Database error"})
		ctx.Logger.WithError(err).Error("Database fail")
		return
	}

	var apiKey string
	var userID int
	if !exists {
		apiKey, err = generateAPIKey() //generate the api key for new users
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_ = json.NewEncoder(w).Encode(map[string]string{"error": "Error generating API key"})
			ctx.Logger.WithError(err).Error("keygen fail")
			return
		}

		userID, err = rt.db.AddNewUser(username, country, city, apiKey)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_ = json.NewEncoder(w).Encode(map[string]string{"error": "Error adding new user"})
			ctx.Logger.WithError(err).Error("database fail")
			return
		}
	} else {
		userID, err = rt.db.GetUserID(username)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_ = json.NewEncoder(w).Encode(map[string]string{"error": "Error retrieving user ID"})
			ctx.Logger.WithError(err).Error("database fail")
			return
		}

		apiKey, err = rt.db.GetUserKey(userID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_ = json.NewEncoder(w).Encode(map[string]string{"error": "Error retrieving API key"})
			ctx.Logger.WithError(err).Error("database fail")

			return
		}
	}

	response := struct {
		UserID int    `json:"userId"`
		APIKey string `json:"apiKey"`
	}{
		UserID: userID,
		APIKey: apiKey,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(response)

}
