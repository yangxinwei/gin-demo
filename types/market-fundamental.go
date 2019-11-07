package types

//个股基本面数据
type FundamentalData struct {
	//总股本
	SharesOut int64 `json:"sharesOut"`
	//总市值
	MarketCap float64 `json:"marketCap"`
	//流通股本数量
	FloatShares int64 `json:"floatShares"`
	//流通市值
	FloatCap float64 `json:"floatCap"`
	//Earnings Per Share 每股收益
	TrailingEps float64 `json:"trailingEps"`
	//市盈率
	TrailingPe float64 `json:"trailingPe"`
	//市净率,净值比 每股市价(Price)/每股净资产(Book Value)
	PriceToBook float64 `json:"priceToBook"`
	//52WKHIGH
	WkHigh52 float64 `json:"wkHigh52"`
	//52WKLOW
	WkLow52 float64 `json:"wkLow52"`
	//26WKHIGH
	//WKHigh26 float64 `json:"wkHigh26"`
	//26WKLOW
	//WKLow26 float64 `json:"wkLow26"`
	//13WKHIGH
	//WKHigh13 float64 `json:"wkHigh13"`
	//13WKLOW
	//WKLow13 float64 `json:"wkLow13"`
	//每股帳面價值
	//BookValuePerShare float64 `json:"bookValuePerShare"`
	//预测每股收益
	//ForwardEps float64 `json:"forwardEps"`
	//预测市盈率
	//ForwardPE float64 `json:"forwardPE"`

	//updated
	//最近一次发财报日期
	//FinanceReportTime string `json:"financeReportTime"`
	//beta系数
	//Beta float64 `json:"beta"`
	//企业估值
	//EnterpriseValue int64 `json:"enterpriseValue"`
	//利润率
	//ProfitMargins  float64 `json:"profitMargins"`
	//sharesShort
	//shortRatio
	//priceHint
	//totalAssets 总资产
	//pegRatio 市盈率与增长比率
	//派息率 Payout Ratio
}

type FundamentalResp struct {
	Response
	Data FundamentalData `json:"data"`
}
