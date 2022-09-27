package main

type Book struct {
	Name   string  `json:"name"`
	Id     string  `json:"id"`
	Author string  `json: "author"`
	Price  float64 `json:"price"`
}
