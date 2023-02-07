package models

import "strings"

type Report struct {
	Categories map[string]CategoryReport `json:"categories"`
}

type CategoryReport struct {
	Name     string             `json:"name"`
	Total    float64            `json:"total"`
	IsParent bool               `json:"isParent"`
	Values   map[string]float64 `json:"values"`
}

func (c CategoryReport) GetParentCategory() string {
	return strings.Split(c.Name, " - ")[0]
}
