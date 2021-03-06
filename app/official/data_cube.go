package official

import (
	"github.com/godcong/wego/core"
	"time"

	"github.com/godcong/wego/log"
	"github.com/godcong/wego/util"
)

/*DataCube DataCube */
type DataCube struct {
	//Config
	*Account
}

func newDataCube(acc *Account) *DataCube {
	return &DataCube{
		//Config:  defaultConfig,
		Account: acc,
	}
}

/*NewDataCube NewDataCube*/
func NewDataCube(config *core.Config) *DataCube {
	return newDataCube(NewOfficialAccount(config))
}

//GetUserSummary 获取用户增减数据（getusersummary）	7
// https://api.weixin.qq.com/datacube/getusersummary?access_token=ACCESS_TOKEN
func (d *DataCube) GetUserSummary(beginDate, endDate time.Time) core.Responder {
	log.Debug("DataCube|GetUserSummary", beginDate, endDate)
	return d.get(
		dataCubeGetUserSummaryURLSuffix,
		beginDate.Format(DatacubeTimeLayout),
		endDate.Format(DatacubeTimeLayout),
	)
}

//GetUserCumulate 获取累计用户数据（getusercumulate）	7
// https://api.weixin.qq.com/datacube/getusercumulate?access_token=ACCESS_TOKEN
func (d *DataCube) GetUserCumulate(beginDate, endDate time.Time) core.Responder {
	log.Debug("DataCube|GetUserCumulate", beginDate, endDate)
	return d.get(
		dataCubeGetUserCumulateURLSuffix,
		beginDate.Format(DatacubeTimeLayout),
		endDate.Format(DatacubeTimeLayout),
	)
}

//GetArticleSummary 获取图文群发每日数据（getarticlesummary）	1
// https://api.weixin.qq.com/datacube/getarticlesummary?access_token=ACCESS_TOKEN
func (d *DataCube) GetArticleSummary(beginDate, endDate time.Time) core.Responder {
	log.Debug("DataCube|GetArticleSummary", beginDate, endDate)
	return d.get(
		dataCubeGetArticleSummaryURLSuffix,
		beginDate.Format(DatacubeTimeLayout),
		endDate.Format(DatacubeTimeLayout),
	)
}

//GetArticleTotal 获取图文群发总数据（getarticletotal）	1
// https://api.weixin.qq.com/datacube/getarticletotal?access_token=ACCESS_TOKEN
func (d *DataCube) GetArticleTotal(beginDate, endDate time.Time) core.Responder {
	log.Debug("DataCube|GetArticleTotal", beginDate, endDate)
	return d.get(
		dataCubeGetArticleTotalURLSuffix,
		beginDate.Format(DatacubeTimeLayout),
		endDate.Format(DatacubeTimeLayout),
	)
}

//GetUserRead 获取图文统计数据（getuserread）	3
// https://api.weixin.qq.com/datacube/getuserread?access_token=ACCESS_TOKEN
func (d *DataCube) GetUserRead(beginDate, endDate time.Time) core.Responder {
	log.Debug("DataCube|GetUserRead", beginDate, endDate)
	return d.get(
		dataCubeGetUserReadURLSuffix,
		beginDate.Format(DatacubeTimeLayout),
		endDate.Format(DatacubeTimeLayout),
	)
}

//GetUserReadHour 获取图文统计分时数据（getuserreadhour）	1
// https://api.weixin.qq.com/datacube/getuserreadhour?access_token=ACCESS_TOKEN
func (d *DataCube) GetUserReadHour(beginDate, endDate time.Time) core.Responder {
	log.Debug("DataCube|GetUserReadHour", beginDate, endDate)
	return d.get(
		dataCubeGetUserReadHourURLSuffix,
		beginDate.Format(DatacubeTimeLayout),
		endDate.Format(DatacubeTimeLayout),
	)
}

//GetUserShare 获取图文分享转发数据（getusershare）	7
// https://api.weixin.qq.com/datacube/getusershare?access_token=ACCESS_TOKEN
func (d *DataCube) GetUserShare(beginDate, endDate time.Time) core.Responder {
	log.Debug("DataCube|GetUserReadHour", beginDate, endDate)
	return d.get(
		dataCubeGetUserShareURLSuffix,
		beginDate.Format(DatacubeTimeLayout),
		endDate.Format(DatacubeTimeLayout),
	)
}

//GetUserShareHour 获取图文分享转发分时数据（getusersharehour）	1
// https://api.weixin.qq.com/datacube/getusersharehour?access_token=ACCESS_TOKEN
func (d *DataCube) GetUserShareHour(beginDate, endDate time.Time) core.Responder {
	log.Debug("DataCube|GetUserReadHour", beginDate, endDate)
	return d.get(
		dataCubeGetUserShareHourURLSuffix,
		beginDate.Format(DatacubeTimeLayout),
		endDate.Format(DatacubeTimeLayout),
	)
}

//GetUpstreamMsg 获取消息发送概况数据（getupstreammsg）	7
// https://api.weixin.qq.com/datacube/getupstreammsg?access_token=ACCESS_TOKEN
func (d *DataCube) GetUpstreamMsg(beginDate, endDate time.Time) core.Responder {
	log.Debug("DataCube|GetUpstreamMsg", beginDate, endDate)
	return d.get(
		dataCubeGetUpstreamMsgURLSuffix,
		beginDate.Format(DatacubeTimeLayout),
		endDate.Format(DatacubeTimeLayout),
	)
}

