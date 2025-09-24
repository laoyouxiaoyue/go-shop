package validator

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

func ValidateMobile(fl validator.FieldLevel) bool {
	mobile := fl.Field().String()
	if mobile == "" {
		return false
	}

	// 中国手机号正则：1开头，第二位3-9，总共11位
	mobileRegex := `^1[3-9]\d{9}$`
	matched, _ := regexp.MatchString(mobileRegex, mobile)
	return matched
}
