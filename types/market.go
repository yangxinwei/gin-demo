package types

import "strings"

type SymbolObject struct {
	//代码
	Symbol string `json:"symbol"`
	//名称
	Name string `json:"name"`
	//国家或地区市场,US,CN,HK
	Market string `json:"market"`
	//发行交易所
	Exchange string `json:"exchange"`
	//所属类型，STK-股票，IND-指数 OPT-期权 FUT-期货 CASH-外汇 NEWS-新闻 FUND-基金
	SecType CategorySecType `json:"type"`
	//交易状态, NORMAL-正常,DELIST-退市,HALTED-停牌,NEW-新股
	Status string `json:"status"`
}

func GetSymbolStatusByIndex(index int) string {
	switch index {
	case 0:
		//
		return "NORMAL"
	case 1:
		//退市
		return "DELIST"
	case 2:
		//停盘
		return "HALTED"
	case 3:
		//新股，未上市
		return "NEW"
	}
	return "NORMAL"
}

func GetMarketIndex(MarketName string) int {
	switch strings.ToUpper(MarketName) {
	case CategoryMarketUS:
		return 2
	case CategoryMarketHK:
		return 1
	case CategoryMarketCN:
		return 0

	}
	return 0
}
func GetMarketString(index int) string {
	switch index {
	case 0:
		return CategoryMarketCN
	case 1:
		return CategoryMarketHK
	case 2:
		return CategoryMarketUS

	}
	return ""
}

type SymbolObjectWithQuote struct {
	SymbolObject
	Quote QuoteData `json:"quote"`
}



type BulletinObject struct {
	//时间
	Time int64 `json:"time"`
	//正文
	Content string `json:"content"`
}

type BulletinsResp struct {
	Response
	Data []BulletinObject `json:"data"`
}
type FinanceResp struct {
	Response
	Data struct {
	} `json:"data"`
}
