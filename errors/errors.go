package errors

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrInvalidParams = status.Error(codes.InvalidArgument, "请求参数不合法")
	ErrDBQuery       = status.Error(codes.Internal, "数据库查询失败")
	ErrDataNotFound  = status.Error(codes.NotFound, "数据不存在")
)
