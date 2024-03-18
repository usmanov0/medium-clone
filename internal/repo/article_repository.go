package repo

import "example.com/my-medium-clone/internal/domain"

type ArticleRepository interface {
	Save(article *domain.Article) (int, error)
	FindById(id int) (*domain.Article, error)
	FindByAuthor(authorID int) ([]domain.Article, error)
	FindByCategory(categoryID int) ([]domain.Article, error)
	FindPublishedArticles() ([]domain.Article, error)
	Update(article *domain.Article) error
	Delete(id int) error
}
