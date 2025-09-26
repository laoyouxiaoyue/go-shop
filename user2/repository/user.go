package repository

import (
	"context"
	"shop/user2/domain"
	"shop/user2/repository/dao"
	"shop/user2/repository/model"
	"time"
)

type UserRepository interface {
	GetUserList(ctx context.Context, pn, psize int) ([]domain.User, error)
	GetUserByMobile(ctx context.Context, mobile string) (domain.User, error)
	GetUserById(ctx context.Context, id string) (domain.User, error)
	CreateUser(ctx context.Context, user domain.User) (domain.User, error)
	UpdateUser(ctx context.Context, user domain.User) (domain.User, error)
	CheckPasswordHash(ctx context.Context, password, hash string) (bool, error)
}

type GormUserRepository struct {
	dao dao.UserDao
}

func NewGormUserRepository(dao dao.UserDao) UserRepository {
	return &GormUserRepository{dao}
}
func Domain2model(user domain.User) model.User {
	birth, _ := time.Parse("2006-01-02", user.Birthday)
	return model.User{
		NickName: user.Nickname,
		Password: user.Password,
		Gender:   user.Gender,
		Birthday: &birth,
	}
}
func Model2Domain(user model.User) domain.User {
	return domain.User{
		Nickname: user.NickName,
		Password: user.Password,
		Gender:   user.Gender,
		Birthday: user.Birthday.Format("2006-01-02"),
		Id:       int64(user.ID),
	}
}
func (g *GormUserRepository) GetUserList(ctx context.Context, pn, psize int) ([]domain.User, error) {
	result, err := g.dao.GetUserList(ctx, pn, psize)
	if err != nil {
		return nil, err
	}
	res := make([]domain.User, 0)
	for _, user := range result {
		res = append(res, Model2Domain(user))
	}
	return res, nil
}

func (g *GormUserRepository) GetUserByMobile(ctx context.Context, mobile string) (domain.User, error) {
	user, err := g.dao.GetUserByMobile(ctx, mobile)
	if err != nil {
		return domain.User{}, err
	}
	return Model2Domain(user), nil
}

func (g *GormUserRepository) GetUserById(ctx context.Context, id string) (domain.User, error) {
	user, err := g.dao.GetUserById(ctx, id)
	if err != nil {
		return domain.User{}, err
	}
	return Model2Domain(user), nil
}

func (g *GormUserRepository) CreateUser(ctx context.Context, user domain.User) (domain.User, error) {
	createUser, err := g.dao.CreateUser(ctx, Domain2model(user))
	if err != nil {
		return domain.User{}, err
	}
	return Model2Domain(createUser), nil
}

func (g *GormUserRepository) UpdateUser(ctx context.Context, user domain.User) (domain.User, error) {
	updateUser, err := g.dao.UpdateUser(ctx, Domain2model(user))
	if err != nil {
		return domain.User{}, err
	}
	return Model2Domain(updateUser), nil
}

func (g *GormUserRepository) CheckPasswordHash(ctx context.Context, password, hash string) (bool, error) {
	passwordHash, err := g.dao.CheckPasswordHash(ctx, password, hash)
	if err != nil {
		return false, err
	}
	if !passwordHash {
		return false, err
	}
	return true, nil
}
