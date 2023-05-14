package model

type Product struct {
	URL    string   `json:"url"`
	Name   string   `json:"name"`
	Price  string   `json:"price"`
	Rating string   `json:"rating"`
	Images[]string `json:"images"`
}
