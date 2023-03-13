package models

type Transaction struct {
	TransactionId         string  `json:"transactionId" dynamo:"TransactionId"`
	CategoryTransactionId string  `json:"-" dynamo:"CategoryTransactionId"`
	Category              string  `json:"category" dynamo:"Category"`
	Value                 float64 `json:"value" dynamo:"Value"`
	Description           string  `json:"description" dynamo:"Description"`
	MonthYear             string  `json:"monthYear" dynamo:"MonthYear"`
	Date                  string  `json:"date" dynamo:"Date"`
	Type                  string  `json:"type" dynamo:"Type"`
	PaymentMethod         string  `json:"paymentMethod" dynamo:"PaymentMethod"`
}

type Category struct {
	Name     string `json:"name" dynamo:"Name"`
	Type     string `json:"type" dynamo:"Type"`
	Priority int    `json:"priority" dynamo:"Priority"`
}
