package server

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"shop/code/errs"
	codev1 "shop/code/proto"
	"shop/code/service"
)

type CodeServer struct {
	codev1.UnimplementedCodeServiceServer
	service *service.CodeService
}

func NewCodeServer(service *service.CodeService) *CodeServer {
	return &CodeServer{
		service: service,
	}
}

func (cs *CodeServer) SendCode(ctx context.Context, req *codev1.SendCodeRequest) (*codev1.SendCodeResponse, error) {
	err := cs.service.SendCode(ctx, req.GetAddr(), req.GetSubject())
	if err != nil {
		return &codev1.SendCodeResponse{
			Success: false,
			Message: err.Error(),
		}, err
	}
	return &codev1.SendCodeResponse{
		Success: true,
		Message: "success",
	}, nil
}
func (cs *CodeServer) VerifyCode(ctx context.Context, req *codev1.VerifyCodeRequest) (*codev1.VerifyCodeResponse, error) {
	flag, err := cs.service.VerifyCode(ctx, req.GetAddr(), req.GetSubject(), req.GetCode())
	if err != nil {
		return &codev1.VerifyCodeResponse{
			Success: false,
			Message: errs.ErrSystemError.Error(),
		}, err
	}
	if flag {
		return &codev1.VerifyCodeResponse{
			Success: true,
			Message: "success",
		}, nil
	} else {
		return &codev1.VerifyCodeResponse{
			Success: false,
			Message: errs.ErrWrongCode.Error(),
		}, err
	}
}
func (cs *CodeServer) SendMessage(ctx context.Context, req *codev1.SendMessageRequest) (*codev1.SendMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
func (cs *CodeServer) Ping(ctx context.Context, req *codev1.PingRequest) (*codev1.PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (cs *CodeServer) Register(server grpc.ServiceRegistrar) {
	codev1.RegisterCodeServiceServer(server, cs)
}
