package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"test/common"
	"test/model"
	"test/response"
	"time"
)

// BallotController 通过密钥进行投票验证并设置 Cookie
func Ballot(c *gin.Context) {
	// 从请求体中获取投票密钥
	var requestBody struct {
		VoteKey string `json:"key"`
	}
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		response.Response(c, http.StatusBadRequest, 400, nil, "无效的请求体")
		return
	}

	fmt.Println("Received voteKey: ", requestBody.VoteKey) // 打印传入的 voteKey

	// 验证投票密钥是否存在于 Voter 模型中
	var voter model.Voter
	db := common.GetDB()
	if err := db.Where("`key` = ?", requestBody.VoteKey).First(&voter).Error; err != nil {
		response.Response(c, http.StatusUnauthorized, 401, nil, "无效的投票密钥")
		return
	}

	// 设置一个短期 Cookie，标识该投票者
	// 创建 Cookie，名称为 "vote_auth"，有效期为 1 小时
	cookie := &http.Cookie{
		Name:     "vote_auth",
		Value:    requestBody.VoteKey,           // 将 voteKey 存储在 Cookie 中
		Path:     "/",                           // 全局可用
		HttpOnly: true,                          // 只允许通过 HTTP 请求访问，防止 JS 读取
		Expires:  time.Now().Add(1 * time.Hour), // 1 小时后过期
	}

	// 将 Cookie 添加到响应头中
	http.SetCookie(c.Writer, cookie)

	// 返回成功响应，密钥验证成功
	response.Success(c, gin.H{"voteID": voter.VoteID}, "密钥验证成功")
}
