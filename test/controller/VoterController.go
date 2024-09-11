package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"test/common"
	"test/model"
)

// 添加投票者
func AddVoter(c *gin.Context) {
	var voter model.Voter
	voteID, _ := strconv.Atoi(c.Param("voteid")) // 从URL参数获取 voteid
	// log.Println("Received voteid: ", voteID)     // 调试输出到日志文件
	fmt.Println("Received voteid: ", c.Param("voteid"))
	if err := c.ShouldBindJSON(&voter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	voter.VoteID = uint(voteID) // 绑定 voteID 到投票者
	db := common.GetDB()        // 获取数据库连接
	if err := db.Create(&voter).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add voter"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Voter added successfully"})
}

// 查看投票的所有投票者
func ShowVoters(c *gin.Context) {
	voteID, _ := strconv.Atoi(c.Param("voteid")) // 从URL参数获取 voteid
	var voters []model.Voter

	db := common.GetDB() // 获取数据库连接
	if err := db.Where("vote_id = ?", voteID).Find(&voters).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch voters"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"voters": voters})
}
