package main

import (
	"encoding/json"
	"io"
	"net/http"

	"go.uber.org/zap"
)

// Colors is response of api
type Colors struct {
	Colors []struct {
		Value string `json:"value"`
	} `json:"colors"`
}

func newColors(body io.Reader) (Colors, error) {
	var colors Colors
	err := json.NewDecoder(body).Decode(&colors)
	return colors, err
}

// Get is function of example of http.Get
func Get(url string) {

	logger, _ := zap.NewDevelopment()
	resp, err := http.Get(url)
	_ = err
	defer resp.Body.Close()

	colors, err := newColors(resp.Body)
	_ = err

	logger.Info("get colors", zap.Reflect("colors", colors))
}

func main() {
	url := "https://api.noopschallenge.com/hexbot"
	Get(url)
}
