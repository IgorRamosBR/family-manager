package handlers

import (
	"encoding/json"

	"github.com/IgorRamos/fm-transaction/internal/repositories"
	log "github.com/sirupsen/logrus"
)

type CategoryHandler struct {
	categoryRepository repositories.CategoryRepository
}

func NewCategoryHandler(categoryRepository repositories.CategoryRepository) CategoryHandler {
	return CategoryHandler{
		categoryRepository: categoryRepository,
	}
}

func toJSON(value interface{}) (string, error) {
	result, err := json.Marshal(value)
	if err != nil {
		log.Errorf("Failed to marshal response body, error: %v", err.Error())
	}
	return string(result), nil
}
