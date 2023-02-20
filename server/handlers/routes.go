package handlers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var companies *gin.RouterGroup
var searchers *gin.RouterGroup

var handleCors = cors.New(cors.Config{
	AllowOrigins:     []string{"*"}, // using '*' for the Dev Purpose only
	AllowMethods:     []string{"PUT", "GET", "POST", "DELETE"},
	AllowHeaders:     []string{"Origin"},
	AllowCredentials: true,
})

func RunServer() *gin.Engine {

	var router *gin.Engine = gin.Default()

	router.Use(handleCors)

	companies = router.Group("/company")
	searchers = router.Group("/searchers")

	return router
}
