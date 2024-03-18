package handler

import (
	"encoding/json"
	"example.com/my-medium-clone/internal/domain"
	"example.com/my-medium-clone/internal/usecase"
	"log"
	"net/http"
)

type ArticleHandler struct {
	ArticleUseCase usecase.ArticleUseCase
}

func NewArticleHandler(articleUseCase usecase.ArticleUseCase) *ArticleHandler {
	return &ArticleHandler{ArticleUseCase: articleUseCase}
}

//GetArticleByID(id int) (*domain.Article, error)
//GetArticlesByAuthor(authorID int) ([]*domain.Article, error)
//UpdateArticle(article *domain.Article) error
//DeleteArticle(id int) error

func (a *ArticleHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req domain.Article

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid JSON request", http.StatusBadRequest)
	}

	id, err := a.ArticleUseCase.CreateArticle(&req)
	if err != nil {
		log.Println("reQ", req)
		http.Error(w, "failed to create new article", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(id)
}
