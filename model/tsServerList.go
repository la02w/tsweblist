package model

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
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
func CreateChannel(data ChannelInfo) *ChannelData {
	var server TsServerInfo
	result := db.First(&server, "id = ?", data.ServerId)
	if result.Error != nil {
		return nil
	}
	fullURL := fmt.Sprintf("%s/1/channelcreate?api-key=%s&channel_name=%s&channel_password=%s&channel_codec=5&channel_codec_quality=10&channel_flag_permanent=1&channel_maxclients=%s&channel_flag_maxclients_unlimited=0&channel_flag_maxfamilyclients_inherited=1&channel_flag_maxfamilyclients_unlimited=0",
		server.WebQuery,
		server.Apikey,
		url.QueryEscape(data.ChannelName),
		data.ChannelPassword,
		data.ChannelMaxclients,
	)
	resp, err := http.Get(fullURL)
	// 处理请求错误
	if err != nil {
		fmt.Println("Error fetching URL:", err)
		return nil
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	// 处理读取错误
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil
	}
	var response ChannelData
	err = json.Unmarshal(body, &response)
	// 处理解析错误
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return nil
	}
	cid := ""
	linkUrl := ""
	for _, item := range response.Body {
		cid = item.CID
	}
	// 判断创建频道状态 不为0则创建失败
	if response.Status.Code != 0 {
		return &response
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
	return &response
}
