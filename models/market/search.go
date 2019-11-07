package market

import (
	"github.com/yangxinwei/gin-demo-api/client/mysql"
	"github.com/go-xorm/builder"
	"strings"
	"log"
	"github.com/yangxinwei/gin-demo-api/models/entity"
)

//过滤特殊字符
func FilterWords(keyword string) string {
	replacer := strings.NewReplacer(
		`!`, ``, `@`, ``, `#`, ``, `$`, ``, `%`, ``,
		`^`, ``, `&`, ``, `*`, ``, `(`, ``, `)`, ``,
		`;`, ``, `'`, ``, `"`, ``, `,`, ` `, `.`, ``,
	)
	keyword = replacer.Replace(keyword)
	keyword = strings.TrimSpace(keyword)
	return keyword
}
func Search(keywords string) ([]entity.FbSymbol, error) {
	var (
		symbolBeans = make([]entity.FbSymbol, 0)
		err         error
	)

	keywords = FilterWords(keywords)
	log.Println(keywords)
	//keyword = strings.Replace(keyword, ",", " ", -1)
	//keywords := strings.Split(keyword, " ")
	//var symbolCond = builder.Like{}
	//var nameCond = builder.Like{}
	//var nmaeCNcond = builder.Like{}
	//for i := 0; i < len(keywords); i++ {
	symbolCond := builder.Like{"symbol", keywords + "%"}
	nameCond := builder.Like{"name", keywords + "%"}
	nmaeCNcond := builder.Like{"name_cn", keywords + "%"}
	pinyincond := builder.Like{"pinyin", keywords + "%"}
	pinyinScond := builder.Like{"pinyin_short", keywords + "%"}
	//}
	err = mysql.GetEngine().Where(symbolCond.Or(nameCond).Or(nmaeCNcond).Or(pinyincond).Or(pinyinScond)).Limit(10).Find(&symbolBeans)

	if err != nil {
		return nil, err
	} else {
		return symbolBeans, nil
	}
}

func SearchBySymbol(symbol string) (entity.FbSymbol, error) {
	var (
		symbolBean entity.FbSymbol
		err        error
		//ok         bool
	)

	_, err = mysql.GetEngine().Where("symbol=?", symbol).Get(&symbolBean)

	if err != nil {
		return symbolBean, err
	} else {
		return symbolBean, nil
	}
}
