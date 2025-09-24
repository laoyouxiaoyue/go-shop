package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"shop/web/global"
	"strings"
)

func HandleValidatorError(c *gin.Context, err error) {
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"error": RemoveTopStruct(errs.Translate(global.Trans)),
	})
	return
}

func RemoveTopStruct(fields map[string]string) map[string]string {
	rsp := make(map[string]string)
	for field, err := range fields {
		// 使用 strings.Cut 来分割字符串
		if _, after, found := strings.Cut(field, "."); found {
			rsp[after] = err
		} else {
			rsp[field] = err
		}
	}
	return rsp
}
