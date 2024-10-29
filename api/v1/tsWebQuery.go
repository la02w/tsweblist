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
		"status":   status.Status,
		"linksrv":  data.LinkSrv,
		"linkcity": data.LinkCity,
		"email":    data.Email,
	})
}

// 获取在线人数
func GetOnlineUserCount(c *gin.Context) {
	id := c.Param("id")
	var body, err = model.GetServerInfoByID(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": gin.H{
				"code":    500,
				"message": err.Error(),
			},
		})
	}
	body.Count = len(body.Body)
	c.JSON(http.StatusOK, body)
}

// 创建频道
func CreateChannel(c *gin.Context) {
	var data model.ChannelInfo
	data.ChannelPassword = utils.GeneratePassword()
	_ = c.ShouldBindJSON(&data)
	body := model.CreateChannel(data)
	c.JSON(http.StatusOK, body)
}
