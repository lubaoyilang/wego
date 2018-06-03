package official_account

import (
	"github.com/godcong/wego/config"
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/core/menu"
	"github.com/godcong/wego/net"
	"github.com/godcong/wego/util"
)

type Menu struct {
	config  config.Config
	account *OfficialAccount
	client  *core.Client
	token   *core.AccessToken
}

func newMenu(account *OfficialAccount) *Menu {
	return &Menu{
		config:  defaultConfig,
		account: account,
		client:  account.client,
		token:   account.token,
		//buttons: make(util.Map),
	}
}

func NewMenu() *Menu {
	return newMenu(account)
}

//func (m *Menu) SetMatchRule(rule *menu.MatchRule) *Menu {
//	m.buttons["matchrule"] = rule
//	return m
//}
//
//func (m *Menu) SetMenuId(id int) *Menu {
//	m.menuid = id
//	return m
//}

//个性化创建
//https://api.weixin.qq.com/cgi-bin/menu/addconditional?access_token=ACCESS_TOKEN
//成功:
//{"errcode":0,"errmsg":"ok"}

//自定义菜单
//https://api.weixin.qq.com/cgi-bin/menu/create?access_token=ACCESS_TOKEN
//成功:
// {"menuid":429680901}]
func (m *Menu) Create(buttons *menu.Button) *net.Response {
	token := m.token.GetToken().KeyMap()
	if buttons.GetMatchRule() == nil {
		resp := m.client.HttpPostJson(
			m.client.Link(MENU_CREATE_URL_SUFFIX),
			token,
			buttons)
		return resp
	}
	resp := m.client.HttpPostJson(
		m.client.Link(MENU_ADDCONDITIONAL_URL_SUFFIX),
		token,
		buttons)
	return resp
}

func (m *Menu) List() *net.Response {
	resp := m.client.HttpGet(m.client.Link(MENU_GET_URL_SUFFIX),
		m.token.GetToken().KeyMap(),
	)
	return resp
}

func (m *Menu) Current() *net.Response {
	resp := m.client.HttpGet(m.client.Link(GET_CURRENT_SELFMENU_INFO_URL_SUFFIX),
		m.token.GetToken().KeyMap())
	return resp
}

func (m *Menu) TryMatch(userId string) *net.Response {
	resp := m.client.HttpPostJson(m.client.Link(MENU_TRYMATCH_URL_SUFFIX),
		m.token.GetToken().KeyMap(),
		util.Map{"user_id": userId})
	return resp
}

func (m *Menu) Delete(menuid int) *net.Response {
	token := m.token.GetToken().KeyMap()
	if menuid == 0 {
		resp := m.client.HttpGet(m.client.Link(MENU_DELETE_URL_SUFFIX),
			token)
		return resp
	}

	resp := m.client.HttpPostJson(m.client.Link(MENU_DELETECONDITIONAL_URL_SUFFIX),
		util.Map{"menuid": menuid},
		token)
	return resp
}
