package models

type Product struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"desc"`
	Quantity    int    `json:"quantity"`
}
