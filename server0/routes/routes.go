package routes

import (
	"github.com/gin-gonic/gin"
	"server0/controller"
	"server0/middleware"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.Use(middleware.CORSMiddleware())

	r.POST("/api/vote/participate/:voteid/:candidateid", controller.AddParticipation)
	r.POST("/api/vote/server0/:voteid", controller.EvalVote)
	r.POST("/api/compute/server0/:voteid", controller.Compute)

	return r
}
