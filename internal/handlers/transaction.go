package handlers

import (
	"encoding/json"
	"strconv"

	"github.com/IgorRamos/fm-create-transaction/internal/repositories"
	log "github.com/sirupsen/logrus"
)

type TransactionHandler struct {
	transactionRepository repositories.TransactionRepository
}

func NewTransactionHandler(transactionRepository repositories.TransactionRepository) TransactionHandler {
	return TransactionHandler{
		transactionRepository: transactionRepository,
	}
}

func toInt32(t string) (int32, error) {
	if t == "" {
		return 0, nil
	}

	n, err := strconv.Atoi(t)
	if err != nil {
		return 0, err
	}
	return int32(n), nil

}

func toJSON(value interface{}) (string, error) {
	result, err := json.Marshal(value)
	if err != nil {
		log.Errorf("Failed to marshal response body, error: ", err.Error())
	}
	return string(result), nil
}
