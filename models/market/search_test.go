package market

import (
	"testing"
	"fmt"
	"github.com/yangxinwei/gin-demo-api/config"
)

func TestSearch(t *testing.T) {
	config.InitConfig()

	//keywords := `%yr,,d&b,b*ff%`
	keywords := `Yiren`
	if result, err := Search(keywords); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}

func TestFilterWords(t *testing.T) {
	keywords:=`%yr,,d&b,b*ff%`

	keywords = FilterWords(keywords)

	fmt.Println(keywords)
	//output:
}