//GetUpstreamMsgHour 获取消息分送分时数据（getupstreammsghour）	1
// https://api.weixin.qq.com/datacube/getupstreammsghour?access_token=ACCESS_TOKEN
func (d *DataCube) GetUpstreamMsgHour(beginDate, endDate time.Time) core.Responder {
	log.Debug("DataCube|GetUpstreamMsgHour", beginDate, endDate)
	return d.get(
		dataCubeGetUpstreamMsgHourURLSuffix,
		beginDate.Format(DatacubeTimeLayout),
		endDate.Format(DatacubeTimeLayout),
	)
}

//GetUpstreamMsgWeek 获取消息发送周数据（getupstreammsgweek）	30
// https://api.weixin.qq.com/datacube/getupstreammsgweek?access_token=ACCESS_TOKEN
func (d *DataCube) GetUpstreamMsgWeek(beginDate, endDate time.Time) core.Responder {
	log.Debug("DataCube|GetUpstreamMsgWeek", beginDate, endDate)
	return d.get(
		dataCubeGetUpstreamMsgWeekURLSuffix,
		beginDate.Format(DatacubeTimeLayout),
		endDate.Format(DatacubeTimeLayout),
	)
}

//GetUpstreamMsgMonth 获取消息发送月数据（getupstreammsgmonth）	30
// https://api.weixin.qq.com/datacube/getupstreammsgmonth?access_token=ACCESS_TOKEN
func (d *DataCube) GetUpstreamMsgMonth(beginDate, endDate time.Time) core.Responder {
	log.Debug("DataCube|GetUpstreamMsgMonth", beginDate, endDate)
	return d.get(
		dataCubeGetUpstreamMsgMonthURLSuffix,
		beginDate.Format(DatacubeTimeLayout),
		endDate.Format(DatacubeTimeLayout),
	)
}

//GetUpstreamMsgDist 获取消息发送分布数据（getupstreammsgdist）	15
// https://api.weixin.qq.com/datacube/getupstreammsgdist?access_token=ACCESS_TOKEN
func (d *DataCube) GetUpstreamMsgDist(beginDate, endDate time.Time) core.Responder {
	log.Debug("DataCube|GetUpstreamMsgDist", beginDate, endDate)
	return d.get(
		dataCubeGetUpstreamMsgDistURLSuffix,
		beginDate.Format(DatacubeTimeLayout),
		endDate.Format(DatacubeTimeLayout),
	)
}

//GetUpstreamMsgDistWeek 获取消息发送分布周数据（getupstreammsgdistweek）	30
// https://api.weixin.qq.com/datacube/getupstreammsgdistweek?access_token=ACCESS_TOKEN
func (d *DataCube) GetUpstreamMsgDistWeek(beginDate, endDate time.Time) core.Responder {
	log.Debug("DataCube|GetUpstreamMsgDistWeek", beginDate, endDate)
	return d.get(
		dataCubeGetUpstreamMsgDistWeekURLSuffix,
		beginDate.Format(DatacubeTimeLayout),
		endDate.Format(DatacubeTimeLayout),
	)
}

// GetUpstreamMsgDistMonth 获取消息发送分布月数据（getupstreammsgdistmonth）	30
// https://api.weixin.qq.com/datacube/getupstreammsgdistmonth?access_token=ACCESS_TOKEN
func (d *DataCube) GetUpstreamMsgDistMonth(beginDate, endDate time.Time) core.Responder {
	log.Debug("DataCube|GetUpstreamMsgDistMonth", beginDate, endDate)
	return d.get(
		dataCubeGetUpstreamMsgDistMonthURLSuffix,
		beginDate.Format(DatacubeTimeLayout),
		endDate.Format(DatacubeTimeLayout),
	)
}

//GetInterfaceSummary 获取接口分析数据（getinterfacesummary）	30
// https://api.weixin.qq.com/datacube/getinterfacesummary?access_token=ACCESS_TOKEN
func (d *DataCube) GetInterfaceSummary(beginDate, endDate time.Time) core.Responder {
	log.Debug("DataCube|GetInterfaceSummary", beginDate, endDate)
	return d.get(
		dataCubeGetInterfaceSummaryURLSuffix,
		beginDate.Format(DatacubeTimeLayout),
		endDate.Format(DatacubeTimeLayout),
	)
}

//GetInterfaceSummaryHour 获取接口分析分时数据（getinterfacesummaryhour）	1
// https://api.weixin.qq.com/datacube/getinterfacesummaryhour?access_token=ACCESS_TOKEN
func (d *DataCube) GetInterfaceSummaryHour(beginDate, endDate time.Time) core.Responder {
	log.Debug("DataCube|GetInterfaceSummaryHour", beginDate, endDate)
	return d.get(
		dataCubeGetInterfaceSummaryHourURLSuffix,
		beginDate.Format(DatacubeTimeLayout),
		endDate.Format(DatacubeTimeLayout),
	)
}

func (d *DataCube) get(url, beginDate, endDate string) core.Responder {
	key := d.accessToken.GetToken().KeyMap()
	resp := core.PostJSON(
		Link(url),
		key,
		util.Map{"begin_date": beginDate, "end_date": endDate})
	return resp
}
