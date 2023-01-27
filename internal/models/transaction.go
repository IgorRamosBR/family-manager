package models

type Type string
type PaymentMethod string

const (
	TypeIncome  Type = "INCOME"
	TypeExpense Type = "EXPENSE"

	PaymentMethodCreditCard PaymentMethod = "CREDIT_CARD"
	PaymentMethodDebitCard  PaymentMethod = "DEBIT_CARD"
	PaymentMethodPix        PaymentMethod = "PIX"
)

type Transaction struct {
	TransactionId         string        `json:"transactionId" dynamo:"TransactionId"`
	CategorySubcategoryId string        `json:"-" dynamo:"CategorySubcategoryId"`
	Category              string        `json:"category" dynamo:"Category"`
	Value                 float64       `json:"value" dynamo:"Value"`
	Description           string        `json:"description" dynamo:"Description"`
	MonthYear             string        `json:"monthYear" dynamo:"MonthYear"`
	Date                  string        `json:"date" dynamo:"Date"`
	Type                  Type          `json:"type" dynamo:"Type"`
	PaymentMethod         PaymentMethod `json:"paymentMethod" dynamo:"PaymentMethod"`
}
