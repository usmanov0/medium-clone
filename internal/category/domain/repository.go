package domain

type CategoryRepository interface {
	Save(category *Category) (int, error)
	GetCategoryById(id int) (*Category, error)
	UpdateCategory(id int, category *Category) error
	Delete(id int) error
	//GetAllCategories() ([]Category, error)
	//SearchCategories(criteria string) ([]*Category, error)
	//DeleteCategory(id int) error
}
