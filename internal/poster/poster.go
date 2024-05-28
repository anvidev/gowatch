package poster

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func PostData(data string) {
	endpoint := "http://localhost:3000/post"

	payload := map[string]string{"content": data}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post(endpoint, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("poster", "response from server:", string(responseBody))
}
