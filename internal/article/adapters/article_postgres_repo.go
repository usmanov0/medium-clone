package adapters

import (
	"example.com/my-medium-clone/internal/article/domain"
	"example.com/my-medium-clone/internal/errors"
	"github.com/jackc/pgx"
)

type articleRepository struct {
	db *pgx.Conn
}

func NewArticleRepo(db *pgx.Conn) domain.ArticleRepository {
	return &articleRepository{db: db}
}

func (a *articleRepository) Save(article *domain.Article) (int, error) {
	query :=
		`INSERT INTO articles(title,body,author_id,category_id,is_draft,published_at,created_at,updated_at)
		VALUES($1,$2,$3,$4,$5,$6,$7,$8)
		RETURNING id `

	var articleId int
	err := a.db.QueryRow(
		query,
		article.Title,
		article.Body,
		article.AuthorId,
		article.CategoryId,
		article.IsDraft,
		article.PublishedAt,
		article.CreatedAt,
		article.UpdatedAt,
	).Scan(&articleId)

	if err != nil {
		return 0, errors.ErrIdScanFailed
	}

	return articleId, nil
}

func (a *articleRepository) FindById(id int) (*domain.Article, error) {
	query :=
		`SELECT id, title, body, author_id, category_id, is_draft, published_at, created_at, updated_at
		FROM articles
		WHERE id = $1`

	var article domain.Article
	err := a.db.QueryRow(
		query,
		id,
	).Scan(
		&article.Id,
		&article.Title,
		&article.Body,
		&article.AuthorId,
		&article.CategoryId,
		&article.IsDraft,
		&article.PublishedAt,
		&article.CreatedAt,
		&article.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &article, nil
}

func (a *articleRepository) FindByAuthor(authorID int) ([]domain.Article, error) {
	query := `
		SELECT id, title, body, author_id, category_id, is_draft, published_at, created_at, updated_at
		FROM articles
		WHERE author_id = $1
	`
	rows, err := a.db.Query(query, authorID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var articles []domain.Article
	for rows.Next() {
		var article domain.Article
		err := rows.Scan(
			&article.Id,
			&article.Title,
			&article.Body,
			&article.AuthorId,
			&article.CategoryId,
			&article.IsDraft,
			&article.PublishedAt,
			&article.CreatedAt,
			&article.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		articles = append(articles, article)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return articles, nil
}

func (a *articleRepository) FindByCategory(categoryID int) ([]domain.Article, error) {
	query := `
		SELECT id, title, body, author_id, category_id, is_draft, published_at, created_at, updated_at
		FROM articles
		WHERE author_id = $1
	`

	rows, err := a.db.Query(query, categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var articles []domain.Article
	for rows.Next() {
		var article domain.Article
		err := rows.Scan(
			&article.Id,
			&article.Title,
			&article.Body,
			&article.AuthorId,
			&article.CategoryId,
			&article.IsDraft,
			&article.PublishedAt,
			&article.CreatedAt,
			&article.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		articles = append(articles, article)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return articles, nil
}

func (a *articleRepository) FindPublishedArticles() ([]domain.Article, error) {
	query := `
		SELECT id, title, body, author_id, category_id, is_draft, published_at, created_at, updated_at
		FROM articles
		WHERE is_draft = false AND published_at IS NOT NULL
	`
	rows, err := a.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var articles []domain.Article
	for rows.Next() {
		var article domain.Article
		err := rows.Scan(
			&article.Id,
			&article.Title,
			&article.Body,
			&article.AuthorId,
			&article.CategoryId,
			&article.IsDraft,
			&article.PublishedAt,
			&article.CreatedAt,
			&article.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		articles = append(articles, article)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return articles, nil
}

func (a *articleRepository) Update(article *domain.Article) error {
	query := `
		UPDATE articles
		SET title = $1, body = $2, category_id = $3,updated_at = $4
		WHERE id = $5
	`
	_, err := a.db.Exec(
		query,
		article.Title,
		article.Body,
		article.CategoryId,
		article.UpdatedAt,
		article.Id,
	)
	if err != nil {
		return err
	}
	return nil
}

func (a *articleRepository) Delete(id int) error {
	query := `
		DELETE FROM articles
		WHERE id = $1
	`

	_, err := a.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
