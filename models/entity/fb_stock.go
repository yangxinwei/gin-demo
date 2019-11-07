package entity

type FbStock struct {
	Id              int    `xorm:"not null pk autoincr INT(11)"`
	Symbol          string `xorm:"comment('股票代码') VARCHAR(50)"`
	Name            string `xorm:"comment('股票名称') VARCHAR(200)"`
	NameCn          string `xorm:"comment('简体中文') VARCHAR(200)"`
	NameTw          string `xorm:"comment('繁体中文') VARCHAR(200)"`
	Market          int64  `xorm:"comment('市场') TINYINT(4)"`
	TradeStatus     string `xorm:"comment('股票交易状态') VARCHAR(10)"`
	Open            string `xorm:"comment('开盘价') DECIMAL(30,4)"`
	Close           string `xorm:"comment('收盘价') DECIMAL(30,4)"`
	High            string `xorm:"DECIMAL(30,4)"`
	Low             string `xorm:"DECIMAL(30,4)"`
	PreClose        string `xorm:"comment('前一交易日收盘价') DECIMAL(30,4)"`
	BidPrice        string `xorm:"comment('最新买价') DECIMAL(30,4)"`
	BidSize         int64  `xorm:"INT(10)"`
	AskPrice        string `xorm:"comment('最新卖价') DECIMAL(30,4)"`
	AskSize         int64  `xorm:"INT(10)"`
	Volume          int64  `xorm:"comment('当日成交量') BIGINT(11)"`
	LastPrice       string `xorm:"DECIMAL(30,4)"`
	LastSize        int64  `xorm:"INT(10)"`
	FirstTradeDate  int64  `xorm:"comment('上市日期') INT(11)"`
	Scale           int64  `xorm:"comment('交易精度') INT(11)"`
	Amount          string `xorm:"comment('成交额') DECIMAL(30,4)"`
	TurnoverRate    string `xorm:"comment('换手率') DECIMAL(20,3)"`
	LastTradeDate   int64  `xorm:"comment('最后交易日期') INT(11)"`
	SharesOut       int64  `xorm:"comment('总股本') BIGINT(20)"`
	MarketCap       string `xorm:"comment('总市值') DECIMAL(30,3)"`
	FloatShares     int64  `xorm:"comment('流通股本数量') BIGINT(20)"`
	FloatCap        string `xorm:"comment('流通市值') DECIMAL(30,3)"`
	TrailingPe      string `xorm:"comment('市盈率') DECIMAL(20,6)"`
	ForwardPe       string `xorm:"comment('预测市盈率') DECIMAL(20,6)"`
	TrailingEps     string `xorm:"comment('每股收益') DECIMAL(20,3)"`
	ForwardEps      string `xorm:"comment('预测每股收益') DECIMAL(20,3)"`
	PriceToBook     string `xorm:"comment('市净率') DECIMAL(20,6)"`
	Beta            string `xorm:"comment('beta系数') DECIMAL(10,3)"`
	EnterpriseValue string `xorm:"comment('企业估值') DECIMAL(30,4)"`
	ProfitMargins   string `xorm:"DECIMAL(20,3)"`
	WkHigh13        string `xorm:"DECIMAL(20,4)"`
	WkLow13         string `xorm:"DECIMAL(20,4)"`
	WkHigh26        string `xorm:"DECIMAL(20,4)"`
	WkLow26         string `xorm:"DECIMAL(20,4)"`
	WkHigh52        string `xorm:"DECIMAL(20,4)"`
	WkLow52         string `xorm:"DECIMAL(20,4)"`
	UpdateAt        int64  `xorm:"comment('更新日期') INT(11)"`
}
