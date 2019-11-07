package mysql

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/yangxinwei/gin-demo-api/config"
	"sync"
	"time"
)

var engine *xorm.Engine
var onceDo = new(sync.Once)

func InitEngine() {
	onceDo.Do(func() {
		var err error
		if dsn := config.GetGlobalConfig().Mysql.DSN; dsn != "" {
			engine, err = xorm.NewEngine("mysql", dsn)
			if err != nil {
				//yrdLog.GetLogger().Alert("错误", "数据库连接错误")
				panic("数据库连接错误")
			}
			engine.ShowSQL(true)
			ping(0)
			go cyclePing()
		} else {
			//yrdLog.GetLogger().Alert("配置", "数据库配置错误")
			panic("数据库配置错误")
		}
	})
}

func GetEngine() *xorm.Engine {
	InitEngine()
	return engine
}

func cyclePing() {
	for {
		time.Sleep(10 * time.Minute)
		ping(0)
	}
}

func ping(tryCount int8) {
	if tryCount > 3 {
		return
	}
	if err := engine.Ping(); err != nil {
		//yrdLog.GetLogger().Alert("ping错误", err)
		time.Sleep(time.Second * 3)
		ping(tryCount + 1)
	}
}
