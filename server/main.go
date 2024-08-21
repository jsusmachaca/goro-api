package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	jsonResponse "github.com/jsusmachaca/goroapi/utils"
)

func getApiData(page int) []byte {
	url := fmt.Sprintf("https://rickandmortyapi.com/api/character?page=%d", page)

	res, err := http.Get(url)

	if err != nil {
		fmt.Println("Ocurrió un error")
	}

	defer res.Body.Close()

	data, _ := io.ReadAll(res.Body)

	return data
}

func sendData(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("page")
	page := 1

	if query != "" {
		data, err := strconv.Atoi(query)
		if err != nil {
			panic(err)
		}
		page = data
	}
	data := getApiData(page)

	var jsonRes jsonResponse.JsonResponse

	error := json.Unmarshal(data, &jsonRes)

	if error != nil {
		w.Write([]byte(`{"message: "han error ocurred"}`))
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(jsonRes.Results)
}

func main() {
	PORT := fmt.Sprintf(":%s", os.Getenv("PORT"))

	http.HandleFunc("/api/RnM", sendData)
	http.ListenAndServe(PORT, nil)
}
