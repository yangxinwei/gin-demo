/**
POD_CONSUL_KEY
POD_CONSUL_URL
POD_CONSUL_TOKEN
 */
package config

import (
	"bytes"
	"encoding/json"
	"github.com/hashicorp/consul/api"
	"github.com/hashicorp/consul/watch"
	"log"
	"os"
	"sync"
)

var c *GlobalConfig
var rwmux = new(sync.RWMutex)
var once = new(sync.Once)
var consulKey = "future-bull.server-api/"

func setGlobalConfig(gc *GlobalConfig) {
	rwmux.Lock()
	defer rwmux.Unlock()
	c = gc
}

func getGlobalConfig() *GlobalConfig {
	rwmux.RLock()
	defer rwmux.RUnlock()
	return c
}
type GlobalConfig struct {
	Mysql    MysqlConfig    `json:"mysql"`
	Redis    RedisConfig    `json:"redis"`
	Rabbitmq RabbitmqConfig `json:"rabbitmq"`
	Alert    AlertConfig    `json:"alert"`
	Services ServicesConfig `json:"services"`
}

type ServiceConfig struct {
	Name string `json:"name"`
	Addr string `json:"addr"`
}
type MysqlConfig struct {
	DSN string `json:"dsn"`
}

type RedisConfig struct {
	Cluster  bool     `json:"cluster"`
	Addr     []string `json:"addr"`
	Password string   `json:"password"`
	Database int      `json:"database"`
}
type RabbitmqConfig struct {
	Addrs []string  `json:"addrs"`
	Queue QueueName `json:"queue"`
}
type AlertConfig struct {
	AppId string `json:"appId"`
	URL   string `json:"url"`
}
type ServicesConfig struct {
	AccountService ServiceConfig `json:"account_service"`
	UserService    ServiceConfig `json:"user_service"`
	MsgService     ServiceConfig `json:"msg_service"`
	SettleService  ServiceConfig `json:"settle_service"`
}


func initClient() {
	once.Do(func() {
		plan := mustParse(`{"type":"key","key":"` + consulKey + `"}`)
		// 环境变量只有在容器里有，如果没有获取到手动赋值 方便本地开发调试
		plan.Token = os.Getenv("POD_CONSUL_TOKEN")
		if plan.Token == "" {

		}
		addr := os.Getenv("POD_CONSUL_URL")
		if addr == "" {
			// addr = "http://consul.default:8500"
		}

		KVGetConfig(addr, plan.Token)
		plan.Handler = func(idx uint64, result interface{}) {
			if entries, ok := result.(*api.KVPair); ok {
				// log.Println("index", idx, "value", string(entries.Value))
				co := GlobalConfig{}
				// 解析json
				err := json.Unmarshal(entries.Value, &co)
				if err != nil {
					log.Println("index", idx, "value", string(entries.Value), "error", "json解析错误")
					return
				}
				setGlobalConfig(&co)
			} else {
				log.Println("error", "类型不为*api.KVPair")
			}
		}

		go func() {
			log.Println("consul error", plan.Run(addr))
		}()
	})
}

func KVGetConfig(addr string, token string) {
	// 先获取一次
	conf := api.DefaultConfig()
	conf.Address = addr
	conf.Token = token
	c, err := api.NewClient(conf)
	if err != nil {
		panic("连接consul出错")
	}

	kv, _, err := c.KV().Get(consulKey, nil)
	if err != nil {
		panic("consul获取KV失败-" + err.Error())
	}

	co := GlobalConfig{}
	// 解析json
	err = json.Unmarshal(kv.Value, &co)
	if err != nil {
		panic("配置文件json解析错误" + err.Error())
	}
	setGlobalConfig(&co)
}

func loadValue() *GlobalConfig {
	initClient()
	return getGlobalConfig()
}

func mustParse(q string) *watch.Plan {
	params := makeParams(q)
	plan, err := watch.Parse(params)
	if err != nil {
		log.Fatalf("plan err: %v", err)
	}
	return plan
}

func makeParams(s string) map[string]interface{} {
	var out map[string]interface{}
	dec := json.NewDecoder(bytes.NewReader([]byte(s)))
	if err := dec.Decode(&out); err != nil {
		log.Fatalf("err: %v", err)
	}
	return out
}

func InitConfig() {
	// 用于本地调试
	token := os.Getenv("POD_CONSUL_TOKEN")
	if token == "" {
		setGlobalConfig(&local_config)
		return
	}
	consulKey = os.Getenv("POD_CONSUL_KEY")
	if consulKey == "" {
		log.Fatal("POD_CONSUL_KEY ")
	}
	loadValue()
}

func GetGlobalConfig() *GlobalConfig {
	return getGlobalConfig()
}

type QueueName struct {
	Withdraw struct {
		Exchange   string `json:"exchange"`
		RoutingKey string `json:"routing_key"`
	} `json:"withdraw"`
	WithdrawQuery struct {
		Exchange   string `json:"exchange"`
		RoutingKey string `json:"routing_key"`
	} `json:"withdraw_query"`
	Bonus struct {
		Exchange   string `json:"exchange"`
		RoutingKey string `json:"routing_key"`
	} `json:"bonus"`
	BonusExtend struct {
		Exchange   string `json:"exchange"`
		RoutingKey string `json:"routing_key"`
	} `json:"bonus_extend"`
}
