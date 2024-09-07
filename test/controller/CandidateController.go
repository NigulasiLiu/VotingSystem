package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strconv"
	"test/common"
	"test/model"
	"test/response"
)

func AddCandidate(ctx *gin.Context) {
	// 获取参数
	var requestClass model.Candidate
	ctx.Bind(&requestClass)
	name := requestClass.Name
	detail := requestClass.Detail

	// 数据验证
	if len(name) > 20 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "名字不超过20位")
		return
	}
	db := common.GetDB()
	var candidate model.Candidate
	db.Where("name = ?", name).First(&candidate)
	if candidate.ID != 0 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "候选人已存在")
		return
	}

	// 创建候选人
	newCandidate := model.Candidate{
		Name:   name,
		Detail: detail,
		Photo:  "",
	}

	db.Create(&newCandidate)

	// 返回结果
	response.Success(ctx, nil, "添加候选人成功")
}

type CandidateInfo struct {
	ID     uint
	Name   string
	Detail string
	Photo  []byte
}

func ShowCandidate(ctx *gin.Context) {
	db := common.GetDB()
	var candidates []model.Candidate
	db.Find(&candidates)

	var candidateInfo []CandidateInfo
	for _, candidate := range candidates {
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
		info := CandidateInfo{
			ID:     candidate.ID,
			Name:   candidate.Name,
			Detail: candidate.Detail,
			Photo:  imageData, // 将图片数据存储到结构体中
		}

		candidateInfo = append(candidateInfo, info)
	}

	ctx.JSON(http.StatusOK, gin.H{"code": 200, "data": candidateInfo})
}

func PhotoUpdate(ctx *gin.Context) {
	// 获取候选人id
	id, _ := strconv.Atoi(ctx.Params.ByName("id"))

	// 获取上传的文件
	photo, err := ctx.FormFile("photo")
	if err != nil {
		response.Response(ctx, http.StatusBadRequest, 400, nil, "未找到上传的文件")
		return
	}

	// 构造文件保存的路径，以选民的id作为文件名
	filename := strconv.Itoa(id) + filepath.Ext(photo.Filename)
	savePath := filepath.Join("tmp", filename)

	// 保存文件到指定路径
	if err := ctx.SaveUploadedFile(photo, savePath); err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "保存文件失败")
		return
	}

	// 绑定body中的参数
	var candidate model.Candidate
	candidate.ID = uint(id)

	// 更新数据库中的图片路径
	db := common.GetDB()
	db.Model(&candidate).Update("photo", savePath)

	response.Success(ctx, nil, "图片上传成功")
}

func UpdateDetail(ctx *gin.Context) {
	candidateId, _ := strconv.Atoi(ctx.Params.ByName("id"))

	var requestClass model.Candidate
	ctx.Bind(&requestClass)
	detail := requestClass.Detail

	db := common.GetDB()
	var candidate model.Candidate
	candidate.ID = uint(candidateId)

	newCandidate := model.Candidate{
		Name:   "",
		Detail: detail,
		Photo:  "",
	}

	db.Model(&candidate).Updates(newCandidate)
	// 返回结果
	response.Success(ctx, nil, "信息修改成功")
}
