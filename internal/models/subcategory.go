package models

type SubCategory struct {
	Name         string `json:"name" dynamo:"Name"`
	CategoryName string `json:"categoryName" dynamo:"CategoryName"`
	Priority     int    `json:"priority" dynamo:"Priority"`
}
