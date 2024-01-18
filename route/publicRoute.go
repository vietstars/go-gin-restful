package route

import (
  "go-gin-restful/controller"
  "github.com/gin-gonic/gin"
)

func PublicRoute (router *gin.RouterGroup) {

  router.POST("/register", controller.Register)
  router.POST("/login", controller.Login)
}

