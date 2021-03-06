package payment

import (
	"strconv"

	"github.com/godcong/wego/core"
	"github.com/godcong/wego/util"
)

/*RedPack RedPack */
type RedPack struct {
	*Payment
}

func newRedPack(payment *Payment) interface{} {
	return &RedPack{
		Payment: payment,
	}
}

//NewRedPack NewRedPack
func NewRedPack(config *core.Config) *RedPack {
	return newRedPack(NewPayment(config)).(*RedPack)
}

/*Info 查询红包记录
接口调用请求说明
请求Url	https://api.mch.weixin.qq.com/mmpaymkttransfers/gethbinfo
是否需要证书	是（证书及使用说明详见商户证书）
请求方式	POST
请求参数
字段名	字段	必填	示例值	类型	说明
商户订单号	mch_billno	是	10000098201411111234567890	String(28)	商户发放红包的商户订单号
*/
func (r *RedPack) Info(mchBillNo string) core.Responder {
	m := util.Map{
		"mch_billno": mchBillNo,
		"appid":      "app_id",
		"bill_type":  "MCHT",
	}
	return r.SafeRequest(mmpaymkttransfersGetHbInfo, m)

}

/*SendNormal 发放普通红包
发放规则
1.发送频率限制------默认1800/min
2.发送个数上限------按照默认1800/min算
3.金额限制------默认红包金额为1-200元，如有需要，可前往商户平台进行设置和申请
4.其他其他限制吗？------单个用户可领取红包上线为10个/天，如有需要，可前往商户平台进行设置和申请
5.如果量上满足不了我们的需求，如何提高各个上限？------金额上限和用户当天领取次数上限可以在商户平台进行设置
注意-红包金额大于200或者小于1元时，请求参数scene_id必传，参数说明见下文。
注意2-根据监管要求，新申请商户号使用现金红包需要满足两个条件:1、入驻时间超过90天 2、连续正常交易30天。
注意3-移动应用的appid无法使用红包接口。
消息触达规则
现金红包发放后会以公众号消息的形式触达用户，不同情况下触达消息的形式会有差别，规则如下:
是否关注	关注时间	是否接收消息	触达消息
否	/	/	模版消息
是	<=50小时	是	模版消息
是	<=50小时	否	模版消息
是	>50小时	是	防伪消息
是	>50小时	否	模版消息
接口调用请求说明
请求Url	https://api.mch.weixin.qq.com/mmpaymkttransfers/sendredpack
是否需要证书	是（证书及使用说明详见商户证书）
请求方式	POST
请求参数
字段名	字段	必填	示例值	类型	说明
随机字符串	nonce_str	是	5K8264ILTKCH16CQ2502SI8ZNMTM67VS	String(32)	随机字符串，不长于32位
签名	sign	是	C380BEC2BFD727A4B6845133519F3AD6	String(32)	详见签名生成算法
商户订单号	mch_billno	是	10000098201411111234567890	String(28)
商户订单号（每个订单号必须唯一。取值范围:0~9，a~z，A~Z）
接口根据商户订单号支持重入，如出现超时可再调用。
商户号	mch_id	是	10000098	String(32)	微信支付分配的商户号
子商户号	sub_mch_id	否	10000090	String(32)	微信支付分配的子商户号，服务商模式下必填
公众账号appid	wxappid	是	wx8888888888888888	String(32)	微信分配的公众账号ID（企业号corpid即为此appId）。接口传入的所有appid应该为公众号的appid（在mp.weixin.qq.com申请的），不能为APP的appid（在open.weixin.qq.com申请的）。
触达用户appid	msgappid	是	wx28b16568a629bb33	String(32)	服务商模式下触达用户时的appid(可填服务商自己的appid或子商户的appid)，服务商模式下必填，服务商模式下填入的子商户appid必须在微信支付商户平台中先录入，否则会校验不过。
商户名称	send_name	是	天虹百货	String(32)	红包发送者名称
用户openid	re_openid	是	oxTWIuGaIt6gTKsQRLau2M0yL16E	String(32)
接受红包的用户
用户在wxappid下的openid，服务商模式下可填入msgappid下的openid。
付款金额	total_amount	是	1000	int	付款金额，单位分
红包发放总人数	total_num	是	1	int
红包发放总人数
total_num=1
红包祝福语	wishing	是	感谢您参加猜灯谜活动，祝您元宵节快乐！	String(128)	红包祝福语
Ip地址	client_ip	是	192.168.0.1	String(15)	调用接口的机器Ip地址
活动名称	act_name	是	猜灯谜抢红包活动	String(32)	活动名称
备注	remark	是	猜越多得越多，快来抢！	String(256)	备注信息
场景id	scene_id	否	PRODUCT_8	String(32)
发放红包使用场景，红包金额大于200或者小于1元时必传
PRODUCT_1:商品促销
PRODUCT_2:抽奖
PRODUCT_3:虚拟物品兑奖
PRODUCT_4:企业内部福利
PRODUCT_5:渠道分润
PRODUCT_6:保险回馈
PRODUCT_7:彩票派奖
PRODUCT_8:税务刮奖
活动信息	risk_info	否	posttime%3d123123412%26clientversion%3d234134%26mobile%3d122344545%26deviceid%3dIOS	String(128)
posttime:用户操作的时间戳
mobile:业务系统账号的手机号，国家代码-手机号。不需要+号
deviceid :mac 地址或者设备唯一标识
clientversion :用户操作的客户端版本
把值为非空的信息用key=value进行拼接，再进行urlencode
urlencode(posttime=xx& mobile =xx&deviceid=xx)
扣钱方mchid	consume_mch_id	否	10000098	String(32)	常规模式下无效，服务商模式下选填，服务商模式下不填默认扣子商户的钱
*/
func (r *RedPack) SendNormal(m util.Map) core.Responder {
	m.Set("total_num", strconv.Itoa(1))
	m.Set("client_ip", core.GetServerIP())
	m.Set("wxappid", r.Get("app_id"))
	return r.SafeRequest(mmpaymkttransfersSendRedPack, m)
}

