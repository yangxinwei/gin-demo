package types

type StockResp struct {
	Response
	Data struct {
		Symbol      SymbolObject    `json:"symbol"`
		Quote       QuoteData       `json:"quote"`
		Fundamental FundamentalData `json:"fundamental"`
		//是否已加入收藏，1-已加入收藏
		Favorite int8 `json:"favorite"`
	} `json:"data"`
}
type StockStatistics struct {
}
