package types

type TimeLineData struct {
	//股票代码
	Symbol string `json:"symbol"`
	//时间戳
	Time int64 `json:"time"`
	//截止到当前时间的最新价
	Price float64 `json:"price"`
	//最新收盘价
	PreClose float64 `json:"preClose"`
	//加权均价
	WAP float64 `json:"wap"`
	//当前时间段内的成交量，1分钟，或则5分钟内
	Volume int64 `json:"volume"`
	//涨跌幅百分比值
	ChangePercent float64 `json:"changePercent"`
	//1- 盘前交易, 2-盘中交易,  3-盘后交易
	Session int `json:"session"`
}
type TimeLineResp struct {
	Response
	Data []TimeLineData `json:"data"`
}

