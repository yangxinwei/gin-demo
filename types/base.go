package types

type APIError struct {
	//状态码
	Status string `json:"errorCode"`
	//状态信息
	Message string `json:"msg"`
	//成功：success、失败：fail
	Result string `json:"result"`
}
type Request struct {
	//设备id
	DeviceId string `json:"deviceId,omitempty"`
	//SessionId
	SessionId string `json:"sessionId,omitempty"`
}
type UserRequest struct {
	Request
	//用户id
	UserId string `json:"userId,omitempty"`
}
type Response struct {
	//状态码
	Status string `json:"errorCode"`
	//状态信息
	Message string `json:"msg"`
	//成功：success、失败：fail
	Result string `json:"result"`
	//毫秒时间戳
	Time int64 `json:"time"`
	//reqId   int64  `json:"reqId"`
}

type CategoryMarket string
type CategorySecType string

const (
	CategoryMarketUS = "US"
	CategoryMarketHK = "HK"
	CategoryMarketCN = "CN"

	CategorySecTypeStock = "STK"
	CategorySecTypeIndex = "IDX"
)

const (
	MarketStatusPrePre   = "PREPRE"
	MarketStatusPre      = "PRE"
	MarketStatusRegular  = "REGULAR"
	MarketStatusPost     = "POST"
	MarketStatusPostPost = "POSTPOST"
)

type Category struct {
	Name   string    `json:"name"`
	Code   string    `json:"code"`
	Parent *Category `json:"parent"`
}
