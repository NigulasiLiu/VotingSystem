package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server0/common"
	"server0/model"
	"server0/response"
	"strconv"
)

func AddParticipation(ctx *gin.Context) {
	// 获取参数
	voteID, _ := strconv.Atoi(ctx.Params.ByName("voteid"))
	candidateID, _ := strconv.Atoi(ctx.Params.ByName("candidateid"))

	db := common.GetDB()

	// 检查是否已经存在相同的记录
	var existingParticipation model.Participate
	db.Where("vote_id = ? AND candidate_id = ?", voteID, candidateID).First(&existingParticipation)
	if existingParticipation.ID != 0 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "该候选人已参加该投票")
		return
	}

	// 添加新候选
	newParticipate := model.Participate{
		CandidateID: uint(candidateID),
		VoteID:      uint(voteID),
	}

	db.Create(&newParticipate)

	// 返回结果
	response.Success(ctx, nil, "添加成功")
}
