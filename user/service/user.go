package service

import (
	"context"
	"shop/user/domain"
	"shop/user/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (u *UserService) GetUserList(ctx context.Context, pn, psize int) ([]domain.User, error) {
	return u.repo.GetUserList(ctx, pn, psize)
}
func (u *UserService) GetUserByMobile(ctx context.Context, mobile string) (domain.User, error) {
	return u.repo.GetUserByMobile(ctx, mobile)
}
func (u *UserService) GetUserById(ctx context.Context, id string) (domain.User, error) {
	return u.repo.GetUserById(ctx, id)

}
func (u *UserService) CreateUser(ctx context.Context, user domain.User) (domain.User, error) {
	return u.repo.CreateUser(ctx, user)
}
func (u *UserService) UpdateUser(ctx context.Context, user domain.User) (domain.User, error) {
	return u.repo.UpdateUser(ctx, user)
}
func (u *UserService) CheckPasswordHash(ctx context.Context, password, hash string) (bool, error) {
	return u.repo.CheckPasswordHash(ctx, password, hash)
}
