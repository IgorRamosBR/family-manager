package category

import "github.com/IgorRamos/fm-transaction/internal/repositories"

type CategoryHandler struct {
	categoryRepository repositories.CategoryRepository
}

func NewCategoryHandler(categoryRepository repositories.CategoryRepository) CategoryHandler {
	return CategoryHandler{
		categoryRepository: categoryRepository,
	}
}
