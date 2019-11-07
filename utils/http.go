package utils

import "github.com/gin-gonic/gin"

func Request(ctx *gin.Context, key string) string {
	value := ctx.PostForm(key)
	if value == "" {
		value = ctx.Query(key)
	}
	if value == "" {
		value = ctx.Param(key)
	}
	return value
}