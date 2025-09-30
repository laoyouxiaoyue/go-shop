package api

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
	"shop/web/forms"
	"shop/web/global"
	"shop/web/global/reponse"
	"shop/web/middlewares"
	"shop/web/models"
	"shop/web/proto"
	"shop/web/utils"
	"strconv"
	"time"
)

func HandleGrpcErrorToHttp(err error, c *gin.Context) {
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				c.JSON(http.StatusNotFound, gin.H{
					"msg": e.Message(),
				})
			case codes.Internal:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg": "内部错误",
				})
			case codes.InvalidArgument:
				c.JSON(http.StatusBadRequest, gin.H{
					"msg": e.Message(),
				})
			case codes.Unavailable:
				c.JSON(http.StatusServiceUnavailable, gin.H{
					"msg": "用户服务不可用",
				})
			default:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg": e.Code(),
				})
			}
		}
	}
}
func GetUserList(ctx *gin.Context) {
	userConn, err := grpc.Dial(fmt.Sprintf("%s:%s", global.ServerConfig.UserServer.Host, global.ServerConfig.UserServer.Port), grpc.WithInsecure())
	if err != nil {
		zap.S().Errorw("[GetUserList] 连接错误", "err", err)
	}
	UserSrvClient := proto.NewUserClient(userConn)

	pn := ctx.DefaultQuery("pn", "0")
	pnInt, _ := strconv.Atoi(pn)
	pSize := ctx.DefaultQuery("psize", "0")
	pSizeInt, _ := strconv.Atoi(pSize)
	rsp, err := UserSrvClient.GetUserList(context.Background(), &proto.PageInfo{
		Pn:    uint32(pnInt),
		PSize: uint32(pSizeInt),
	})
	if err != nil {
		zap.S().Errorw("[GetUserList] 查询用户列表错误", "err", err)
		HandleGrpcErrorToHttp(err, ctx)
		return
	}
	result := make([]interface{}, 0)
	for _, value := range rsp.Data {
		user := reponse.UserResponse{
			Id:       value.Id,
			Nickname: value.NickName,
			Birthday: reponse.JsonTime(time.Unix(int64(value.BirthDay), 0)),
			Gender:   value.Gender,
			Mobile:   value.Mobile,
		}

		result = append(result, user)
	}

	ctx.JSON(http.StatusOK, result)
}

func PassWordLogin(c *gin.Context) {
	passwordLoginForm := forms.PassWordLoginForm{}

	if err := c.ShouldBind(&passwordLoginForm); err != nil {
		utils.HandleValidatorError(c, err)
	}
	zap.S().Infof("%v", passwordLoginForm)

	if !store.Verify(passwordLoginForm.CaptchaId, passwordLoginForm.Captcha, true) {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "验证码错误",
		})
		return
	}
	userConn, err := grpc.Dial(fmt.Sprintf("%s:%s", global.ServerConfig.UserServer.Host, global.ServerConfig.UserServer.Port), grpc.WithInsecure())
	if err != nil {
		zap.S().Errorw("[PassWordLogin] 连接错误", "err", err)
	}
	UserSrvClient := proto.NewUserClient(userConn)

	if rsp, err := UserSrvClient.GetUserByMobile(context.Background(), &proto.MobileRequest{
		Mobile: passwordLoginForm.Mobile,
	}); err != nil {
		zap.S().Errorf("%v", err)
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				c.JSON(http.StatusNotFound, map[string]string{
					"msg":    e.Message(),
					"mobile": "用户不存在",
				})
			default:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg": "登录失败,内部错误",
				})
			}
			return
		}
	} else {

		if passRsp, pasErr := UserSrvClient.CheckPassword(context.Background(), &proto.PassWordCheckInfo{
			Password:          passwordLoginForm.PassWord,
			EncryptedPassword: rsp.Password,
		}); pasErr != nil {
			c.JSON(http.StatusInternalServerError, map[string]string{
				"msg": "登录失败",
			})
		} else {
			if passRsp.Success {
				j := middlewares.NewJWT()
				token, err := j.CreateToken(models.CustomClaims{
					ID:          uint64(rsp.Id),
					Nickname:    rsp.NickName,
					AuthorityId: 0,
					StandardClaims: jwt.StandardClaims{
						NotBefore: time.Now().Unix(),
						ExpiresAt: time.Now().Unix() + 86400,
						Issuer:    "Shuai",
					},
				})
				if err != nil {
					c.JSON(http.StatusInternalServerError, map[string]string{
						"msg": "登录失败",
					})
					return
				}
				c.Header("x-token", token)
				c.JSON(http.StatusOK, map[string]string{
					"msg": "登录成功",
				})
			} else {
				c.JSON(http.StatusBadRequest, map[string]string{
					"msg": "登录失败，用户名或密码错误",
				})
			}
		}
	}
}

func Register(c *gin.Context) {
	registerForm := forms.RegisterForm{}
	if err := c.ShouldBind(&registerForm); err != nil {
		utils.HandleValidatorError(c, err)
		return
	}
	smsConn, err := grpc.Dial(fmt.Sprintf("127.0.0.1:50052"), grpc.WithInsecure())
	if err != nil {
		zap.S().Errorw("[SendSms] 连接错误", "err", err)
		HandleGrpcErrorToHttp(err, c)
		return
	}
	CodeSrvClient := proto.NewCodeServiceClient(smsConn)
	rsp, err := CodeSrvClient.VerifyCode(context.Background(), &proto.VerifyCodeRequest{
		Addr:    registerForm.Mobile,
		Subject: registerForm.Subject,
		Code:    registerForm.Code,
	})
	if !rsp.Success || err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "验证码错误",
		})
	}
	userConn, err := grpc.Dial(fmt.Sprintf("%s:%s", global.ServerConfig.UserServer.Host, global.ServerConfig.UserServer.Port), grpc.WithInsecure())
	if err != nil {
		zap.S().Errorw("[PassWordLogin] 连接错误", "err", err)
		HandleGrpcErrorToHttp(err, c)
		return
	}
	UserSrvClient := proto.NewUserClient(userConn)
	UserSrvClient.CreateUser(c, &proto.CreateUserInfo{
		Mobile:   registerForm.Mobile,
		PassWord: registerForm.PassWord,
	})

}
