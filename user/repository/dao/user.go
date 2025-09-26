package dao

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"shop/user/repository/model"
	"shop/user/utils"
)

type UserDao interface {
	GetUserList(ctx context.Context, pn, psize int) ([]model.User, error)
	GetUserByMobile(ctx context.Context, mobile string) (model.User, error)
	GetUserById(ctx context.Context, id string) (model.User, error)
	CreateUser(ctx context.Context, user model.User) (model.User, error)
	UpdateUser(ctx context.Context, user model.User) (model.User, error)
	CheckPasswordHash(ctx context.Context, password, hash string) (bool, error)
}

type UserDaoImpl struct {
	db *gorm.DB
}

func NewUserDaoImpl(db *gorm.DB) *UserDaoImpl {
	return &UserDaoImpl{db: db}
}

func (u *UserDaoImpl) GetUserList(ctx context.Context, pn, psize int) ([]model.User, error) {
	var users []model.User
	result := u.db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	u.db.Scopes(utils.Paginate(int(pn), int(psize))).Find(&users)
	return users, nil
}

func (u *UserDaoImpl) GetUserByMobile(ctx context.Context, mobile string) (model.User, error) {
	var user model.User
	result := u.db.Where("mobile = ?", mobile).First(&user)
	fmt.Printf("%+v", result.Error)
	if result.RowsAffected == 0 {
		return user, status.Error(codes.NotFound, "用户不存在")
	}
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}

func (u *UserDaoImpl) GetUserById(ctx context.Context, id string) (model.User, error) {
	var user model.User
	result := u.db.Where("id = ?", id).First(&user)
	fmt.Printf("%+v", result.Error)
	if result.RowsAffected == 0 {
		return user, status.Error(codes.NotFound, "用户不存在")
	}
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}

func (u *UserDaoImpl) CreateUser(ctx context.Context, req model.User) (model.User, error) {
	result := u.db.Where("mobile = ?", req.Mobile).First(&model.User{})
	if result.RowsAffected != 0 {
		return model.User{}, status.Error(codes.AlreadyExists, "用户已存在")
	}
	var user model.User
	user.Mobile = req.Mobile
	user.NickName = req.NickName
	user.Gender = req.Gender

	var err error
	user.Password, err = utils.HashPassword(req.Password)
	if err != nil {
		return model.User{}, status.Error(codes.Internal, err.Error())
	}
	result = u.db.Create(&user)
	if result.Error != nil {
		return model.User{}, status.Error(codes.Internal, result.Error.Error())
	}
	return user, nil
}

func (u *UserDaoImpl) UpdateUser(ctx context.Context, req model.User) (model.User, error) {
	var user model.User
	result := u.db.Where("id = ?", req.ID).First(&user)
	if result.RowsAffected == 0 {
		return model.User{}, status.Error(codes.NotFound, "用户不存在")
	}

	user.NickName = req.NickName
	user.Birthday = req.Birthday
	user.Gender = req.Gender

	result = u.db.Save(&user)
	if result.Error != nil {
		return model.User{}, status.Errorf(codes.Internal, "内部错误:%s", result.Error.Error())
	}
	return user, nil
}

func (u *UserDaoImpl) CheckPasswordHash(ctx context.Context, password, hashPassword string) (bool, error) {
	err := utils.CheckPassword(password, hashPassword)
	fmt.Printf("%+v", err)
	if err != nil {
		return false, status.Errorf(codes.InvalidArgument, "无效参数: %v", err)
	}
	return true, nil
}

func NewGormUserDao(db *gorm.DB) UserDao {
	return &UserDaoImpl{db}

}
