package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"server1/common"
	"server1/model"
	"server1/utils"
	"strconv"
)

func EvalVote(ctx *gin.Context) {
	// 从 URL 参数中获取 voteId
	voteID, _ := strconv.Atoi(ctx.Param("voteid"))
	ID, _ := strconv.Atoi(ctx.Param("id"))

	// 读取请求体中的数据
	k, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read request body"})
		return
	}

	fmt.Println("Received data:", k)

	// 处理获取到的数据，这里简单打印出来
	utils.Eval(ID, voteID, k)

}

func Compute(ctx *gin.Context) {
	// 从 URL 参数中获取 voteId
	voteID, err := strconv.Atoi(ctx.Param("voteid"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid vote ID"})
		return
	}

	// 获取数据库连接
	db := common.GetDB()
	var N int64
	db.Model(&model.Argument{}).Where("vote_id = ?", voteID).Count(&N)

	var eta int64
	db.Model(&model.Participate{}).Where("vote_id = ?", voteID).Count(&eta)

	outs := utils.Tallying(voteID, int(N), int(eta))

	//outs = utils.Verify(voteID, int(N), outs)

	outsString := make([]string, len(outs))
	for i, out := range outs {
		outsString[i] = out.String()
	}

	// 返回 JSON 响应
	ctx.JSON(http.StatusOK, gin.H{
		"outs":    outsString,
		"message": "获取成功",
	})
}
