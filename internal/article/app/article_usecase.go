package app

import (
	"example.com/my-medium-clone/internal/article/domain"
	"time"
)

type ArticleUseCase interface {
	CreateArticle(article *domain.Article) (int, error)
	GetArticleByID(id int) (*domain.Article, error)
	GetArticlesByAuthor(authorID int) ([]*domain.Article, error)
	UpdateArticle(article *domain.Article) error
	DeleteArticle(id int) error
}

type articleUseCase struct {
	articleRepo domain.ArticleRepository
}

func NewArticleUseCase(articleRepo domain.ArticleRepository) ArticleUseCase {
	return &articleUseCase{articleRepo: articleRepo}
}

func (a *articleUseCase) CreateArticle(article *domain.Article) (int, error) {
	var userId int
	if err := domain.ValidateArticle(article); err != nil {
		return 0, err
	}
	article.AuthorId = userId
	article.CreatedAt = time.Now()

	articleId, err := a.articleRepo.Save(article)
	if err != nil {
		return 0, err
	}
	return articleId, nil
}

func (a *articleUseCase) GetArticleByID(id int) (*domain.Article, error) {
	article, err := a.articleRepo.FindById(id)
	if err != nil {
		return nil, err
	}

	return article, nil
}

func (a *articleUseCase) GetArticlesByAuthor(authorID int) ([]*domain.Article, error) {
	articles, err := a.articleRepo.FindByAuthor(authorID)
	if err != nil {
		return nil, err
	}

	return articles, nil
}

func (a *articleUseCase) UpdateArticle(article *domain.Article) error {
	if err := domain.ValidateArticle(article); err != nil {
		return err
	}

	err := a.articleRepo.Update(article)
	if err != nil {
		return err
	}

	return nil
}

func (a *articleUseCase) DeleteArticle(id int) error {
	err := a.articleRepo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
