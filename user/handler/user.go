package handler

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	"shop/user/proto"
	"shop/user/service"
)

type UserServer struct {
	service *service.UserService
	proto.UnimplementedUserServer
}

func NewUserServer(userService *service.UserService) *UserServer {
	return &UserServer{
		service: userService,
	}
}
func (s *UserServer) GetUserList(ctx context.Context, req *proto.PageInfo) (*proto.UserListResponse, error) {
	return s.service.GetUserList(ctx, req)
}
func (s *UserServer) GetUserByMobile(ctx context.Context, req *proto.MobileRequest) (*proto.UserInfoResponse, error) {
	return s.service.GetUserByMobile(ctx, req)
}
func (s *UserServer) GetUserById(ctx context.Context, req *proto.IdRequest) (*proto.UserInfoResponse, error) {
	return s.service.GetUserById(ctx, req)
}

func (s *UserServer) CreateUser(ctx context.Context, req *proto.CreateUserInfo) (*proto.UserInfoResponse, error) {
	return s.service.CreateUser(ctx, req)
}

func (s *UserServer) UpdateUser(ctx context.Context, req *proto.UpdateUserInfo) (*emptypb.Empty, error) {
	return s.service.UpdateUser(ctx, req)
}

func (s *UserServer) CheckPassword(ctx context.Context, req *proto.PassWordCheckInfo) (*proto.CheckResponse, error) {
	return s.service.CheckPassword(ctx, req)
}
