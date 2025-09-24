package main

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"shop/web/global"
	"shop/web/initialize"
	validator2 "shop/web/validator"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	initialize.InitLogger()
	initialize.InitConfig()
	err := initialize.InitTrans("zh")

	//// 自定义验证器
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 注册自定义验证规则
		err := v.RegisterValidation("mobile", validator2.ValidateMobile)
		if err != nil {
			return
		}

		// 注册字段名函数（让错误信息更友好）
		_ = v.RegisterTranslation("mobile", global.Trans,
			func(ut ut.Translator) error {
				return ut.Add("mobile", "{0} 非法的手机号码!", true)
			},
			func(ut ut.Translator, fe validator.FieldError) string {
				t, _ := ut.T("mobile", fe.Field())
				return t
			},
		)
	}

	if err != nil {
		return
	}
	Router := initialize.Routers()
	port := global.ServerConfig.Port

	zap.S().Infof("启动服务器,端口:%s", port)

	err = Router.Run(fmt.Sprintf(":%s", port))
	if err != nil {
		zap.S().Panic("启动失败:", zap.Error(err))
		return
	}
}
