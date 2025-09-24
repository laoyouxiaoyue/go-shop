package initialize

import (
	"errors"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
	"shop/web/global"
	"strings"
)

func InitTrans(locale string) error {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
		zhT := zh.New()
		enT := en.New()

		uni := ut.New(enT, zhT, enT)
		global.Trans, ok = uni.GetTranslator(locale)
		if !ok {
			return errors.New("uni.GetTranslator failed")
		}
		switch locale {
		case "zh":
			err := zh_translations.RegisterDefaultTranslations(v, global.Trans)
			if err != nil {
				return err
			}
		case "en":
			err := en_translations.RegisterDefaultTranslations(v, global.Trans)
			if err != nil {
				return err
			}
		default:
			err := en_translations.RegisterDefaultTranslations(v, global.Trans)
			if err != nil {
				return err
			}
		}
		return nil

	}
	return nil
}