/*SendGroup 裂变红包
发放规则
裂变红包:一次可以发放一组红包。首先领取的用户为种子用户，种子用户领取一组红包当中的一个，并可以通过社交分享将剩下的红包给其他用户。裂变红包充分利用了人际传播的优势。
消息触达规则
现金红包发放后会以公众号消息的形式触达用户，不同情况下触达消息的形式会有差别，规则如下:
是否关注	关注时间	是否接收消息	触达消息
否	/	/	模版消息
是	<=50小时	是	模版消息
是	<=50小时	否	模版消息
是	>50小时	是	防伪消息
是	>50小时	否	模版消息
接口调用请求说明
请求Url	https://api.mch.weixin.qq.com/mmpaymkttransfers/sendgroupredpack
是否需要证书	是（证书及使用说明详见商户证书）
请求方式	POST
请求参数
字段名	字段	必填	示例值	类型	说明
随机字符串	nonce_str	是	5K8264ILTKCH16CQ2502SI8ZNMTM67VS	String(32)	随机字符串，不长于32位
签名	sign	是	C380BEC2BFD727A4B6845133519F3AD6	String(32)	详见签名生成算法
商户订单号	mch_billno	是	10000098201411111234567890	String(28)
商户订单号（每个订单号必须唯一）
组成:mch_id+yyyymmdd+10位一天内不能重复的数字。
接口根据商户订单号支持重入，如出现超时可再调用。
商户号	mch_id	是	10000098	String(32)	微信支付分配的商户号
子商户号	sub_mch_id	否	10000090	String(32)	微信支付分配的子商户号，服务商模式下必填
公众账号appid	wxappid	是	wx8888888888888888	String(32)	微信分配的公众账号ID（企业号corpid即为此appId），接口传入的所有appid应该为公众号的appid（在mp.weixin.qq.com申请的），不能为APP的appid（在open.weixin.qq.com申请的）。
触达用户appid	msgappid	否	wx28b16568a629bb33	String(32)	服务商模式下触达用户时的appid(可填服务商自己的appid或子商户的appid)，服务商模式下必填，服务商模式下填入的子商户appid必须在微信支付商户平台中先录入，否则会校验不过。
商户名称	send_name	是	天虹百货	String(32)	红包发送者名称
用户openid	re_openid	是	oxTWIuGaIt6gTKsQRLau2M0yL16E	String(32)
接收红包的种子用户（首个用户）
用户在wxappid下的openid
付款金额	total_amount	是	1000	int	红包发放总金额，即一组红包金额总和，包括分享者的红包和裂变的红包，单位分
红包发放总人数	total_num	是	3	int	红包发放总人数，即总共有多少人可以领到该组红包（包括分享者）
红包金额设置方式	amt_type	是	ALL_RAND	String(32)
红包金额设置方式
ALL_RAND—全部随机,商户指定总金额和红包发放总人数，由微信支付随机计算出各红包金额
红包祝福语	wishing	是	感谢您参加猜灯谜活动，祝您元宵节快乐！	String(128)	红包祝福语
活动名称	act_name	是	猜灯谜抢红包活动	String(32)	活动名称
备注	remark	是	猜越多得越多，快来抢！	String(256)	备注信息
场景id	scene_id	否	PRODUCT_8	String(32)
PRODUCT_1:商品促销
PRODUCT_2:抽奖
PRODUCT_3:虚拟物品兑奖
PRODUCT_4:企业内部福利
PRODUCT_5:渠道分润
PRODUCT_6:保险回馈
PRODUCT_7:彩票派奖
PRODUCT_8:税务刮奖
活动信息	risk_info	否	posttime%3d123123412%26clientversion%3d234134%26mobile%3d122344545%26deviceid%3dIOS	String(128)
posttime:用户操作的时间戳
mobile:业务系统账号的手机号，国家代码-手机号。不需要+号
deviceid :mac 地址或者设备唯一标识
clientversion :用户操作的客户端版本
把值为非空的信息用key=value进行拼接，再进行urlencode
urlencode(posttime=xx& mobile =xx&deviceid=xx)
资金授权商户号	consume_mch_id	否	1222000096	String(32)
资金授权商户号
服务商替特约商户发放时使用
扣钱方mchid	consume_mch_id	否	10000098	String(32)	常规模式下无效，服务商模式下选填，服务商模式下不填默认扣子商户的钱
*/
func (r *RedPack) SendGroup(m util.Map) core.Responder {
	m.Set("amt_type", "ALL_RAND")
	m.Set("wxappid", r.Get("app_id"))
	return r.SafeRequest(mmpaymkttransfersSendGroupRedPack, m)
}
