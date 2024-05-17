package routes

import (
	"github.com/gin-gonic/gin"
	"server1/controller"
	"server1/middleware"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.Use(middleware.CORSMiddleware())

	r.POST("/api/vote/server1/:voteid/:id", controller.EvalVote)
	r.POST("/api/compute/server1/:voteid", controller.Compute)

	return r
}
