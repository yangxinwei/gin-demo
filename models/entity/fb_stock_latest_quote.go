package entity

type FbStockLatestQuote struct {
	Id            int    `xorm:"not null pk autoincr INT(11)"`
	Symbol        string `xorm:"VARCHAR(30)"`
	Section       int64  `xorm:"comment('1-盘前,2-盘中,3-盘后') TINYINT(10)"`
	Time          int64  `xorm:"INT(11)"`
	Price         string `xorm:"DECIMAL(30,4)"`
	Volume        int64  `xorm:"BIGINT(20)"`
	Change        string `xorm:"DECIMAL(20,4)"`
	ChangePercent string `xorm:"DECIMAL(20,3)"`
}
