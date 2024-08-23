package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/jsusmachaca/goroapi/internal/utils"
	"github.com/jsusmachaca/goroapi/pkg/models"
)

func SendData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	query := r.URL.Query().Get("page")
	token := r.Header.Get("authorization")

	if !strings.HasPrefix(token, "Bearer") {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"error": "token is not provided"}`))
		return
	}

	token = token[7:]

	err := utils.VerifyToken(token)

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"error": "token is not valid"}`))
		return
	}

	page := 1
	if query != "" {
		data, err := strconv.Atoi(query)
		if err != nil {
			w.WriteHeader(http.StatusNotAcceptable)
			w.Write([]byte(`{"error": "page is not valid"}`))
			return
		}
		page = data
	}

	data := utils.GetApiData(page)
	if data == nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "failed to fetch API data"}`))
		return
	}

	var jsonRes models.JsonResponse
	error := json.Unmarshal(data, &jsonRes)
	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message: "han error ocurred"}`))
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(jsonRes.Results)
}
