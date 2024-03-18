package repo

import "example.com/my-medium-clone/internal/domain"

type CategoryRepository interface {
	Save(category *domain.Category) (int, error)
	GetCategoryById(id int) (*domain.Category, error)
	UpdateCategory(id int, category *domain.Category) error
	Delete(id int) error
	//GetAllCategories() ([]Category, error)
	//SearchCategories(criteria string) ([]*Category, error)
	//DeleteCategory(id int) error
}
