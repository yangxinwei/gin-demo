package market

import (
	"github.com/yangxinwei/gin-demo-api/client/mysql"
	"github.com/yangxinwei/gin-demo-api/models/entity"
)

func GetStockInfo(symbol string) (entity.FbStock, error) {
	var (
		quote entity.FbStock
		err   error
	)

	_, err = mysql.GetEngine().Where("symbol=?", symbol).Get(&quote)

	if err != nil {
		return quote, err
	} else {
		return quote, nil
	}
}

func GetUSStockLatestQuote(symbol string) (pre, regular, post entity.FbStockLatestQuote, err error) {
	var (
		quotes []entity.FbStockLatestQuote
	)
	err = mysql.GetEngine().Where("symbol=?", symbol).Find(&quotes)

	if err != nil {

	} else {
		for _, quote := range quotes {
			if quote.Section == 1 {
				pre = quote
			} else if quote.Section == 2 {
				regular = quote
			} else if quote.Section == 3 {
				post = quote
			}
		}
	}

	return
}
