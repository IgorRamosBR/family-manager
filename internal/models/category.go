package models

type CategoryType string

const (
	ExpenseCategoryType CategoryType = "EXPENSE"
	IncomeCategoryType  CategoryType = "INCOME"
)

type Category struct {
	Name     string       `json:"name" dynamo:"Name"`
	Type     CategoryType `json:"type" dynamo:"Type"`
	Priority int          `json:"priority" dynamo:"Priority"`
}
