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
	voteKey := ctx.Params.ByName("votekey")

	db := common.GetDB()
	// 验证投票密钥是否存在
	var voter model.Voter
	if err := db.Where("`key` = ?", voteKey).First(&voter).Error; err != nil {
		response.Response(ctx, http.StatusUnauthorized, 401, nil, "无效的投票密钥")
		return
	}
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

	//var user model.User
	//db.First(&user, userID)
	//if user.ID == 0 {
	//	response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "选民不存在")
	//	return
	//}

	// 获取当前最大的 voteIndex
	var maxVoteIndex uint
	db.Model(&model.Voted{}).Where("vote_id = ? AND vote_key = ?", voteID, voteKey).Select("max(vote_index)").Row().Scan(&maxVoteIndex)

	//var count int64
	//db.Model(&model.Voted{}).Where("vote_id = ? AND vote_key = ?", voteID, voteKey).Count(&count)
	//var maxVoteIndex = uint(count)
	// 新的 voteIndex
	newVoteIndex := maxVoteIndex + 1

	// 检查是否超过了投票允许的最大次数
	if newVoteIndex > uint(vote.Num) {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "投票次数已达上限, 当前尝试为第 "+strconv.Itoa(int(newVoteIndex))+" 次")
		return
	}

	// 记录该次投票：userid：voteid
	newVoted := model.Voted{
		VoteKey:   voteKey,
		VoteID:    uint(voteID),
		VoteIndex: newVoteIndex,
	}

	db.Create(&newVoted)

	// 返回结果
	response.Success(ctx, nil, "投票成功")
}
