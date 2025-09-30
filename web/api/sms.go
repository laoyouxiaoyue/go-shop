package api

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"shop/code/proto"
	"shop/web/forms"
	"shop/web/utils"
)

func SendSms(c *gin.Context) {
	sendSmsForm := forms.SendSmsForm{}

	if err := c.ShouldBind(&sendSmsForm); err != nil {
		utils.HandleValidatorError(c, err)
	}
	smsConn, err := grpc.Dial(fmt.Sprintf("127.0.0.1:50052"), grpc.WithInsecure())
	if err != nil {
		zap.S().Errorw("[SendSms] 连接错误", "err", err)
	}
	UserSrvClient := proto.NewCodeServiceClient(smsConn)
	rsp, err := UserSrvClient.SendCode(context.Background(), &proto.SendCodeRequest{
		Addr:    sendSmsForm.Mobile,
		Subject: sendSmsForm.Subject,
	})
	if !rsp.Success || err != nil {
		c.JSON(200, map[string]string{
			"msg": "验证码发送失败，请重试",
		})
	}
}
