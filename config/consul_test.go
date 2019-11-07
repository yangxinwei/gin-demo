package config

import (
	"testing"
	"time"
)

func TestGetGlobalConfig(t *testing.T) {
	conf := GetGlobalConfig()
	t.Logf("%+v\n", conf)
}

func TestInitConfig(t *testing.T) {
	InitConfig()
	println(GetGlobalConfig().Mysql.DSN)
	time.Sleep(time.Second*5)
}

func TestKVGetConfig(t *testing.T) {
	KVGetConfig("", "")
	t.Log(c)
}
