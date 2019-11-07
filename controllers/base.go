package controllers

import (
	"github.com/sirupsen/logrus"
	"github.com/go-redis/redis"
	"github.com/gin-gonic/gin"
	"github.com/yangxinwei/gin-demo-api/config"
	"github.com/yangxinwei/gin-demo-api/utils"
	"net/http"
	"time"
)

type BaseController struct {
	//MySQL
	Logger       *logrus.Logger
	RedisCluster *redis.ClusterClient
	RedisSingle  *redis.Client
	Config       *config.GlobalConfig
	Env          *utils.Env
}

func (c *BaseController) Ping(ctx *gin.Context) {
	type pingResp struct {
		Pong string
	}
	resp := new(pingResp)
	resp.Pong = "pong"

	c.Output(ctx, "200", resp)
}

func (c *BaseController) Output(ctx *gin.Context, status string, data interface{}) {
	var result = "fail"
	if status == "200" {
		result = "success"
	}
	ctx.JSON(http.StatusOK, gin.H{
		"errorCode": status,
		"result":    result,
		"msg":       utils.Status(status),
		"time":      time.Now().UnixNano() / 1e6,
		"data":      data,
	})
	return
}
