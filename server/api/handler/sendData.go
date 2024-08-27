package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/jsusmachaca/goroapi/internal/util"
	"github.com/jsusmachaca/goroapi/pkg/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

func SendData(w http.ResponseWriter, r *http.Request, client *mongo.Client) {
	w.Header().Set("Content-Type", "application/json")

	token := r.Header.Get("authorization")
	if !strings.HasPrefix(token, "Bearer") {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"error": "token is not provided"}`))
		return
	}

	token = token[7:]

	err := util.VerifyToken(token)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"error": "token is not valid"}`))
		return
	}

	data, err := repository.GetAll(client)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "failed to fetch API data"}`))
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "failed to fetch API data"}`))
		return
	}
}
