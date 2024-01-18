package route

import (
  "go-gin-restful/controller"
  "github.com/gin-gonic/gin"
)

func AuthRoute (router *gin.RouterGroup) {

  router.POST("/news", controller.Register)
  router.POST("/post", controller.Login)
}

