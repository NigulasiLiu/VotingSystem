package controller

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"server0/common"
	"server0/model"
	"server0/response"
	"server0/utils"
	"strconv"
)

func EvalVote(ctx *gin.Context) {
	// 从 URL 参数中获取 voteId
	voteID, err := strconv.Atoi(ctx.Param("voteid"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid vote ID"})
		return
	}

	// 读取请求体中的数据
	k, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read request body"})
		return
	}

	// fmt.Println("Received data:", k)

	newArgument := model.Argument{
		VoteID: uint(voteID),
	}
	db := common.GetDB()
	db.Create(&newArgument)

	utils.Eval(newArgument.ID, voteID, k)

	response.Success(ctx, gin.H{"id": newArgument.ID}, "")
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

	outs = utils.Verify(voteID, int(N), outs)

	response.Success(ctx, gin.H{"outs": outs}, "获取成功")
}
