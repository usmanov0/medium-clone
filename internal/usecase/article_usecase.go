package usecase

import (
	"example.com/my-medium-clone/internal/common"
	domain2 "example.com/my-medium-clone/internal/domain"
	"example.com/my-medium-clone/internal/repo"
	"time"
)

type ArticleUseCase interface {
	CreateArticle(article *domain2.Article) (int, error)
	GetArticleByID(id int) (*domain2.Article, error)
	GetArticlesByAuthor(authorID int) ([]domain2.Article, error)
	UpdateArticle(article *domain2.Article) error
	DeleteArticle(id int) error
}

type articleUseCase struct {
	articleRepo repo.ArticleRepository
}

func NewArticleUseCase(articleRepo repo.ArticleRepository) ArticleUseCase {
	return &articleUseCase{articleRepo: articleRepo}
}

func (a *articleUseCase) CreateArticle(article *domain2.Article) (int, error) {
	var userId int
	if err := common.ValidateArticle(article); err != nil {
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

func (a *articleUseCase) GetArticleByID(id int) (*domain2.Article, error) {
	article, err := a.articleRepo.FindById(id)
	if err != nil {
		return nil, err
	}

	return article, nil
}

func (a *articleUseCase) GetArticlesByAuthor(authorID int) ([]domain2.Article, error) {
	articles, err := a.articleRepo.FindByAuthor(authorID)
	if err != nil {
		return nil, err
	}

	return articles, nil
}

func (a *articleUseCase) UpdateArticle(article *domain2.Article) error {
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
