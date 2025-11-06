package logic

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"shop/userop/internal/model"

	"shop/userop/internal/svc"
	"shop/userop/userop"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateMessageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateMessageLogic {
	return &CreateMessageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateMessageLogic) CreateMessage(in *userop.MessageRequest) (*userop.MessageResponse, error) {
	var message model.LeavingMessages

	message.User = in.UserId
	message.MessageType = in.MessageType
	message.Subject = in.Subject
	message.Message = in.Message
	message.File = in.File

	if err := l.svcCtx.Db.Save(&message).Error; err != nil {
		return nil, status.Errorf(codes.Internal, "创建留言失败: %v", err)
	}

	return &userop.MessageResponse{Id: message.ID}, nil
}
