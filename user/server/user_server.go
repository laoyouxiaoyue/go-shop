package server

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"shop/api/gen/user"
	"shop/user/domain"
	"shop/user/service"
	"strconv"
	"time"
)

type UserServer struct {
	service *service.UserService
	userv1.UnimplementedUserServer
}

func NewUserServer(s *service.UserService) *UserServer {
	return &UserServer{
		service: s,
	}

}
func Domain2Rsp(user domain.User) userv1.UserInfoResponse {
	birthday, err := time.Parse("2006-01-02", user.Birthday)
	if err != nil {
		fmt.Println("日期解析错误:", err)
		return userv1.UserInfoResponse{}
	}
	// 转换为Unix时间戳（uint64）
	timestamp := uint64(birthday.Unix())
	return userv1.UserInfoResponse{
		Id:       int32(user.Id),
		NickName: user.Nickname,
		Password: user.Password,
		Gender:   user.Gender,
		BirthDay: timestamp,
	}
}
func (s *UserServer) Register(server grpc.ServiceRegistrar) {
	userv1.RegisterUserServer(server, &UserServer{})
}
func (s *UserServer) GetUserList(ctx context.Context, req *userv1.PageInfo) (*userv1.UserListResponse, error) {
	list, err := s.service.GetUserList(ctx, int(req.GetPn()), int(req.GetPSize()))
	if err != nil {
		return nil, err
	}
	rsp := &userv1.UserListResponse{
		Total: int32(len(list)),
	}
	for _, user := range list {
		userInfoRsp := Domain2Rsp(user)
		rsp.Data = append(rsp.Data, &userInfoRsp)
	}
	return rsp, nil

}
func (s *UserServer) GetUserByMobile(ctx context.Context, req *userv1.MobileRequest) (*userv1.UserInfoResponse, error) {
	user, err := s.service.GetUserByMobile(ctx, req.GetMobile())
	if err != nil {
		return nil, err
	}
	rsp := Domain2Rsp(user)
	return &rsp, nil
}
func (s *UserServer) GetUserById(ctx context.Context, req *userv1.IdRequest) (*userv1.UserInfoResponse, error) {
	user, err := s.service.GetUserById(ctx, string(req.GetId()))
	if err != nil {
		return nil, err
	}
	rsp := Domain2Rsp(user)
	return &rsp, nil
}

func (s *UserServer) CreateUser(ctx context.Context, req *userv1.CreateUserInfo) (*userv1.UserInfoResponse, error) {
	user, err := s.service.CreateUser(ctx, domain.User{})
	if err != nil {
		return nil, err
	}
	rsp := Domain2Rsp(user)
	return &rsp, nil

}
func RspToDomain(req *userv1.UpdateUserInfo) domain.User {
	return domain.User{
		Id:       int64(req.GetId()),
		Nickname: req.GetNickName(),
		Gender:   req.GetGender(),
		Birthday: strconv.FormatUint(req.BirthDay, 10),
	}
}
func (s *UserServer) UpdateUser(ctx context.Context, req *userv1.UpdateUserInfo) (*emptypb.Empty, error) {
	_, err := s.service.UpdateUser(ctx, RspToDomain(req))
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *UserServer) CheckPassword(ctx context.Context, req *userv1.PassWordCheckInfo) (*userv1.CheckResponse, error) {
	hash, err := s.service.CheckPasswordHash(ctx, req.GetPassword(), req.GetEncryptedPassword())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "无效参数: %v", err)
	}
	if !hash {
		return &userv1.CheckResponse{
			Success: false,
		}, nil
	}
	return &userv1.CheckResponse{
		Success: true,
	}, nil
}
