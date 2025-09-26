package dao

import (
	"context"
	"gorm.io/gorm"
	"reflect"
	"shop/user/ioc"
	"shop/user/repository/model"
	"testing"
)

func TestNewGormUserDao(t *testing.T) {
	db := ioc.InitDB()
	NewGormUserDao(db)
}

func TestNewUserDaoImpl(t *testing.T) {
	type args struct {
		db *gorm.DB
	}
	tests := []struct {
		name string
		args args
		want *UserDaoImpl
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserDaoImpl(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserDaoImpl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserDaoImpl_CheckPasswordHash(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx          context.Context
		password     string
		hashPassword string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserDaoImpl{
				db: tt.fields.db,
			}
			got, err := u.CheckPasswordHash(tt.args.ctx, tt.args.password, tt.args.hashPassword)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckPasswordHash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CheckPasswordHash() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserDaoImpl_CreateUser(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx context.Context
		req model.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    model.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserDaoImpl{
				db: tt.fields.db,
			}
			got, err := u.CreateUser(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserDaoImpl_GetUserById(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    model.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserDaoImpl{
				db: tt.fields.db,
			}
			got, err := u.GetUserById(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserById() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserDaoImpl_GetUserByMobile(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx    context.Context
		mobile string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    model.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserDaoImpl{
				db: tt.fields.db,
			}
			got, err := u.GetUserByMobile(tt.args.ctx, tt.args.mobile)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserByMobile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserByMobile() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserDaoImpl_GetUserList(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx   context.Context
		pn    int
		psize int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []model.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserDaoImpl{
				db: tt.fields.db,
			}
			got, err := u.GetUserList(tt.args.ctx, tt.args.pn, tt.args.psize)
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

func TestUserDaoImpl_UpdateUser(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx context.Context
		req model.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    model.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserDaoImpl{
				db: tt.fields.db,
			}
			got, err := u.UpdateUser(tt.args.ctx, tt.args.req)
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
