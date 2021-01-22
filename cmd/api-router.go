package cmd

import "github.com/gin-gonic/gin"

func registerAPIRouter(router *gin.Engine) {

	api := APIHandlers{}

	tz := router.Group("/tz")

	tz.GET("/banner/list", api.GetBanners)

}
