package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/jsusmachaca/goroapi/pkg/models"
)

func GetApiData() ([]models.ResultsResponse, error) {
	allData := make([]models.ResultsResponse, 0, 400)

	for pages := 1; pages < 21; pages++ {
		url := fmt.Sprintf("https://rickandmortyapi.com/api/character?page=%d", pages)

		client := &http.Client{}
		res, err := client.Get(url)
		if err != nil {
			fmt.Println("An error occurred while fetching the API:", err)
			return nil, err
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			fmt.Printf("Received non-OK HTTP status: %s\n", res.Status)
			return nil, err
		}

		data, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Println("An error occurred while reading the response body:", err)
			return nil, err
		}

		var jsonRes models.JsonResponse
		marshErr := json.Unmarshal(data, &jsonRes)
		if marshErr != nil {
			fmt.Println("An error occurred while marshall data:", err)
			return nil, marshErr
		}

		allData = append(allData, jsonRes.Results...)
	}

	return allData, nil
}
