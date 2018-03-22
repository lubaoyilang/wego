package official_account

import (
	"time"

	"github.com/godcong/wego/core"
)

type DataCube struct {
	core.Config
	*OfficialAccount
}

func newDataCube(officialAccount *OfficialAccount) *DataCube {
	return &DataCube{
		Config:          defaultConfig,
		OfficialAccount: officialAccount,
	}
}

func NewDataCube() *DataCube {
	return newDataCube(account)
}

// 获取用户增减数据（getusersummary）	7	https://api.weixin.qq.com/datacube/getusersummary?access_token=ACCESS_TOKEN
// 成功:
// {"list":[{"ref_date":"2018-03-19","user_source":0,"new_user":0,"cancel_user":0},{"ref_date":"2018-03-19","user_source":17,"new_user":1,"cancel_user":0}]}
// 成功:
// {"list":[{"ref_date":"2018-03-19","user_source":0,"new_user":0,"cancel_user":0},{"ref_date":"2018-03-19","user_source":17,"new_user":1,"cancel_user":0}]}
// 失败:
// {"errcode":61501,"errmsg":"date range error hint: [_muTLA05701504]"}
func (d *DataCube) GetUserSummary(beginDate, endDate time.Time) *core.Response {
	core.Debug("DataCube|GetUserSummary", beginDate, endDate)
	return d.get(
		DATACUBE_GETUSERSUMMARY_URL_SUFFIX,
		beginDate.Format("2006-01-02"),
		endDate.Format("2006-01-02"),
	)
}

// 获取累计用户数据（getusercumulate）	7	https://api.weixin.qq.com/datacube/getusercumulate?access_token=ACCESS_TOKEN
// 成功:
// {"list":[{"ref_date":"2018-03-18","user_source":0,"cumulate_user":5},{"ref_date":"2018-03-19","user_source":0,"cumulate_user":6},{"ref_date":"2018-03-20","user_source":0,"cumulate_user":6}]}
// 失败:
// {"errcode":61501,"errmsg":"date range error hint: [_muTLA05701504]"}
func (d *DataCube) GetUserCumulate(beginDate, endDate time.Time) *core.Response {
	core.Debug("DataCube|GetUserCumulate", beginDate, endDate)
	return d.get(
		DATACUBE_GETUSERCUMULATE_URL_SUFFIX,
		beginDate.Format("2006-01-02"),
		endDate.Format("2006-01-02"),
	)
}

// 获取图文群发每日数据（getarticlesummary）	1	https://api.weixin.qq.com/datacube/getarticlesummary?access_token=ACCESS_TOKEN
// 成功:
// {"list":[]}
// 失败:
// {"errcode":61501,"errmsg":"date range error hint: [_muTLA05701504]"}
func (d *DataCube) GetArticleSummary(beginDate, endDate time.Time) *core.Response {
	core.Debug("DataCube|GetArticleSummary", beginDate, endDate)
	return d.get(
		DATACUBE_GETARTICLESUMMARY_URL_SUFFIX,
		beginDate.Format("2006-01-02"),
		endDate.Format("2006-01-02"),
	)
}

// 获取图文群发总数据（getarticletotal）	1	https://api.weixin.qq.com/datacube/getarticletotal?access_token=ACCESS_TOKEN
// 成功:
// {"list":[]}
// 失败:
// {"errcode":61501,"errmsg":"date range error hint: [_muTLA05701504]"}
func (d *DataCube) GetArticleTotal(beginDate, endDate time.Time) *core.Response {
	core.Debug("DataCube|GetArticleTotal", beginDate, endDate)
	return d.get(
		DATACUBE_GETARTICLETOTAL_URL_SUFFIX,
		beginDate.Format("2006-01-02"),
		endDate.Format("2006-01-02"),
	)
}

// 获取图文统计数据（getuserread）	3	https://api.weixin.qq.com/datacube/getuserread?access_token=ACCESS_TOKEN
// 成功:
// {"list":[]}
// 失败:
// {"errcode":61501,"errmsg":"date range error hint: [_muTLA05701504]"}
func (d *DataCube) GetUserRead(beginDate, endDate time.Time) *core.Response {
	core.Debug("DataCube|GetUserRead", beginDate, endDate)
	return d.get(
		DATACUBE_GETUSERREAD_URL_SUFFIX,
		beginDate.Format("2006-01-02"),
		endDate.Format("2006-01-02"),
	)
}

// 获取图文统计分时数据（getuserreadhour）	1	https://api.weixin.qq.com/datacube/getuserreadhour?access_token=ACCESS_TOKEN
// 成功:
// {"list":[]}
// 失败:
// {"errcode":61501,"errmsg":"date range error hint: [_muTLA05701504]"}
func (d *DataCube) GetUserReadHour(beginDate, endDate time.Time) *core.Response {
	core.Debug("DataCube|GetUserReadHour", beginDate, endDate)
	return d.get(
		DATACUBE_GETUSERREADHOUR_URL_SUFFIX,
		beginDate.Format("2006-01-02"),
		endDate.Format("2006-01-02"),
	)
}

