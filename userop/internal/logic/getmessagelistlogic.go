package logic

import (
	"context"
	"shop/userop/internal/model"

	"shop/userop/internal/svc"
	"shop/userop/userop"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMessageListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMessageListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMessageListLogic {
	return &GetMessageListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMessageListLogic) GetMessageList(in *userop.MessageRequest) (*userop.MessageListResponse, error) {
	var rsp userop.MessageListResponse
	var messages []model.LeavingMessages
	var messageList []*userop.MessageResponse

	result := l.svcCtx.Db.Where(&model.LeavingMessages{User: in.UserId}).Find(&messages)
	rsp.Total = int32(result.RowsAffected)

	for _, message := range messages {
		messageList = append(messageList, &userop.MessageResponse{
			Id:          message.ID,
			UserId:      message.User,
			MessageType: message.MessageType,
			Subject:     message.Subject,
			Message:     message.Message,
			File:        message.File,
		})
	}

	rsp.Data = messageList
	return &rsp, nil
}
