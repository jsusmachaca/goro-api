package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jsusmachaca/goroapi/utils"
)

func verifyToken(tokenstring string) error {
	key, fileErr := os.ReadFile("jwt.key.pub")

	if fileErr != nil {
		panic(fileErr)
	}
	pubKey, keyErr := jwt.ParseRSAPublicKeyFromPEM(key)

	if keyErr != nil {
		panic(keyErr)
	}

	token, err := jwt.Parse(tokenstring, func(t *jwt.Token) (interface{}, error) {
		return pubKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}

func getApiData(page int) []byte {
	url := fmt.Sprintf("https://rickandmortyapi.com/api/character?page=%d", page)

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	res, err := client.Get(url)
	if err != nil {
		fmt.Println("An error occurred while fetching the API:", err)
		return nil
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Printf("Received non-OK HTTP status: %s\n", res.Status)
		return nil
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("An error occurred while reading the response body:", err)
		return nil
	}

	return data
}

func sendData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	query := r.URL.Query().Get("page")
	token := r.Header.Get("authorization")

	if !strings.HasPrefix(token, "Bearer") {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"error": "token is not provided"}`))
		return
	}

	token = token[7:]

	err := verifyToken(token)
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

	data := getApiData(page)
	if data == nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "failed to fetch API data"}`))
		return
	}

	var jsonRes utils.JsonResponse
	error := json.Unmarshal(data, &jsonRes)
	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message: "han error ocurred"}`))
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(jsonRes.Results)
}

func main() {
	PORT := fmt.Sprintf(":%s", os.Getenv("PORT"))

	http.HandleFunc("/api/RnM", sendData)
	http.ListenAndServe(PORT, nil)
}
