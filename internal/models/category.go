package models

type CategoryType string

const (
	ExpenseCategoryType CategoryType = "EXPENSE"
	IncomeCategoryType  CategoryType = "INCOME"
)

type Category struct {
	Name     string       `json:"name" dynamo:"name"`
	Type     CategoryType `json:"type" dynamo:"type"`
	Priority int          `json:"priority" dynamo:"priority"`
}
