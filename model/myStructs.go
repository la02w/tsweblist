package model

type ClientList struct {
	Count int `json:"count"`
	Body  []struct {
		ClientNickname string `json:"client_nickname"`
	} `json:"body"`
	Status struct {
		Code int    ` json:"code"`
		Msg  string ` json:"message"`
	} `json:"status"`
}
type ChannelData struct {
	Body []struct {
		CID string `json:"cid"`
	} `json:"body"`
	Status struct {
		Code int    ` json:"code"`
		Msg  string ` json:"message"`
	} `json:"status"`
}

type TsJsonData struct {
	Body   []Body     `json:"body"`
	Status StatusInfo `json:"status"`
}
type Body struct {
	ClientNickname string `json:"client_nickname,omitempty"`
	CID            string `json:"cid"`
}
type Status struct {
	Status StatusInfo ` json:"status"`
}
type StatusInfo struct {
	Code int    ` json:"code"`
	Msg  string ` json:"message"`
}
type ChannelInfo struct {
	ServerId        string `json:"sid"`
	ChannelName     string `json:"channel_name"`
	ChannelPassword string `json:"channel_password"`
}
