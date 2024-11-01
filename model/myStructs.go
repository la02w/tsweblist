package model

// 服务器在线列表响应体
type ClientList struct {
	Count int `json:"count"`
	Body  []struct {
		ClientNickname string `json:"client_nickname"`
	} `json:"body"`
	Status StatusInfo `json:"status"`
}

// 创建频道响应体
type ChannelData struct {
	Body []struct {
		CID string `json:"cid"`
	} `json:"body"`
	Status StatusInfo `json:"status"`
}

// 状态响应体
type Status struct {
	Status StatusInfo ` json:"status"`
}
type StatusInfo struct {
	Code int    ` json:"code"`
	Msg  string ` json:"message"`
}

// 创建频道请求体
type ChannelInfo struct {
	ServerId          string `json:"sid"`
	Email             string `json:"email"`
	ChannelName       string `json:"channel_name"`
	ChannelPassword   string `json:"channel_password"`
	ChannelMaxclients string `json:"channel_maxclients"`
}

type ServerChannel struct {
	ServerName      string            `json:"server_name"`
	ServerClient    string            `json:"server_client"`
	ServerMaxClient string            `json:"server_max_client"`
	ServerMessage   string            `json:"server_message"`
	ChannelList     []ChannelListInfo `json:"channel_list"`
}
type ChannelListInfo struct {
	ChannelName      string `json:"channel_name"`
	ChannelID        string `json:"cid"`
	ChannelClient    string `json:"total_clients"`
	ChannelMaxClient string `json:"max_client"`
}
