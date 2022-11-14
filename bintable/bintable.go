package bintable

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
)

type LookupResponse_t struct {
	Result  int                  `json:"result"`
	Message string               `json:"message"`
	Data    LookupDataResponse_t `json:"data"`
}

type LookupDataResponse_t struct {
	Card    LookupDataCardResponse_t    `json:"card"`
	Country LookupDataCountryResponse_t `json:"country"`
	Bank    LookupDataBankResponse_t    `json:"bank"`
}

type LookupDataCardResponse_t struct {
	Scheme    string `json:"scheme"`
	Type      string `json:"type"`
	Category  string `json:"category"`
	Length    int    `json:"length"`
	CheckLuhn int    `json:"checluhn"`
	CVVlen    int    `json:"cvvlen"`
}

type LookupDataCountryResponse_t struct {
	Name         string `json:"name"`
	Code         string `json:"code"`
	Flag         string `json:"flag"`
	Currency     string `json:"currency"`
	CurrencyCode string `json:"currency_code"`
}

type LookupDataBankResponse_t struct {
	Name    string `json:"name"`
	Website string `json:"website"`
	Phone   string `json:"phone"`
}

type Bintable interface {
	New(key string) *BintableConfig
}

type BintableConfig struct {
	Key      string
	Endpoint string
}

func New(key string) *BintableConfig {
	if key == "" {
		log.Fatalln("You must set api key")
	}

	return &BintableConfig{
		Key:      key,
		Endpoint: "https://api.bintable.com/v1/",
	}
}

// no working because the endpoint /balance not working in bintable

func (bc *BintableConfig) GetBalance() (map[string]interface{}, error) {
	if bc.Key == "" {
		return nil, errors.New("You must set api key")
	}

	body := Get(fmt.Sprintf("%vbalance?api_key=%v", bc.Endpoint, bc.Key))

	fmt.Println(string(body))

	return nil, nil
}

func (bc *BintableConfig) Lookup(bin string) (LookupResponse_t, error) {
	if bc.Key == "" {
		return LookupResponse_t{}, errors.New("You must set api key")
	}

	var lookup LookupResponse_t

	body := Get(fmt.Sprintf("%v%v?api_key=%v", bc.Endpoint, bin, bc.Key))

	if err := json.Unmarshal(body, &lookup); err != nil {
		return LookupResponse_t{}, err
	}

	if lookup.Result == 404 {
		return LookupResponse_t{}, errors.New("Metadata not found")
	}

	return lookup, nil

}
