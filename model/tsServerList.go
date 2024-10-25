package model

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

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
//
//	@Param	data TsServerInfo 服务器数据信息
//
//	@Reutrn	Status	状态
func AddServerv2(data TsServerInfo) *Status {
	var server = TsServerInfo{
		LinkSrv:  data.LinkSrv,
		LinkCity: data.LinkCity,
		Apikey:   data.Apikey,
		Email:    data.Email,
		LinkTime: time.Now().Unix() + 604800000,
		WebQuery: data.WebQuery,
	}
	var status = checkApikey(server.WebQuery + "/serverinfo?api-key=" + server.Apikey)
	if status.Status.Code != 0 {
		return status
	}
	db.Create(&server)
	return status
}

// 验证WebQuery和APIkey
//
//	@Param	fullURL string 完整WebQueryURL地址
//
//	@Reutrn	Status	状态
func checkApikey(fullURL string) *Status {
	var resp, _ = http.Get(fullURL)
	defer resp.Body.Close()
	var body, _ = io.ReadAll(resp.Body)
	var status Status
	err = json.Unmarshal(body, &status)
	if err != nil {
		return nil
	}
	return &status
}

// 根据数据库索引ID获取服务器信息
//
//	@Param	id string 服务器索引ID
//
//	@Return	int,[]Client,int
//	在线人数，在线列表，状态码
func GetServerInfoByID(id string) (*ClientList, error) {
	var server TsServerInfo
	// 根据ID查询服务器记录
	result := db.First(&server, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	// 根据服务器保存的WebQuery和APIkey查询在线人数
	return getUserCount(&server)
}

// 根据数据库WebQuery和APIkey查询TeamSpeak服务器在线人数
//
//	@Param server TsServerInfo 服务器信息
//
//	@Return int,[]Client,int
//	在线人数，在线列表，状态码
func getUserCount(server *TsServerInfo) (*ClientList, error) {
	fullURL := server.WebQuery + "/1/clientlist?api-key=" + server.Apikey
	resp, err := http.Get(fullURL)
	// 处理请求错误
	if err != nil {
		fmt.Println("Error fetching URL:", err)
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	// 处理读取错误
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil, err
	}
	var response ClientList
	err = json.Unmarshal(body, &response)
	// 处理解析错误
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return nil, err
	}
	return &response, nil

}

// 创建永久服务器
func CreateChannel(data ChannelInfo) *ChannelData {
	var server TsServerInfo
	result := db.First(&server, "id = ?", data.ServerId)
	if result.Error != nil {
		return nil
	}
	fullURL := fmt.Sprintf("%s/1/channelcreate?api-key=%s&channel_name=%s&channel_password=%s&channel_codec=5&channel_codec_quality=10&channel_flag_permanent=1&channel_maxclients=8&channel_flag_maxclients_unlimited=0&channel_flag_maxfamilyclients_inherited=1&channel_flag_maxfamilyclients_unlimited=0",
		server.WebQuery,
		server.Apikey,
		url.QueryEscape(data.ChannelName),
		data.ChannelPassword,
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
	return &response
}
