package types

type StockProfileResp struct {
	Response
	Data struct {
		Sector   Category      `json:"sector"`
		Industry Category      `json:"industry"`
		Company  CompanyObject `json:"company"`
		//主要股东
		Shareholder []ShareholderObject `json:"shareholder"`
		//股东更新时间
		ShareholderUpdateAt int64 `json:"shareholderUpdateAt"`
	} `json:"data"`
}

type CompanyObject struct {
	//国家
	Country Category `json:"country"`
	//城市
	City Category `json:"city"`
	//公司中文名
	CompanyNameCN string `json:"companyNameCN"`
	//公司英文名
	CompanyNameEN string `json:"companyNameEN"`
	//公司简介,经营业务介绍
	BusinessSummary string `json:"businessSummary"`
	//网址
	Website string `json:"website"`
	//
	Address1 string `json:"address1"`
	//
	Address2 string `json:"address2"`
	//邮编
	PostalCode string `json:"postalCode"`
}

//股东持股
type ShareholderObject struct {
	//股东名
	Name string `json:"name"`
	//持股数
	HoldingQuantity float64 `json:"holdingQuantity"`
	//占比
	Proportion float64 `json:"proportion"`
	//较上期变动数
	Change int64 `json:"change"`
	//较上期变动百分比
	ChangePercent string `json:"changePercent"`
}

type GatheredCompanyResult struct {
	Address      string `json:"address"`
	Code         string `json:"code"`
	Introduction string `json:"introduction"`
	Name         string `json:"name"`
	Website      string `json:"website"`
}
type GatheredCompanyResponse struct {
	ErrorCode int                   `json:"errorCode"`
	Result    GatheredCompanyResult `json:"result"`
}
