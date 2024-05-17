package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"test/common"
	"test/model"
	"test/response"
	"time"
)

func AddVoted(ctx *gin.Context) {
	// 获取参数
	voteID, _ := strconv.Atoi(ctx.Params.ByName("voteid"))
	userID, _ := strconv.Atoi(ctx.Params.ByName("userid"))

	db := common.GetDB()
	// 数据验证
	var vote model.Vote
	db.First(&vote, voteID)
	if vote.ID == 0 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "投票不存在")
		return
	}

	// 将字符串格式的截止日期转换为时间类型
	deadlineTime, err := time.Parse("2006-01-02 15:04:05", vote.Deadline)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "日期解析错误")
		return
	}

	// 检查投票是否过期
	if time.Now().After(deadlineTime) {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "投票已经过期")
		return
	}

	var user model.User
	db.First(&user, userID)
	if user.ID == 0 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "选民不存在")
		return
	}

	// 检查是否已经存在相同的记录
	var existingVoted model.Voted
	db.Where("vote_id = ? AND user_id = ?", voteID, userID).First(&existingVoted)
	if existingVoted.ID != 0 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "已参加该投票")
		return
	}

	// 创建新投票
	newVoted := model.Voted{
		UserID: uint(userID),
		VoteID: uint(voteID),
	}

	db.Create(&newVoted)

	// 返回结果
	response.Success(ctx, nil, "投票成功")
}
