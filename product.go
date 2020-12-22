package main

type product struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Price    float32  `json:"price"`
	Quantity int      `json:"quantity"`
	Status   bool     `json:"status"`
	Category category `json:"category"`
}
		