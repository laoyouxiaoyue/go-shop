package service

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"
	"shop/user/model"
	"shop/user/proto"
	"shop/user/utils"
	"time"
)

type UserService struct {
	DB *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{
		DB: db,
	}
}

func (s *UserService) GetUserList(ctx context.Context, req *proto.PageInfo) (*proto.UserListResponse, error) {
	var users []model.User
	result := s.DB.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	rsp := &proto.UserListResponse{
		Total: int32(result.RowsAffected),
	}
	s.DB.Scopes(utils.Paginate(int(req.Pn), int(req.PSize))).Find(&users)
	for _, user := range users {
		userInfoRsp := utils.ModelToResponse(user)
		rsp.Data = append(rsp.Data, &userInfoRsp)
	}
	return rsp, nil
}
func (s *UserService) GetUserByMobile(ctx context.Context, req *proto.MobileRequest) (*proto.UserInfoResponse, error) {
	var user model.User
	result := s.DB.Where("mobile = ?", req.Mobile).First(&user)
	fmt.Printf("%+v", result.Error)
	if result.RowsAffected == 0 {
		return nil, status.Error(codes.NotFound, "用户不存在")
	}
	if result.Error != nil {
		return nil, result.Error
	}
	userInfoRsp := utils.ModelToResponse(user)
	return &userInfoRsp, nil
}
func (s *UserService) GetUserById(ctx context.Context, req *proto.IdRequest) (*proto.UserInfoResponse, error) {
	var user model.User
	result := s.DB.Where("id = ?", req.Id).First(&user)
	if result.RowsAffected == 0 {
		return nil, status.Error(codes.NotFound, "用户不存在")
	}
	if result.Error != nil {
		return nil, result.Error
	}
	userInfoRsp := utils.ModelToResponse(user)
	return &userInfoRsp, nil
}

func (s *UserService) CreateUser(ctx context.Context, req *proto.CreateUserInfo) (*proto.UserInfoResponse, error) {
	result := s.DB.Where("mobile = ?", req.Mobile).First(&model.User{})
	if result.RowsAffected != 0 {
		return nil, status.Error(codes.AlreadyExists, "用户已存在")
	}
	var user model.User
	user.Mobile = req.Mobile
	user.NickName = req.NickName
	var err error
	user.Password, err = utils.HashPassword(req.PassWord)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	result = s.DB.Create(&user)
	if result.Error != nil {
		return nil, status.Error(codes.Internal, result.Error.Error())
	}

	userInfoRsp := utils.ModelToResponse(user)
	return &userInfoRsp, nil
}

func (s *UserService) UpdateUser(ctx context.Context, req *proto.UpdateUserInfo) (*emptypb.Empty, error) {
	var user model.User
	result := s.DB.Where("id = ?", req.Id).First(&user)
	if result.RowsAffected == 0 {
		return nil, status.Error(codes.NotFound, "用户不存在")
	}

	birthDay := time.Unix(int64(req.BirthDay), 0)
	user.NickName = req.NickName
	user.Birthday = &birthDay
	user.Gender = req.Gender

	result = s.DB.Save(&user)
	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, "内部错误:%s", result.Error.Error())
	}
	return &empty.Empty{}, nil
}

func (s *UserService) CheckPassword(ctx context.Context, req *proto.PassWordCheckInfo) (*proto.CheckResponse, error) {
	err := utils.CheckPassword(req.Password, req.EncryptedPassword)
	fmt.Printf("%+v", err)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "无效参数: %v", err)
	}
	return &proto.CheckResponse{
		Success: true,
	}, nil
}
