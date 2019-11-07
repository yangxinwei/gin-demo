package market

import (
	"testing"
	"log"
	"os"
	"github.com/yangxinwei/gin-demo-api/config"
)

func TestGetUSStockLatestQuote(t *testing.T) {
	setTestEnv()
	pre, regular, post, err := GetUSStockLatestQuote("yrd")
	if err != nil {

	}
	log.Printf("pre:%+v", pre)
	log.Printf("regular:%+v", regular)
	log.Printf("post:%+v", post)
}
func setTestEnv() {
	os.Setenv("POD_CONSUL_URL", "")
	os.Setenv("POD_CONSUL_KEY", "")
	os.Setenv("POD_CONSUL_TOKEN", "")
	config.InitConfig()
}
