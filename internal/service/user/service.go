package user

import (
	validate "buf.build/go/protovalidate"
	"rovioletta/todo-list-api/internal/db"
	pb "rovioletta/todo-list-api/pkg/pb/user"
)

type UserService struct {
	pb.UnimplementedAuthServiceServer

	db *db.DB
	v  validate.Validator
}

func NewUserService(db *db.DB) *UserService {
	v, err := validate.New()
	if err != nil {
		panic(err)
	}

	return &UserService{
		db: db,
		v:  v,
	}
}
