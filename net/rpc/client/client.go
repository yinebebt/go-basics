package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type RPCRequest struct {
	Method string      `json:"method"`
	Params interface{} `json:"params"`
	ID     int         `json:"id"`
}

type Param struct {
	ID int `json:"id"`
}

func main() {
	params := []Param{{ID: 1}}
	request := RPCRequest{
		Method: "JSONServer.TwitterProfileDetail",
		Params: params,
		ID:     1,
	}

	reqBody, err := json.Marshal(request)
	if err != nil {
		log.Fatalf("Error marshalling request: %v", err)
	}

	resp, err := http.Post("http://localhost:9000/rpc", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		log.Fatalf("Error sending request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response: %v", err)
	}

	log.Printf("Response: %s", body)
}
