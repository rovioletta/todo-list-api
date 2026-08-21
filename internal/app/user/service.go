package user

import (
	pb "rovioletta/todo-list-api/pkg/pb/user"
)

type UserService struct{
	pb.UnimplementedAuthServiceServer

}
func NewUserService() *UserService{
	return &UserService{}
}