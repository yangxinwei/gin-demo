package types

type NewsResp struct {
	Response
	Data []NewsAbstract `json:"data"`
}

//新闻摘要
type NewsAbstract struct {
	//标题
	Title string `json:"title"`
	//发布时间
	Time int64 `json:"time"`
	//新闻源
	Source string `json:"source"`
	//摘要
	abstract string `json:"abstract"`
	//来源地址
	SourceUrl string `json:"sourceUrl"`
	//内容地址
	Url string `json:"url"`
}

type NewsDetail struct {
	NewsAbstract
	//新闻内容
	Content string `json:"content"`
}
