package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type alfabank struct {
	symbols []string
	url     string
}

type rates struct {
	Rates []rate
}

type rate struct {
	BuyIso   string
	SellIso  string
	BuyRate  float32
	SellRate float32
}

func NewAlfaBank(symbols []string) alfabank {
	alfaBank := alfabank{}
	alfaBank.symbols = symbols
	alfaBank.url = "https://developerhub.alfabank.by:8273/partner/1.0.0/public/rates"

	return alfaBank
}

func (a alfabank) GetPair() map[string]string {
	result := make(map[string]string)
	resp, err := http.Get(a.url)

	if err != nil {
		fmt.Printf("error response - %s\n", err)
		return result
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("error response - %s\n", err)
		return result
	}

	var currentRates rates
	error := json.Unmarshal(body, &currentRates)
	if error != nil {
		fmt.Printf("error from decode body - %s\n", err)
		return result
	}

	for _, pair := range currentRates.Rates {
		for _, symbol := range a.symbols {
			sellIso := symbol[:3]
			buyIso := symbol[3:]
			if pair.SellIso == sellIso && pair.BuyIso == buyIso {
				result["USDRUB"] = fmt.Sprintf("Buy - %f,  Sell - %f", pair.BuyRate, pair.SellRate)
			}
		}

	}

	return result
}

func (a alfabank) Content() string {
	var result string
	for k, v := range a.GetPair() {
		result += k + ": " + v + "\n"
	}

	return result
}
