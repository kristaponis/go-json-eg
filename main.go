package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	demo4()
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
	fmt.Println(prdct.toStr())
}

func demo3() {
	prdct := product{
		ID:       "1",
		Name:     "product1",
		Price:    3,
		Quantity: 10,
		Status:   true,
		Category: category{
			ID:   "101",
			Name: "category1",
		},
	}
	res, err := json.Marshal(prdct)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(res))
}

func demo4() {
	prdct := product{
		ID:       "1",
		Name:     "product1",
		Price:    3,
		Quantity: 10,
		Status:   true,
		Category: category{
			ID:   "101",
			Name: "category1",
		},
		Comments: []comment{
			{
				Title: "First comment",
				Text:  "A very nice comment",
			},
			{
				Title: "Second comment",
				Text:  "Also very good comment",
			},
			{
				Title: "Third comment",
				Text:  "Last, but not least comment",
			},
		},
	}
	res, err := json.Marshal(prdct)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(res))
}
