package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	demo2()
}

func demo() {
	prdct := product{
		ID:       "1",
		Name:     "product1",
		Price:    3,
		Quantity: 10,
		Status:   true,
	}
	res, err := json.Marshal(prdct)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(res))
}

func demo2() {
	var prdct product
	str := `{
		"id":"1",
		"name":"product1",
		"price":3,
		"quantity":10,
		"status":true
	}`
	json.Unmarshal([]byte(str), &prdct)
	fmt.Println(prdct)
	fmt.Println(prdct.toStr())
}
