package pay

// 支付类型
const (
	TypeWechatApp = iota + 1
	TypeWechatWap
	TypeAlipayApp
	TypeAlipayWap
)

// 第三方支付商名称
const (
	ServiceAlipay = "Alipay"
	ServiceWechat = "Wechat"
)

// 支付操作
const (
	MethodApp   = "App"
	MethodWap   = "Wap"
	MethodQuery = "Query"
)
