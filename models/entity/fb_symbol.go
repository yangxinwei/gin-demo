package entity

type FbSymbol struct {
	Id          int    `xorm:"not null pk autoincr INT(11)"`
	Symbol      string `xorm:"comment('代码') index VARCHAR(50)"`
	SymbolIb    string `xorm:"comment('ib代码') VARCHAR(50)"`
	SecType     string `xorm:"comment('STK股票，INX指数') VARCHAR(10)"`
	Name        string `xorm:"comment('股票名称') VARCHAR(200)"`
	NameCn      string `xorm:"comment('简体中文') VARCHAR(200)"`
	NameTw      string `xorm:"comment('繁体中文') VARCHAR(200)"`
	Market      int    `xorm:"comment('市场,CN:0，HK:1, US:2') VARCHAR(30)"`
	Exchange    string `xorm:"comment('交易所，NYSE') VARCHAR(30)"`
	Close       string `xorm:"comment('最近一个交易日的收盘价') DECIMAL(30,4)"`
	TradeHours  string `xorm:"VARCHAR(300)"`
	Status      int    `xorm:"default 0 comment('0正常交易 1 IPO之前 新股  2 停牌 3 退市') TINYINT(4)"`
	Pinyin      string `xorm:"comment('拼音') VARCHAR(100)"`
	PinyinShort string `xorm:"comment('拼音首字母') VARCHAR(50)"`
}
