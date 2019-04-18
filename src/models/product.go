package models

type Product struct {
	Serial      string `json:"serial"`
	Name        string `json:"name"`
	Description string `json:"desc"`
	Quantity    int    `json:"quantity"`
}
