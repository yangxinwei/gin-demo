package utils

//00 系统错误
//01 用户相关
//02 交易相关
//03 行情相关
//04 资讯相关

//1 不能为空
//2 不能大于多少个字符
//3 不能小于多少个字符
//4 格式不正确
//5

var status = map[string]string{
	"200":       "成功",
	"100000001": "服务异常",
	"100000002": "未知错误",
}

func Status(code string) string {
	if msg, ok := status[code]; ok {
		return msg
	} else {
		return "未知错误"
	}
}
