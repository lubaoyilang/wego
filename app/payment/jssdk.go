package payment

import (
	"strings"

	"github.com/godcong/wego/config"
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/util"
)

type JSSDK struct {
	config.Config
	*Payment
}

func newJSSDK(p *Payment) *JSSDK {
	return &JSSDK{
		Config:  defaultConfig,
		Payment: p,
	}
}

func NewJSSDK() *JSSDK {
	return newJSSDK(payment)
}

func (j *JSSDK) getUrl() string {
	return core.GetServerIp()
}

func (j *JSSDK) BridgeConfig(pid string) util.Map {
	appid := j.Config.Get("sub_appid")
	if appid == "" {
		appid = j.Config.Get("app_id")
	}

	m := util.Map{
		"appId":     appid,
		"timeStamp": util.Time(),
		"nonceStr":  util.GenerateNonceStr(),
		"package":   strings.Join([]string{"prepay_id", pid}, "="),
		"signType":  "MD5",
	}

	m.Set("paySign", core.GenerateSignature(m, j.Config.Get("key"), core.MakeSignMD5))

	return m
}

func (j *JSSDK) SdkConfig(pid string) util.Map {
	config := j.BridgeConfig(pid)

	config.Set("timestamp", config.Get("timeStamp"))
	config.Delete("timeStamp")

	return config
}

func (j *JSSDK) AppConfig(pid string) util.Map {
	m := util.Map{
		"appid":     j.Config.Get("app_id"),
		"partnerid": j.Config.Get("mch_id"),
		"prepayid":  pid,
		"noncestr":  util.GenerateNonceStr(),
		"timestamp": util.Time(),
		"package":   "Sign=WXPay",
	}

	m.Set("sign", core.GenerateSignature(m, j.Config.Get("aes_key"), core.MakeSignMD5))
	return m
}

func (j *JSSDK) ShareAddressConfig(accessToken interface{}) util.Map {
	token := ""
	switch v := accessToken.(type) {
	case core.AccessToken:
		token0 := v.GetToken()
		token = token0.ToJson()
	case string:
		token = accessToken.(string)
	}
	m := util.Map{
		"appId":     j.Config.Get("app_id"),
		"scope":     "jsapi_address",
		"timeStamp": util.Time(),
		"nonceStr":  util.GenerateNonceStr(),
		"signType":  "SHA1",
	}

	sm := util.Map{
		"appid":       m.Get("appId"),
		"url":         j.getUrl(),
		"timestamp":   m.Get("timeStamp"),
		"noncestr":    m.Get("nonceStr"),
		"accesstoken": token,
	}

	m.Set("addrSign", util.SHA1(sm.UrlEncode()))

	return m
}