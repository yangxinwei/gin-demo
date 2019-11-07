package types

type MarketStatusResp struct {
	Response
	Data struct {
		//市场代码
		Market string `json:"market"`
		//市场状态（未开盘 1，交易中 2，休市 3，盘前 4，盘后 5）
		Status string `json:"status"`
		//最近开盘、交易时间,时间戳
		OpenTime string `json:"openTime"`
		//服务器时间,时间戳
		ServerTime int64 `json:"serverTime"`
	} `json:"data"`
}
