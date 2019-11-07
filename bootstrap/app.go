package bootstrap

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/yangxinwei/gin-demo-api/config"
	"github.com/yangxinwei/gin-demo-api/controllers"
	"github.com/yangxinwei/gin-demo-api/utils"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type APP struct {
	Router *gin.Engine
	//Config  *config.Config
	Controller *controllers.BaseController
}

func (a *APP) Initialize() {
	a.config()
	config.InitConfig()
	//a.Service = new(controller.Service)
	log.Println("检测服务加载")
	a.logger()
	//a.db()
	//a.queue()
	//a.redis()
	log.Println("服务加载成功")
}

func (a *APP) config() {
	//var err error
	config.InitConfig()

	a.Controller.Config = config.GetGlobalConfig()
	//if err != nil {
	//	panic(err)
	//}
	log.Println("配置加载成功")
}
func (a *APP) db() {

}
func (a *APP) redis()  {
	if config.GetGlobalConfig().Redis.Cluster {
		a.Controller.RedisCluster = redis.NewClusterClient(&redis.ClusterOptions{
			Addrs:        config.GetGlobalConfig().Redis.Addr,
			Password:     config.GetGlobalConfig().Redis.Password,
			DialTimeout:  10 * time.Second,
			ReadTimeout:  50 * time.Second,
			WriteTimeout: 50 * time.Second,
		})
		pong, err := a.Controller.RedisCluster.Ping().Result()
		if err != nil {
			a.Controller.Logger.Fatalln("redis集群链接失败: ", err)
		}
		if pong == "PONG" {
			a.Controller.Logger.Info("redis集群连接成功")
		}
	} else {
		a.Controller.RedisSingle = redis.NewClient(&redis.Options{
			Addr: config.GetGlobalConfig().Redis.Addr[0],
			DB: config.GetGlobalConfig().Redis.Database,
		})
	}
}
func (a *APP) logger() {
	a.Controller.Logger = utils.NewLogger()
}


func (a *APP) Run(addr string ) {
	handleSigterm(func() {
		//a.Controller.MySQL.Close()
		a.Controller.RedisCluster.Close()
	})
	log.Println("路由加载", a.Router)
	srv := &http.Server{
		Addr:    addr,
		Handler: a.Router,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}

func handleSigterm(handleExit func()) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	go func() {
		<-c
		handleExit()
	}()
}
