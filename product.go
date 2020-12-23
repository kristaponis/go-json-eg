package main

import "fmt"

type product struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	Price    float32   `json:"price"`
	Quantity int       `json:"quantity"`
	Status   bool      `json:"status"`
	Category category  `json:"category"`
	Comments []comment `json:"comments"`
}

func (p product) toStr() string {
	return fmt.Sprintf("Id: %s\nName: %s\nPrice: %f\nQuantity: %d\nStatus: %t", p.ID, p.Name, p.Price, p.Quantity, p.Status)
}
