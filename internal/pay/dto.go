package pay

type PayType int8

func (p PayType) String() string {
	switch int(p) {
	case TypeAlipayApp:
		return "支付宝APP"
	case TypeAlipayWap:
		return "支付宝WAP"
	case TypeWechatApp:
		return "微信APP"
	case TypeWechatWap:
		return "微信WAP"
	}

	return "未知"
}
