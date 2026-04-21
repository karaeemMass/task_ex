package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ApiResponse struct {
	Status string `json:"status"`
	Data   Data   `json:"data"`
}

type Data struct {
	BaseCurrency string                `json:"base_currency"`
	Metals       string                `json:"metals"`
	MetalPrices  map[string]MetalPrice `json:"metal_prices"`
}

type MetalPrice struct {
	Price float64 `json:"price"`
	Open  float64 `json:"open"`
	High  float64 `json:"high"`
	Low   float64 `json:"low"`
}

func GetGoldPrice() {

	url := "https://gold.g.apised.com/v1/latest?metals=XAU%2CXAG%2CXPT%2CXPD&base_currency=KWD&currencies=EUR%2CKWD%2CGBP%2CUSD&weight_unit=gram"
	//url := "https://gold.g.apised.com/v1/latest?metals=XAU%2CXAG%2CXPT%2CXPD&base_currency=EGP&currencies=EUR%2CKWD%2CGBP%2CUSD&weight_unit=gram"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("x-api-key", "sk_8F58a63Df5eeC3E1BD910aEa1673FA112Dd887Ca9E4aFe9B")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)

		return
	}
	fmt.Println(string(body))

	// var result ApiResponse
	// err = json.NewDecoder(res.Body).Decode(&result)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(err)
}
