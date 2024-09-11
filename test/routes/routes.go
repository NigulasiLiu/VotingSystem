package routes

import (
	"github.com/gin-gonic/gin"
	"test/controller"
	"test/middleware"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.Use(middleware.CORSMiddleware())

	// 需要登录和认证的路由
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login", controller.Login)
	r.GET("/api/auth/info", middleware.AuthMiddleware(), controller.Info)

	r.POST("/api/auth/addvote", middleware.AuthMiddleware(), controller.AddVote)
	r.GET("/api/auth/showvote", controller.ShowVote)
	r.GET("/api/auth/voteinfo/:id", controller.GetVoteInfo)
	r.POST("/api/auth/openvote/:id", middleware.AuthMiddleware(), controller.OpenVote)
	r.POST("/api/auth/endvote/:id", middleware.AuthMiddleware(), controller.EndVote)
	r.POST("/api/auth/computevote/:id", middleware.AuthMiddleware(), controller.ComputeVote)
	r.POST("/api/auth/outs/:id", middleware.AuthMiddleware(), controller.PutOuts)

	r.POST("/api/auth/addcandidate", middleware.AuthMiddleware(), controller.AddCandidate)
	r.POST("/api/auth/photo/:id", controller.PhotoUpdate)
	r.GET("/api/auth/showcandidate", controller.ShowCandidate)
	r.POST("/api/auth/candidate/detail/:id", middleware.AuthMiddleware(), controller.UpdateDetail)

	r.POST("/api/auth/participate/:voteid/:candidateid", middleware.AuthMiddleware(), controller.AddParticipation)
	r.GET("/api/auth/showparticipate/:voteid", controller.ShowParticipate)

	// 修改为 voteKey，而不是 userid
	r.POST("/api/auth/vote/:voteid/:votekey", controller.AddVoted)

	// 仍然需要认证的路由
	r.POST("/api/vote/participate/:voteid/:candidateid", middleware.AuthMiddleware(), controller.Redirect)
	r.POST("/api/vote/server0/:voteid", middleware.AuthMiddleware(), controller.Redirect)
	r.POST("/api/compute/server0/:voteid", middleware.AuthMiddleware(), controller.Redirect)

	r.POST("/api/vote/server1/:voteid/:id", middleware.AuthMiddleware(), controller.Redirect1)
	r.POST("/api/compute/server1/:voteid", middleware.AuthMiddleware(), controller.Redirect1)

	// 添加投票者相关的路由
	r.POST("/api/auth/addvoter/:voteid", middleware.AuthMiddleware(), controller.AddVoter)
	r.GET("/api/auth/showVoters/:voteid", controller.ShowVoters)

	// 投票者通过 voteKey 进行验证，而不需要认证 Token
	r.POST("/api/auth/ballot", controller.Ballot)

	return r
}
