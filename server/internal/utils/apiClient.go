package utils

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func GetApiData(page int) []byte {
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
