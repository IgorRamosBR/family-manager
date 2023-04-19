package models

import (
	"errors"
	"regexp"
)

type Type string
type PaymentMethod string
type CategoryType string

const (
	TypeIncome  Type = "INCOME"
	TypeExpense Type = "EXPENSE"

	PaymentMethodCreditCard PaymentMethod = "CREDIT_CARD"
	PaymentMethodDebitCard  PaymentMethod = "DEBIT_CARD"
	PaymentMethodPix        PaymentMethod = "PIX"

	ExpenseCategoryType CategoryType = "EXPENSE"
	IncomeCategoryType  CategoryType = "INCOME"
)

type TransactionRequest struct {
	Category      CategoryRequest `json:"category"`
	Value         float64         `json:"value"`
	Description   string          `json:"description"`
	MonthYear     string          `json:"monthYear"`
	Date          string          `json:"date"`
	Type          Type            `json:"type"`
	PaymentMethod PaymentMethod   `json:"paymentMethod"`
}

type CategoryRequest struct {
	Name     string       `json:"name"`
	Type     CategoryType `json:"type"`
	Priority int          `json:"priority"`
}

func (t TransactionRequest) ToDynamoModel(category Category) Transaction {
	return Transaction{
		Category:      t.Category.Name,
		Value:         t.Value,
		Description:   t.Description,
		MonthYear:     t.MonthYear,
		Date:          t.Date,
		Type:          string(category.Type),
		PaymentMethod: string(t.PaymentMethod),
	}
}

func (t TransactionRequest) Validate() error {
	if err := t.Category.Validate(); err != nil {
		return err
	}

	if err := validateValue(t.Value); err != nil {
		return err
	}

	if err := validateMonthYear(t.MonthYear); err != nil {
		return err
	}

	if err := validateDate(t.Date); err != nil {
		return err
	}

	if err := validatePaymentMethod(PaymentMethod(t.PaymentMethod)); err != nil {
		return err
	}

	return nil
}

func validateValue(value float64) error {
	if value <= 0.0 {
		return errors.New("value field is required")
	}
	return nil
}

func validateMonthYear(monthYear string) error {
	pattern := regexp.MustCompile(`^(Jan|Feb|Mar|Apr|May|Jun|Jul|Aug|Sep|Oct|Nov|Dec)\-\d{4}$`)
	if monthYear == "" {
		return errors.New("monthYear field is required")
	}
	if !pattern.MatchString(monthYear) {
		return errors.New("monthYear should match with the pattern MMM/yyyy")
	}
	return nil
}

func validateDate(date string) error {
	pattern := regexp.MustCompile(`^(0?[1-9]|[12]\d|3[01])\/(0?[1-9]|1[0-2])\/\d{4}$`)
	if date == "" {
		return errors.New("date field is required")
	}
	if !pattern.MatchString(date) {
		return errors.New("date should match with the pattern MMM/yyyy")
	}

	return nil
}

func validatePaymentMethod(paymentMethod PaymentMethod) error {
	switch paymentMethod {
	case PaymentMethodCreditCard, PaymentMethodDebitCard, PaymentMethodPix:
		return nil
	default:
		return errors.New("type should be DEBIT_CARD, CREDIT_CARD or PIX")
	}
}

func (c CategoryRequest) ToDynamoModel() Category {
	return Category{
		Name:     c.Name,
		Type:     string(c.Type),
		Priority: c.Priority,
	}
}

func (c CategoryRequest) Validate() error {
	if err := validateName(c.Name); err != nil {
		return err
	}
	if err := validateCategoryType(c.Type); err != nil {
		return err
	}
	return nil
}

func validateName(name string) error {
	if name == "" {
		return errors.New("name is required")
	}
	return nil
}

func validateCategoryType(categoryType CategoryType) error {
	switch categoryType {
	case ExpenseCategoryType, IncomeCategoryType:
		return nil
	default:
		return errors.New("type should be EXPENSE or INCOME")
	}
}
