package handlers

import (
	"net/http"
	"strings"
	"sync"

	"github.com/IgorRamos/fm-transaction/internal/models"
	"github.com/IgorRamos/fm-transaction/internal/repositories"
	"github.com/Rhymond/go-money"
	"github.com/aws/aws-lambda-go/events"
	log "github.com/sirupsen/logrus"
)

type transactionsQueryResult struct {
	Transactions []models.Transaction
	Error        error
}

func (h TransactionHandler) ReportTransactions(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	periods := strings.Split(req.QueryStringParameters["periods"], ",")
	if len(periods) == 0 || (len(periods) == 1 && periods[0] == "") {
		return events.APIGatewayProxyResponse{
			Body:       "Query Parameter [periods] required",
			StatusCode: http.StatusBadRequest,
		}, nil
	}

	in := h.getPeriods(periods)
	results := h.getTransactionsResult(in)
	transactions := h.splitTransactions(results)
	report := h.categorizeTransactions(transactions)

	responseBody, err := toJSON(report)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       err.Error(),
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       responseBody,
		Headers: map[string]string{
			"Content-type":                     "application/json",
			"Access-Control-Allow-Origin":      "*",
			"Access-Control-Allow-Credentials": "true",
		},
	}, nil
}

func (h TransactionHandler) getPeriods(periods []string) chan string {
	out := make(chan string)
	go func() {
		for _, period := range periods {
			out <- period
		}
		close(out)
	}()

	return out
}

func (h TransactionHandler) getTransactionsResult(period <-chan string) chan transactionsQueryResult {
	out := make(chan transactionsQueryResult)
	go func() {
		for p := range period {
			page, err := h.transactionRepository.GetTransactions(repositories.QueryParameters{
				Period: p,
				Limit:  1000,
			})
			out <- transactionsQueryResult{Transactions: page.Results}
			if err != nil {
				out <- transactionsQueryResult{Error: err}
			}
		}
		close(out)
	}()

	return out
}

func (h TransactionHandler) splitTransactions(transactions chan transactionsQueryResult) chan models.Transaction {
	out := make(chan models.Transaction)
	go func() {
		for ch := range transactions {
			for _, t := range ch.Transactions {
				out <- t
			}
		}
		close(out)
	}()

	return out
}

func (h TransactionHandler) categorizeTransactions(transactions chan models.Transaction) []models.CategoryReport {
	transactionsByCategory := make(map[string]models.CategoryReport)
	transactionsByParentCategory := make(map[string]models.CategoryReport)

	mu := sync.Mutex{}

	categories := h.getCategories()
	for _, cat := range categories {
		transactionsByParentCategory[cat.GetParentCategory()] = models.CategoryReport{Name: cat.GetParentCategory(), Values: map[string]float64{}}
		transactionsByCategory[cat.Name] = models.CategoryReport{Name: cat.Name, Values: map[string]float64{}}
	}

	for t := range transactions {
		mu.Lock()
		transactionsByParentCategory[t.GetParentCategory()] = mergeTransaction(transactionsByParentCategory[t.GetParentCategory()], t)
		transactionsByCategory[t.Category] = mergeTransaction(transactionsByCategory[t.Category], t)
		mu.Unlock()
	}

	report := []models.CategoryReport{}
	for _, cat := range categories {
		if parentTransactions, ok := transactionsByParentCategory[cat.GetParentCategory()]; ok {
			report = append(report, parentTransactions)
			delete(transactionsByParentCategory, cat.GetParentCategory())
		}
		report = append(report, transactionsByCategory[cat.Name])
	}

	return report
}

func (h TransactionHandler) getCategories() []models.Category {
	categories, err := h.categoryRepository.GetAllCategories()
	if err != nil {
		log.Errorf("failed to get categories, error: [%w]", err)
	}

	return categories
}

func mergeTransaction(category models.CategoryReport, transaction models.Transaction) models.CategoryReport {
	month := extractMonth(transaction.MonthYear)
	if value, ok := category.Values[month]; ok {
		sum, _ := money.NewFromFloat(value, money.BRL).Add(money.NewFromFloat(transaction.Value, money.BRL))
		total, _ := money.NewFromFloat(category.Total, money.BRL).Add(money.NewFromFloat(transaction.Value, money.BRL))
		category.Values[month] = sum.AsMajorUnits()
		category.Total = total.AsMajorUnits()
		return category
	}
	category.Values[month] = transaction.Value
	total, _ := money.NewFromFloat(category.Total, money.BRL).Add(money.NewFromFloat(transaction.Value, money.BRL))
	category.Total = total.AsMajorUnits()
	return category
}

func extractMonth(period string) string {
	return strings.Split(period, "-")[0]
}
