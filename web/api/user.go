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
	codev1 "shop/api/gen/code"
	"shop/api/gen/user"
	"shop/web/forms"
	"shop/web/global"
	"shop/web/global/reponse"
	"shop/web/middlewares"
	"shop/web/models"
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
	addr, err := utils.DiscoverAddr("user-service")
	if err != nil {
		HandleGrpcErrorToHttp(err, ctx)
		return
	}
	userConn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		zap.S().Errorw("[GetUserList] 连接错误", "err", err)
	}
	UserSrvClient := userv1.NewUserClient(userConn)

	pn := ctx.DefaultQuery("pn", "0")
	pnInt, _ := strconv.Atoi(pn)
	pSize := ctx.DefaultQuery("psize", "0")
	pSizeInt, _ := strconv.Atoi(pSize)
	rsp, err := UserSrvClient.GetUserList(context.Background(), &userv1.PageInfo{
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
	UserSrvClient := userv1.NewUserClient(userConn)

	if rsp, err := UserSrvClient.GetUserByMobile(context.Background(), &userv1.MobileRequest{
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

		if passRsp, pasErr := UserSrvClient.CheckPassword(context.Background(), &userv1.PassWordCheckInfo{
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
	codeAddr, err := utils.DiscoverAddr("code-service")
	if err != nil {
		HandleGrpcErrorToHttp(err, c)
		return
	}
	smsConn, err := grpc.Dial(codeAddr, grpc.WithInsecure())
	if err != nil {
		zap.S().Errorw("[SendSms] 连接错误", "err", err)
		HandleGrpcErrorToHttp(err, c)
		return
	}
	CodeSrvClient := codev1.NewCodeServiceClient(smsConn)
	rsp, err := CodeSrvClient.VerifyCode(context.Background(), &codev1.VerifyCodeRequest{
		Addr:    registerForm.Mobile,
		Subject: registerForm.Subject,
		Code:    registerForm.Code,
	})
	if !rsp.Success || err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "验证码错误",
		})
	}
	addr, err := utils.DiscoverAddr("user-service")
	if err != nil {
		HandleGrpcErrorToHttp(err, c)
		return
	}
	userConn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		zap.S().Errorw("[PassWordLogin] 连接错误", "err", err)
		HandleGrpcErrorToHttp(err, c)
		return
	}
	UserSrvClient := userv1.NewUserClient(userConn)
	UserSrvClient.CreateUser(c, &userv1.CreateUserInfo{
		Mobile:   registerForm.Mobile,
		PassWord: registerForm.PassWord,
	})

}
