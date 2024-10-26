package http_tester

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type HTTPClient struct {
	instances int
	duration  int
	port      string
}

type Message struct {
	Message string `json:"message"`
}

func NewClient(instances int, duration int, port string) *HTTPClient {
	return &HTTPClient{
		instances: instances,
		duration:  duration,
		port:      port,
	}
}

func (h *HTTPClient) Run() {
	address := "http://localhost"
	httpClient := &http.Client{}
	resp, err := httpClient.Get(address + h.port)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var result Message
	json.Unmarshal(body, &result)

	log.Println(result)
}
