package types

type IndicesResp struct {
	Response
	Data []SymbolObjectWithQuote `json:"data"`
}

type HotStocksResp struct {
	Response
	Data []SymbolObjectWithQuote `json:"data"`
}
type ChineseResp struct {
	Response
	Data []SymbolObjectWithQuote `json:"data"`
}

type IndexResp struct {
	Response
	Data SymbolObjectWithQuote `json:"data"`
}

type IndexStocksResp struct {
	Response
	Data SymbolObjectWithQuote `json:"data"`
}
