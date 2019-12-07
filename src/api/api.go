package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type GetPairData interface {
	GetPair() map[string]string
	Content() string
}

type binance struct {
	Symbols  []string
	url      string
	response *binanceResponse
}

type binanceResponse struct {
	Mins  int    `json:"mins"`
	Price string `json:"price"`
}

func NewBinance(symbols []string) binance {
	binance := binance{}
	binance.Symbols = symbols
	binance.url = "https://www.binance.com/api/v3/avgPrice?symbol=%s"
	binance.response = new(binanceResponse)

	return binance
}

func (b binance) GetPair() map[string]string {
	result := make(map[string]string)
	for _, symbol := range b.Symbols {
		url := fmt.Sprintf(b.url, symbol)
		resp, err := http.Get(url)

		if err != nil {
			fmt.Printf("error response - %s\n", err)
			continue
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			fmt.Printf("error from read body - %s\n", err)
			continue
		}
		error := json.Unmarshal(body, b.response)

		if error != nil {
			fmt.Printf("error from decode body - %s\n", err)
			continue
		}

		result[symbol] = b.response.Price
	}
	return result
}

func (b binance) Content() string {
	var result string
	for k, v := range b.GetPair() {
		result += k + ": " + v + "\n"
	}

	return result
}
