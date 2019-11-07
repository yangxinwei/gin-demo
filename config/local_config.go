package config

var local_config = GlobalConfig{
	Mysql: struct {
		DSN string `json:"dsn"`
	}{DSN: "futurebull_dev:futurebull_dev_123@tcp(10.141.8.141:3306)/futurebull_dev?charset=utf8mb4&collation=utf8mb4_general_ci"},
	// }{DSN: "zhuanduoduo_dev:zhuanduoduo_dev_123@tcp(10.141.8.141:3306)/zhuanduoduo_dev?charset=utf8mb4&collation=utf8mb4_general_ci"},
	Redis: redisConfig,
		Rabbitmq: struct {
		Addrs []string  `json:"addrs"`
		Queue QueueName `json:"queue"`
	}{Addrs: []string{
		"amqp://admin:123456@localhost:5672/zdd_account",
		// "amqp://zhuandd:zhuandd@10.141.4.87:5673/zhuandd",
		// "amqp://zhuandd:zhuandd@10.141.5.190:36348/zhuandd",
	}, Queue: queueName},
	Alert: struct {
		AppId string `json:"appId"`
		URL   string `json:"url"`
	}{AppId: "", URL: ""},
	Services: struct {
		AccountService ServiceConfig `json:"account_service"`
		UserService    ServiceConfig `json:"user_service"`
		MsgService     ServiceConfig `json:"msg_service"`
		SettleService  ServiceConfig `json:"settle_service"`
	}{AccountService: ServiceConfig{
		Name: "user_service",
		Addr: "localhost:8801",
	}, UserService: ServiceConfig{
		Name: "account_service",
		Addr: "localhost:8802",
	}, MsgService: ServiceConfig{
		Name: "msg_service",
		Addr: "http://msgcenter-service-operations.kpl.yixinonline.org",
	}, SettleService: ServiceConfig{
		Name: "settle_service",
		Addr: "http://10.141.4.207:8180",
	}},
}

var queueName = QueueName{
	Withdraw: struct {
		Exchange   string `json:"exchange"`
		RoutingKey string `json:"routing_key"`
	}{"", "zdd_withdraw"},
	WithdrawQuery: struct {
		Exchange   string `json:"exchange"`
		RoutingKey string `json:"routing_key"`
	}{"", "zdd_withdraw_query"},
	Bonus: struct {
		Exchange   string `json:"exchange"`
		RoutingKey string `json:"routing_key"`
	}{"", "zdd_bonus_notice"},
	BonusExtend: struct {
		Exchange   string `json:"exchange"`
		RoutingKey string `json:"routing_key"`
	}{"", "zdd_bonus_extend"},
}

var redisConfig = RedisConfig{
	Cluster:false,
	Addr:[]string{"127.0.0.1:6379"},
	Database:1,
	Password:"",
}
