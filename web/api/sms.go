package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SendSms(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{})
}
