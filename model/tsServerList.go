package model

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"
	"tsweblist/utils"

	ts3 "github.com/la02w/ts3-webquery"
	"gorm.io/gorm"
)

type TsServerInfo struct {
	gorm.Model
	LinkSrv  string `json:"linksrv"`
	LinkCity string `json:"linkcity"`
	Apikey   string `json:"apikey"`
	Email    string `json:"email"`
	LinkTime int64  `json:"linktime"`
	WebQuery string `json:"webquery"`
}

// 添加服务器信息
func AddServerv2(data TsServerInfo) ts3.Status {
	var server = TsServerInfo{
		LinkSrv:  data.LinkSrv,
		LinkCity: data.LinkCity,
		Apikey:   data.Apikey,
		Email:    data.Email,
		LinkTime: time.Now().Unix() + 604800000,
		WebQuery: data.WebQuery,
	}
	c, _ := ts3.Login(server.WebQuery, server.Apikey)
	resp, _ := c.ServerInfo()
	if resp.Status.Code != 0 {
		return resp.Status
	}
	db.Create(&server)
	return resp.Status
}

// 创建永久服务器
func CreateChannel(data ChannelInfo) (*ts3.Status, error) {
	var server TsServerInfo
	result := db.First(&server, "id = ?", data.ServerId)
	if result.Error != nil {
		return nil, result.Error
	}
	channelinfo := map[string]string{
		"channel_password":                        data.ChannelPassword,
		"channel_codec":                           "5",
		"channel_codec_quality":                   "10",
		"channel_flag_permanent":                  "1",
		"channel_maxclients":                      data.ChannelMaxclients,
		"channel_flag_maxclients_unlimited":       "0",
		"channel_flag_maxfamilyclients_inherited": "1",
		"channel_flag_maxfamilyclients_unlimited": "0",
	}
	c, _ := ts3.Login(server.WebQuery, server.Apikey)
	resp, _ := c.ChannelCreate(url.QueryEscape(data.ChannelName), channelinfo)
	cid := ""
	linkUrl := ""
	for _, item := range resp.Body {
		cid = item.CID
	}
	// 判断创建频道状态 不为0则创建失败
	if resp.Status.Code != 0 {
		return &resp.Status, nil
	}
	// 如果是ip+端口或者域名+端口 则切割字符串
	parts := strings.Split(server.LinkSrv, ":")

	if len(parts) == 2 {
		linkUrl = fmt.Sprintf("ts3server://%s?port=%s&cid=%s&channelpassword=%s", parts[0], parts[1], cid, data.ChannelPassword)
	} else {
		linkUrl = fmt.Sprintf("ts3server://%s?cid=%s&channelpassword=%s", server.LinkSrv, cid, data.ChannelPassword)
	}
	// 发送邮件
	utils.SeedEmail(data.Email, linkUrl)
	// log.Println(linkUrl)
	return &resp.Status, nil
}

// 服务器频道列表获取
func GetServerChannel() []ServerChannel {
	var server []TsServerInfo
	var list []ServerChannel
	db.Find(&server)
	for _, s := range server {
		var channel ServerChannel
		// 获取服务器信息
		c, _ := ts3.Login(s.WebQuery, s.Apikey)
		resp, _ := c.ServerInfo()
		for _, si := range resp.Body {
			channel.ServerName = si.ServerName
			channel.ServerMessage = si.ServerMessage
			channel.ServerMaxClient = si.ServerMaxClient
		}
		// 获取服务器频道信息
		channellist, _ := c.ChannelList()
		for _, cl := range channellist.Body {
			var chl ChannelListInfo
			chl.ChannelName = cl.ChannelName
			chl.ChannelID = cl.ChannelID
			chl.ChannelClient = cl.ChannelClient
			chl.ChannelMaxClient = "?"
			data, _ := c.ChannelInfo(cl.ChannelID)
			for _, info := range data.Body {
				chl.ChannelMaxClient = info.ChannelMaxClient
			}
			channel.ChannelList = append(channel.ChannelList, chl)
		}
		//获取服务器在线人数
		clientlist, _ := c.ClientList()
		channel.ServerClient = strconv.Itoa(len(clientlist.Body))
		list = append(list, channel)
	}
	return list
}
