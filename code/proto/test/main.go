package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"shop/code/proto"
)

//func main() {
//	userConn, err := grpc.Dial(fmt.Sprintf("127.0.0.1:50052"), grpc.WithInsecure())
//	if err != nil {
//		zap.S().Errorw("[GetUserList] 连接错误", "err", err)
//	}
//	UserSrvClient := proto.NewCodeServiceClient(userConn)
//	_, err = UserSrvClient.SendCode(context.Background(), &proto.SendCodeRequest{
//		Addr:    "996948441@qq.com",
//		Subject: "register",
//	})
//	if err != nil {
//		return
//	}
//}

func main() {
	userConn, err := grpc.Dial(fmt.Sprintf("127.0.0.1:50052"), grpc.WithInsecure())
	UserSrvClient := proto.NewCodeServiceClient(userConn)
	rsp, err := UserSrvClient.VerifyCode(context.Background(), &proto.VerifyCodeRequest{
		Addr:    "996948441@qq.com",
		Subject: "register",
		Code:    "555188",
	})
	fmt.Printf("%v", rsp)
	if err != nil {
		return
	}
}
