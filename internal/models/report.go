package models

type Report struct {
	Categories map[string]CategoryReport `json:"categories"`
}

type CategoryReport struct {
	Name   string             `json:"name"`
	Total  float64            `json:"total"`
	Values map[string]float64 `json:"values"`
}
