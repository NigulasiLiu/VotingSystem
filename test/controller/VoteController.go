package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"test/common"
	"test/dto"
	"test/model"
	"test/response"
	"time"
)

func AddVote(ctx *gin.Context) {
	// 获取参数
	var requestClass model.Vote
	ctx.Bind(&requestClass)
	name := requestClass.Name
	ddl := requestClass.Deadline
	num := requestClass.Num
	// 数据验证
	if len(name) > 30 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "投票名不超过30位")
		return
	}
	if num < 1 || num >= 10 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "可选人数不符合")
		return
	}

	// 创建新投票
	newVote := model.Vote{
		Name:     name,
		Deadline: ddl,
		Num:      num,
		State:    0,
	}
	db := common.GetDB()
	db.Create(&newVote)

	// 返回结果
	response.Success(ctx, gin.H{"id": newVote.ID}, "添加投票成功")
}

type VoteInfo struct {
	ID       uint
	Name     string
	Deadline string
	Num      int
	Outs     string
	State    int
}

func ShowVote(ctx *gin.Context) {
	db := common.GetDB()
	var voteinfo []VoteInfo
	db.Raw("SELECT ID,name AS Name,deadline AS Deadline,num AS Num,state AS State FROM votes").Scan(&voteinfo)
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "data": voteinfo})
}

func GetVoteInfo(ctx *gin.Context) {
	voteId, _ := strconv.Atoi(ctx.Params.ByName("id"))
	db := common.GetDB()
	var voteinfo model.Vote
	db.Where("id = ?", voteId).First(&voteinfo)
	if voteinfo.ID == 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "投票不存在"})
		return
	}
	var participate int64
	db.Model(&model.Participate{}).Where("vote_id = ?", voteId).Count(&participate)
	// 返回结果
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "data": gin.H{"vote": dto.ToVoteDto(voteinfo, int(participate))}})
}

func OpenVote(ctx *gin.Context) {
	voteId, _ := strconv.Atoi(ctx.Params.ByName("id"))

	db := common.GetDB()
	var vote model.Vote

	// 查找指定 ID 的投票项
	if err := db.First(&vote, voteId).Error; err != nil {
		response.Response(ctx, http.StatusNotFound, 404, nil, "投票不存在")
		return
	}

	// 检查投票项的 state 是否已经为 true
	if vote.State != 0 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "投票状态错误")
		return
	}

	// 更新投票项的 state 为 true
	vote.State = 1
	if err := db.Save(&vote).Error; err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "更新投票状态失败")
		return
	}

	response.Success(ctx, nil, "已开放投票")
}

func EndVote(ctx *gin.Context) {
	voteId, err := strconv.Atoi(ctx.Params.ByName("id"))

	db := common.GetDB()
	var vote model.Vote

	// 查找指定 ID 的投票项
	if err := db.First(&vote, voteId).Error; err != nil {
		response.Response(ctx, http.StatusNotFound, 404, nil, "投票不存在")
		return
	}

	// 将字符串格式的截止日期转换为时间类型
	deadlineTime, err := time.Parse("2006-01-02 15:04:05", vote.Deadline)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "日期解析错误")
		return
	}

	// 检查投票是否截止
	if time.Now().After(deadlineTime) {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "投票已截止")
		return
	}

	// 更新截止时间为当前时间
	vote.Deadline = time.Now().Format("2006-01-02 15:04:05")
	vote.State = 2
	if err := db.Save(&vote).Error; err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "更新投票截止时间失败")
		return
	}

	response.Success(ctx, nil, "投票截止")
}

func ComputeVote(ctx *gin.Context) {
	voteId, _ := strconv.Atoi(ctx.Params.ByName("id"))

	db := common.GetDB()
	var vote model.Vote

	// 查找指定 ID 的投票项
	if err := db.First(&vote, voteId).Error; err != nil {
		response.Response(ctx, http.StatusNotFound, 404, nil, "投票不存在")
		return
	}

	// 检查投票项的 state 是否已经为 true
	if vote.State == 0 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "投票未开放")
		return
	}

	if vote.State == 1 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "投票未截止")
		return
	}

	if vote.State != 2 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "投票状态错误")
		return
	}

	// 更新投票项的 state 为 true
	vote.State = 3
	if err := db.Save(&vote).Error; err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "更新投票状态失败")
		return
	}

	response.Success(ctx, nil, "开始计票")
}

type OutsRequest struct {
	Outs0 []int `json:"outs0"`
	Outs1 []int `json:"outs1"`
}

func PutOuts(ctx *gin.Context) {
	voteId, _ := strconv.Atoi(ctx.Params.ByName("id"))

	var outsRequest OutsRequest

	// 解析JSON请求体
	if err := ctx.ShouldBindJSON(&outsRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if len(outsRequest.Outs0) != len(outsRequest.Outs1) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Outs0 and Outs1 arrays must have the same length"})
		return
	}

	outs := make([]int, len(outsRequest.Outs0))

	// 按位相加
	for i := 0; i < len(outsRequest.Outs0); i++ {
		outs[i] = outsRequest.Outs0[i] + outsRequest.Outs1[i]
	}

	db := common.GetDB()
	// 查找并加载要更新的记录
	var vote model.Vote
	if err := db.Where("id = ?", voteId).First(&vote).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}

	vote.State = 4
	// 保存更改
	if err := db.Save(&vote).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update record"})
		return
	}

	// 获取参与记录，按ID递增排序
	var participates []model.Participate
	if err := db.Where("vote_id = ?", voteId).Order("id asc").Find(&participates).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve participates"})
		return
	}

	// 检查参与记录数量是否与outs长度一致
	if len(participates) != len(outs) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Number of participate records does not match number of outs"})
		return
	}

	// 按顺序给participates表中vote_id=voteId的项赋上值outs[i]
	for i, participate := range participates {
		if err := db.Model(&participate).Update("outs", outs[i]).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update participates"})
			return
		}
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Outs received successfully"})
}
