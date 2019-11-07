package types

type SearchReq struct {
	Keywords string `json:"keywords"`
}

type SearchRespData struct {
	SymbolObject
	//是否已加入收藏，1-已加入收藏
	Favorite int8 `json:"favorite"`
}
type SearchResp struct {
	Response
	Data struct{
		Result []SearchRespData `json:"result"`
		Keywords string `json:"keywords"`
	} `json:"data"`
}
