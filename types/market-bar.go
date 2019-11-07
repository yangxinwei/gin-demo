package types

type BarData struct {
	//股票代码
	Symbol string `json:"symbol"`
	//时间戳
	Time int64 `json:"time"`
	//开盘价
	Open float64 `json:"open"`
	//最高价
	High float64 `json:"high"`
	//最低价
	Low float64 `json:"low"`
	//收盘价
	Close float64 `json:"close"`
	//加权均价
	WAP float64 `json:"wap"`
	//成交量
	Volume int64 `json:"volume"`
	//前日收盘价
	PreClose float64 `json:"preClose"`
	//涨跌额
	Change float64 `json:"change"`
	//涨跌幅百分比值
	ChangePercent float64 `json:"changePercent"`
	//5日均价
	MA5 float64 `json:"ma5"`
	//10日均价
	MA10 float64 `json:"ma10"`
	//20日均价
	MA20 float64 `json:"ma20"`
	//5日均量
	VMA5 float64 `json:"vma5"`
	//10日均量
	VMA10 float64 `json:"vma10"`
	//20日均量
	VMA20 float64 `json:"vma20"`
	//换手率百分比
	Turnover float64 `json:"turnover"`
}

type BarResp struct {
	Response
	Data []BarData `json:"data"`
}
