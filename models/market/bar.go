package market

import (
	"github.com/yangxinwei/gin-demo-api/client/mysql"
	"github.com/yangxinwei/gin-demo-api/models/entity"
	"github.com/go-xorm/builder"
)

func GetDayBar(market, symbol string, begin, end int64) ([]*entity.FbStockMarketBarDay, error) {
	var (
		symbolBean []*entity.FbStockMarketBarDay
		err        error
	)

	sqlBuilder := builder.Eq{"symbol": symbol}
	if begin > 0 {
		sqlBuilder.And(builder.Gt{"time": begin})
	}
	if end > 0 {
		sqlBuilder.And(builder.Lt{"time": end})
	}
	err = mysql.GetEngine().Desc("time").Where(sqlBuilder).Find(&symbolBean)

	if err != nil {
		return symbolBean, err
	} else {
		return symbolBean, nil
	}
}
