package types

//最新报价、成交量、涨跌
type QuoteData struct {
	//股票代码
	Symbol string `json:"symbol"`
	//最新成交时间
	Time int64 `json:"time"`
	//开盘价
	Open float64 `json:"open"`
	//最高价
	High float64 `json:"high"`
	//最低价
	Low float64 `json:"low"`
	//收盘价
	Close float64 `json:"close"`
	//前一交易日收盘价
	PreClose float64 `json:"preClose"`
	//最新价
	Price float64 `json:"Price"`
	//最新成交数量
	LatestSize int64 `json:"latestSize"`
	//卖盘价
	AskPrice float64 `json:"askPrice"`
	//卖盘数量
	AskSize int64 `json:"askSize"`
	//买盘价
	BidPrice float64 `json:"bidPrice"`
	//买盘数量
	BidSize int64 `json:"bidSize"`
	//成交量
	Volume int64 `json:"volume"`
	//成交额
	Amount float64 `json:"amount"`
	//换手率百分比
	TurnoverRate float64 `json:"turnoverRate"`
	//交易状态, NORMAL-正常,DELIST-退市,HALTED-停牌,NEW-新股
	TradeStatus string `json:"tradeStatus"`
	//涨跌额
	Change float64 `json:"change"`
	//涨跌幅
	ChangePercent float64 `json:"changePercent"`
	//市场状态 PREPRE-盘前之前，PRE-盘前，REGULAR-盘中，POST-盘后，POSTPOST-盘后之后
	MarketStatus string `json:"marketStatus"`
	//盘前数据
	PreMarket MarketQuoteData `json:"preMarket"`
	//盘后数据
	PostMarket MarketQuoteData `json:"postMarket"`
	//盘中数据
	RegularMarket MarketQuoteData `json:"regularMarket"`
}

//行情数据
type MarketQuoteData struct {
	//更新时间
	Time int64 `json:"time"`
	//最新价
	Price float64 `json:"price"`
	//成交量
	Volume int64 `json:"volume"`
	//涨跌额
	Change float64 `json:"change"`
	//涨跌幅
	ChangePercent float64 `json:"changePercent"`
	//数据来源，DELAYED- /REALTIME
	//Source string `json:"source"`
}
type QuotesResp struct {
	Response
	Data []QuoteData `json:"data"`
}

type QuoteResp struct {
	Response
	Data QuoteData `json:"data"`
}
