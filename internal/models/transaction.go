package models

type Type string

const (
	TypeCredit Type = "CREDIT"
	TypeDebit  Type = "DEBIT"
)

type Transaction struct {
	TransactionId         string  `json:"transactionId" dynamo:"TransactionId"`
	CategorySubcategoryId string  `json:"-" dynamo:"CategorySubcategoryId"`
	Value                 float32 `json:"value" dynamo:"Value"`
	MonthYear             string  `json:"monthYear" dynamo:"MonthYear"`
	Type                  Type    `json:"type" dynamo:"Type"`
	Category              string  `json:"category" dynamo:"Category"`
	Subcatetegory         string  `json:"subcategory" dynamo:"Subcategory"`
	Date                  string  `json:"date" dynamo:"Date"`
}
