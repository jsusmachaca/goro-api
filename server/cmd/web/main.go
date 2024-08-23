package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/jsusmachaca/goroapi/api/handler"
)

func main() {
	PORT := fmt.Sprintf(":%s", os.Getenv("PORT"))

	http.HandleFunc("/api/rnm", handler.SendData)
	http.ListenAndServe(PORT, nil)
}
