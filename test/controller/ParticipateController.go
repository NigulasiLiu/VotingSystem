package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strconv"
	"test/common"
	"test/model"
	"test/response"
)

import (
	"time"
)

func AddParticipation(ctx *gin.Context) {
	// 获取参数
	voteID, _ := strconv.Atoi(ctx.Params.ByName("voteid"))
	candidateID, _ := strconv.Atoi(ctx.Params.ByName("candidateid"))

	db := common.GetDB()
	// 数据验证
	var vote model.Vote
	db.First(&vote, voteID)
	if vote.ID == 0 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "投票不存在")
		return
	}

	// 检查投票是否已开放
	if vote.State == 1 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "投票已开始")
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

	var candidate model.Candidate
	db.First(&candidate, candidateID)
	if candidate.ID == 0 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "候选人不存在")
		return
	}

	// 检查是否已经存在相同的记录
	var existingParticipation model.Participate
	db.Where("vote_id = ? AND candidate_id = ?", voteID, candidateID).First(&existingParticipation)
	if existingParticipation.ID != 0 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "该候选人已参加该投票")
		return
	}

	// 创建新投票
	newParticipate := model.Participate{
		CandidateID: uint(candidateID),
		VoteID:      uint(voteID),
		Outs:        "",
	}

	db.Create(&newParticipate)

	// 返回结果
	response.Success(ctx, nil, "添加成功")
}

type ParticipateInfo struct {
	ID     uint
	Name   string
	Detail string
	Photo  []byte
	X      int    // 添加X字段，表示顺序索引
	Out    string // 添加Out字段
}

func ShowParticipate(ctx *gin.Context) {
	voteid, _ := strconv.Atoi(ctx.Params.ByName("voteid"))
	db := common.GetDB()

	var participateinfo []struct {
		ID     uint
		Name   string
		Photo  string
		Detail string
		Outs   string
	}

	// 使用 gorm 方法构建查询
	db.Table("participates p").
		Select("p.candidate_id AS ID, c.name AS Name, c.photo AS Photo, c.detail AS Detail, p.outs AS Outs").
		Joins("JOIN candidates c ON p.candidate_id = c.id").
		Where("p.vote_id = ?", voteid).
		Scan(&participateinfo)

	var participants []ParticipateInfo
	for index, candidate := range participateinfo {
		var imageData []byte
		if candidate.Photo != "" {
			// 从数据库中读取相对路径
			Path := candidate.Photo

			// 读取图片文件
			var err error
			imageData, err = ioutil.ReadFile(Path)
			if err != nil {
				fmt.Println("Error reading file:", err)
				continue
			}
		}

		// 构造候选人信息结构体
		info := ParticipateInfo{
			ID:     candidate.ID,
			Name:   candidate.Name,
			Detail: candidate.Detail,
			Photo:  imageData,      // 将图片数据存储到结构体中
			X:      index,          // 设置顺序索引
			Out:    candidate.Outs, // 设置候选人的Out
		}

		participants = append(participants, info)
	}

	ctx.JSON(http.StatusOK, gin.H{"code": 200, "data": participants})
}
