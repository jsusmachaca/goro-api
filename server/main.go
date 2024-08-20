package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	jsonResponse "github.com/jsusmachaca/goroapi/utils"
)

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

	url := fmt.Sprintf("https://rickandmortyapi.com/api/character?page=%d", page)

	res, err := http.Get(url)

	if err != nil {
		fmt.Println("Ocurrió un error")
	}

	defer res.Body.Close()

	data, _ := io.ReadAll(res.Body)

	var jsonRes jsonResponse.JsonResponse

	error := json.Unmarshal(data, &jsonRes)

	if error != nil {
		w.Write([]byte("Hubo un error"))
	}
	json.NewEncoder(w).Encode(jsonRes.Results)
}

func main() {
	http.HandleFunc("/api/RnM", sendData)

	http.ListenAndServe(":8080", nil)
}
