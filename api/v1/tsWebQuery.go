package v1

import (
	"net/http"
	"tsweblist/model"
	"tsweblist/utils"

	"github.com/gin-gonic/gin"
)

// 添加服务器信息到数据库
func AddServerInfo(c *gin.Context) {
	var data model.TsServerInfo
	_ = c.ShouldBindJSON(&data)
	var status = model.AddServerv2(data)
	c.JSON(http.StatusOK, gin.H{
		"status":   status,
		"linksrv":  data.LinkSrv,
		"linkcity": data.LinkCity,
		"email":    data.Email,
	})
}

// 创建频道
func CreateChannel(c *gin.Context) {
	var data model.ChannelInfo
	data.ChannelPassword = utils.GeneratePassword()
	_ = c.ShouldBindJSON(&data)
	body, _ := model.CreateChannel(data)
	c.JSON(http.StatusOK, body)
}

// 获取服务器频道列表
func GetServerChannel(c *gin.Context) {
	data := model.GetServerChannel()
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

// 修改ChannelPassword
func ChangeChannelPassword(c *gin.Context) {
	var data model.ChangeChannelPassword
	_ = c.ShouldBindJSON(&data)
	c.JSON(http.StatusOK, data)
}
