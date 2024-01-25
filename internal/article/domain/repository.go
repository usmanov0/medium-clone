package domain

type ArticleRepository interface {
	Save(article *Article) (int, error)
	FindById(id int) (*Article, error)
	FindByAuthor(authorID int) ([]Article, error)
	FindByCategory(categoryID int) ([]Article, error)
	FindPublishedArticles() ([]Article, error)
	Update(article *Article) error
	Delete(id int) error
}
