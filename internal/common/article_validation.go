package common

import (
	"example.com/my-medium-clone/internal/domain"
	"example.com/my-medium-clone/internal/errors"
)

func ValidateArticle(article *domain.Article) error {
	if len(article.Title) == 0 {
		return errors.ErrArticleTitleEmpty
	}
	if len(article.Body) == 0 {
		return errors.ErrArticleBodyEmpty
	}
	if len(article.Title) > 255 {
		return errors.ErrArticleExceedTitle
	}
	if len(article.Body) > 5000 {
		return errors.ErrArticleExceedBody
	}
	return nil
}
