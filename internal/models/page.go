package models

type Page struct {
	Results []Transaction `json:"results"`
	Next    string        `json:"next"`
}
