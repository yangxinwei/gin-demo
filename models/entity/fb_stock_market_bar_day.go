package entity

import (
	"time"
)

type FbStockMarketBarDay struct {
	Id            int       `xorm:"not null pk autoincr INT(11)"`
	Symbol        string    `xorm:"VARCHAR(30)"`
	Market        int       `xorm:"INT(11)"`
	Day           time.Time `xorm:"comment('交易日') DATE"`
	Time          int64     `xorm:"comment('交易日时间戳') INT(11)"`
	Open          string    `xorm:"comment('开盘价') DECIMAL(30,4)"`
	Close         string    `xorm:"comment('收盘价') DECIMAL(30,4)"`
	High          string    `xorm:"comment('开盘价') DECIMAL(30,4)"`
	Low           string    `xorm:"comment('最低价') DECIMAL(30,4)"`
	Wap           string    `xorm:"comment('均价') DECIMAL(30,4)"`
	Preclose      string    `xorm:"comment('前一交易日收盘价') DECIMAL(30,4)"`
	Volume        int64     `xorm:"comment('成交量') BIGINT(11)"`
	Count         int64     `xorm:"comment('盘内成交量') INT(11)"`
	Amount        string    `xorm:"comment('成交额') DECIMAL(30,4)"`
	Change        string    `xorm:"comment('价格变动') DECIMAL(20,4)"`
	ChangePercent string    `xorm:"comment('涨跌幅') DECIMAL(20,3)"`
	Ma5           string    `xorm:"comment('5日均价') DECIMAL(30,4)"`
	Ma10          string    `xorm:"comment('10日均价') DECIMAL(30,4)"`
	Ma20          string    `xorm:"comment('20日均价') DECIMAL(30,4)"`
	VMa5          int       `xorm:"comment('5日均量') INT(11)"`
	VMa10         int       `xorm:"comment('10日均量') INT(11)"`
	VMa20         int       `xorm:"comment('20日均量') INT(11)"`
	Turnover      string    `xorm:"comment('换手率') DECIMAL(6,3)"`
}
