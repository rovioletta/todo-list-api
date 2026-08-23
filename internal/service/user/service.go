package user

import "rovioletta/todo-list-api/internal/db"

type Service struct {
	db *db.Queries
}

func NewService(db *db.Queries) *Service {
	return &Service{
		db: db,
	}
}
