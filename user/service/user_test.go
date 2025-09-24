package service

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"reflect"
	"shop/user/model"
	"shop/user/proto"
	"testing"
	"time"
)

func setupTestDB(t *testing.T) *gorm.DB {
	dsn := "root:root@tcp(127.0.0.1:3306)/shop?charset=utf8&parseTime=True&loc=Local"
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		})
	db1, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}
	return db1
}

func TestUserService_CheckPassword(t *testing.T) {
	type fields struct {
		DB *gorm.DB
	}
	type args struct {
		ctx context.Context
		req *proto.PassWordCheckInfo
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *proto.CheckResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &UserService{
				DB: tt.fields.DB,
			}
			got, err := s.CheckPassword(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckPassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CheckPassword() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserService_CreateUser(t *testing.T) {
	db := setupTestDB(t)
	svc := NewUserService(db)

	// 1. 创建用户
	req := &proto.CreateUserInfo{
		NickName: "Alice",
		PassWord: "123456",
		Mobile:   "18888888888",
	}
	resp, err := svc.CreateUser(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, "Alice", resp.NickName)
	assert.Equal(t, "18888888888", resp.Mobile)

	// 2. 通过手机号获取用户
	getReq := &proto.MobileRequest{Mobile: "18888888888"}
	userResp, err := svc.GetUserByMobile(context.Background(), getReq)
	assert.NoError(t, err)
	assert.Equal(t, resp.Id, userResp.Id)
}

func TestUserService_GetUserById(t *testing.T) {
	db := setupTestDB(t)
	svc := NewUserService(db)

	// 1. 创建用户
	req := &proto.CreateUserInfo{
		NickName: "Alice",
		PassWord: "123456",
		Mobile:   "18888888888",
	}
	resp, err := svc.CreateUser(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, "Alice", resp.NickName)
	assert.Equal(t, "18888888888", resp.Mobile)

	// 2. 通过手机号获取用户
	getReq := &proto.MobileRequest{Mobile: "18888888888"}
	userResp, err := svc.GetUserByMobile(context.Background(), getReq)
	assert.NoError(t, err)
	assert.Equal(t, resp.Id, userResp.Id)
}

func TestUserService_GetUserByMobile(t *testing.T) {
	db := setupTestDB(t)
	svc := NewUserService(db)
	//
	//// 1. 创建用户
	//req := &proto.CreateUserInfo{
	//	NickName: "Alice",
	//	PassWord: "123456",
	//	Mobile:   "18888888889",
	//}
	//resp, err := svc.CreateUser(context.Background(), req)
	//assert.NoError(t, err)
	//assert.Equal(t, "Alice", resp.NickName)
	//assert.Equal(t, "18888888889", resp.Mobile)
	//
	//// 2. 通过手机号获取用户
	//getReq := &proto.MobileRequest{Mobile: "18888888889"}
	//	userResp, _ := svc.GetUserByMobile(context.Background(), getReq)
	//	fmt.Printf("%+v\n", userResp)
	passwordRsp, err := svc.CheckPassword(context.Background(), &proto.PassWordCheckInfo{
		Password:          "123456",
		EncryptedPassword: "$2a$10$cNAohNrzqUyOcdlSBIUR3eK.wMbrOxODc4EwJplvMN0DhNy.lggW2",
	})
	fmt.Print(err)
	fmt.Printf("%+v", passwordRsp)

	//assert.Equal(t, resp.Id, userResp.Id)
}

func TestUserService_GetUserList(t *testing.T) {
	type fields struct {
		DB *gorm.DB
	}
	type args struct {
		ctx context.Context
		req *proto.PageInfo
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *proto.UserListResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &UserService{
				DB: tt.fields.DB,
			}
			got, err := s.GetUserList(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserList() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserService_UpdateUser(t *testing.T) {
	type fields struct {
		DB *gorm.DB
	}
	type args struct {
		ctx context.Context
		req *proto.UpdateUserInfo
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *emptypb.Empty
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &UserService{
				DB: tt.fields.DB,
			}
			got, err := s.UpdateUser(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestGetUserList(t *testing.T) {
	db := setupTestDB(t)
	svc := NewUserService(db)

	// 批量插入用户
	for i := 0; i < 5; i++ {
		u := model.User{
			Mobile:   "1888888888" + string(rune('0'+i)),
			NickName: "User" + string(rune('0'+i)),
			Password: "pwd",
		}
		db.Create(&u)
	}

	// 获取第一页
	req := &proto.PageInfo{Pn: 1, PSize: 2}
	rsp, err := svc.GetUserList(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, int32(5), rsp.Total)
	assert.Len(t, rsp.Data, 2)

	// 获取第二页
	req = &proto.PageInfo{Pn: 2, PSize: 2}
	rsp, err = svc.GetUserList(context.Background(), req)
	assert.NoError(t, err)
	assert.Len(t, rsp.Data, 2)
}
