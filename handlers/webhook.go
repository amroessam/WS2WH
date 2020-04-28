package handlers

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

//Message type
type Message struct {
	Data    []byte
	Address string
}

//WebhookHandler handles webhooks
func WebhookHandler(m Message) {

	requestBody, err := json.Marshal(map[string]string{
		"text": string(m.Data),
	})

	if err != nil {
		log.Fatalln(err)
	}

	resp, err := http.Post(m.Address, "application/json", bytes.NewBuffer(requestBody))

	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

}
