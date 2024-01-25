package adapters

import (
	"example.com/my-medium-clone/internal/category/domain"
	"github.com/jackc/pgx"
)

type categoryRepository struct {
	db *pgx.Conn
}

func NewCategoryRepo(db *pgx.Conn) domain.CategoryRepository {
	return &categoryRepository{db: db}
}

func (c *categoryRepository) Save(category *domain.Category) (int, error) {
	query := `
		INSERT INTO category(name)
		VALUES($1)
		RETURNING id`

	var categoryId int
	err := c.db.QueryRow(
		query,
		category.Name,
	).Scan(categoryId)
	if err != nil {
		return 0, err
	}
	return categoryId, nil
}

func (c *categoryRepository) GetCategoryById(id int) (*domain.Category, error) {
	query :=
		`SELECT c.name
		FROM category c
		WHERE c.id = $1 
		`

	var category domain.Category
	err := c.db.QueryRow(
		query,
		id).Scan(&category.Id, &category.Name)

	if err != nil {
		return nil, err
	}

	return &category, nil
}

func (c *categoryRepository) UpdateCategory(id int, category *domain.Category) error {
	query :=
		`UPDATE category
		SET name = name = $2
		WHERE id = $1`

	_, err := c.db.Exec(query, id, category.Name)
	if err != nil {
		return err
	}
	return nil
}

func (c *categoryRepository) Delete(id int) error {
	query :=
		`DELETE FROM category
		WHERE id = $1`

	_, err := c.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
