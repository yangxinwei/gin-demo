package types

type SortOrderBy int

const (
	SortOrderByDefault    SortOrderBy = iota
	SortOrderByPriceAsc
	SortOrderByChangeAsc
	SortOrderByChangeDesc
)

//添加股票
type FavoritesAddReq struct {
	UserRequest
	Stock string `json:"symbol"`
}
type FavoritesAddResp struct {
	Response
	Data struct {
	} `json:"data"`
}

//移除股票
type FavoritesRemoveReq struct {
	UserRequest
	Stock string `json:"symbol"`
}
type FavoritesRemoveResp struct {
	Response
	Data struct {
	} `json:"data"`
}

//自定义排序
type FavoritesSortReq struct {
	UserRequest
	Stocks string `json:"symbols"`
}
type FavoritesSortResp struct {
	Response
	Data struct {
	} `json:"data"`
}

//排序设置
type FavoritesSortSetReq struct {
	UserRequest
	OrderBy SortOrderBy `json:"orderBy"`
}
type FavoritesSortSetResp struct {
	Response
	Data struct {
	} `json:"data"`
}

//股票列表
type FavoritesListReq struct {
	UserRequest
}
type FavoritesListResp struct {
	Response
	Data struct {
		Stocks []SymbolObjectWithQuote `json:"symbols"`
	} `json:"data"`
}

//报价
type FavoritesQuoteReq struct {
	UserRequest
	Stocks string `json:"symbols"`
}
type FavoritesQuoteResp struct {
	Response
	Data struct {
		Quotes []QuoteData `json:"quotes"`
	} `json:"data"`
}
