package config

import "github.com/tal-tech/go-zero/rest"

type Config struct {
	rest.RestConf
	Debug   bool
	Db      Db
	Service Service
}

type Db struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
	Prefix   string
	Charset  string
}

type Service struct {
	Alipay Alipay
	Wechat Wechat
}

type INotify struct {
	App string
	Wap string
}

type Alipay struct {
	AppPrivateKey     string
	AppPublicKey      string
	EncryptContent    string
	PublicKey         string
	AppId             string
	SellerId          string
	SellerName        string
	SandboxPublicKey  string
	SandboxAppId      string
	SandboxSellerId   string
	SandboxSellerName string
	Notify            INotify
}
type Wechat struct {
	NNAppAppId     string
	NNAppAppSecret string
	MiniAppId      string
	PublicAppId    string
	OpenAppId      string
	MchId          string
	MchAPIKey      string
	Notify         INotify
}
