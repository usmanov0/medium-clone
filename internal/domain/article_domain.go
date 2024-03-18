package domain

import "time"

type Article struct {
	Id          int
	Title       string
	Body        string
	AuthorId    int
	CategoryId  int
	IsDraft     bool
	PublishedAt time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