// 获取图文分享转发数据（getusershare）	7	https://api.weixin.qq.com/datacube/getusershare?access_token=ACCESS_TOKEN
// 成功:
// {"list":[]}
// 失败:
// {"errcode":61501,"errmsg":"date range error hint: [_muTLA05701504]"}
func (d *DataCube) GetUserShare(beginDate, endDate time.Time) *core.Response {
	core.Debug("DataCube|GetUserReadHour", beginDate, endDate)
	return d.get(
		DATACUBE_GETUSERSHARE_URL_SUFFIX,
		beginDate.Format("2006-01-02"),
		endDate.Format("2006-01-02"),
	)
}

// 获取图文分享转发分时数据（getusersharehour）	1	https://api.weixin.qq.com/datacube/getusersharehour?access_token=ACCESS_TOKEN
// 失败:
// {"errcode":61501,"errmsg":"date range error hint: [_muTLA05701504]"}
func (d *DataCube) GetUserShareHour(beginDate, endDate time.Time) *core.Response {
	core.Debug("DataCube|GetUserReadHour", beginDate, endDate)
	return d.get(
		DATACUBE_GETUSERSHAREHOUR_URL_SUFFIX,
		beginDate.Format("2006-01-02"),
		endDate.Format("2006-01-02"),
	)
}

// 获取消息发送概况数据（getupstreammsg）	7	https://api.weixin.qq.com/datacube/getupstreammsg?access_token=ACCESS_TOKEN
func (d *DataCube) GetUpstreamMsg(beginDate, endDate time.Time) *core.Response {
	core.Debug("DataCube|GetUpstreamMsg", beginDate, endDate)
	return d.get(
		DATACUBE_GETUPSTREAMMSG_URL_SUFFIX,
		beginDate.Format("2006-01-02"),
		endDate.Format("2006-01-02"),
	)
}

// 获取消息分送分时数据（getupstreammsghour）	1	https://api.weixin.qq.com/datacube/getupstreammsghour?access_token=ACCESS_TOKEN
func (d *DataCube) GetUpstreamMsgHour(beginDate, endDate time.Time) *core.Response {
	core.Debug("DataCube|GetUpstreamMsgHour", beginDate, endDate)
	return d.get(
		DATACUBE_GETUPSTREAMMSGHOUR_URL_SUFFIX,
		beginDate.Format("2006-01-02"),
		endDate.Format("2006-01-02"),
	)
}

// 获取消息发送周数据（getupstreammsgweek）	30	https://api.weixin.qq.com/datacube/getupstreammsgweek?access_token=ACCESS_TOKEN
func (d *DataCube) GetUpstreamMsgWeek(beginDate, endDate time.Time) *core.Response {
	core.Debug("DataCube|GetUpstreamMsgWeek", beginDate, endDate)
	return d.get(
		DATACUBE_GETUPSTREAMMSGWEEK_URL_SUFFIX,
		beginDate.Format("2006-01-02"),
		endDate.Format("2006-01-02"),
	)
}

// 获取消息发送月数据（getupstreammsgmonth）	30	https://api.weixin.qq.com/datacube/getupstreammsgmonth?access_token=ACCESS_TOKEN
func (d *DataCube) GetUpstreamMsgMonth(beginDate, endDate time.Time) *core.Response {
	core.Debug("DataCube|GetUpstreamMsgMonth", beginDate, endDate)
	return d.get(
		DATACUBE_GETUPSTREAMMSGMONTH_URL_SUFFIX,
		beginDate.Format("2006-01-02"),
		endDate.Format("2006-01-02"),
	)
}

// 获取消息发送分布数据（getupstreammsgdist）	15	https://api.weixin.qq.com/datacube/getupstreammsgdist?access_token=ACCESS_TOKEN
func (d *DataCube) GetUpstreamMsgDist(beginDate, endDate time.Time) *core.Response {
	core.Debug("DataCube|GetUpstreamMsgDist", beginDate, endDate)
	return d.get(
		DATACUBE_GETUPSTREAMMSGDIST_URL_SUFFIX,
		beginDate.Format("2006-01-02"),
		endDate.Format("2006-01-02"),
	)
}

// 获取消息发送分布周数据（getupstreammsgdistweek）	30	https://api.weixin.qq.com/datacube/getupstreammsgdistweek?access_token=ACCESS_TOKEN
func (d *DataCube) GetUpstreamMsgDistWeek(beginDate, endDate time.Time) *core.Response {
	core.Debug("DataCube|GetUpstreamMsgDistWeek", beginDate, endDate)
	return d.get(
		DATACUBE_GETUPSTREAMMSGDISTWEEK_URL_SUFFIX,
		beginDate.Format("2006-01-02"),
		endDate.Format("2006-01-02"),
	)
}

// 获取消息发送分布月数据（getupstreammsgdistmonth）	30	https://api.weixin.qq.com/datacube/getupstreammsgdistmonth?access_token=ACCESS_TOKEN
func (d *DataCube) GetUpstreamMsgDistMonth(beginDate, endDate time.Time) *core.Response {
	core.Debug("DataCube|GetUpstreamMsgDistMonth", beginDate, endDate)
	return d.get(
		DATACUBE_GETUPSTREAMMSGDISTMONTH_URL_SUFFIX,
		beginDate.Format("2006-01-02"),
		endDate.Format("2006-01-02"),
	)
}

func (d *DataCube) get(url, beginDate, endDate string) *core.Response {
	key := d.token.GetToken().KeyMap()
	resp := d.client.HttpPostJson(
		d.client.Link(url),
		core.Map{"begin_date": beginDate, "end_date": endDate},
		core.Map{core.REQUEST_TYPE_QUERY.String(): key})
	return resp
}
