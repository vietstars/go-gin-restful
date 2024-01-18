package route

import (
  "go-gin-restful/controller"
  "go-gin-restful/middleware"
  "github.com/gin-gonic/gin"
)

func ProtectedRoutes (router *gin.RouterGroup) {

  router.Group("/api")
  router.Use(middleware.JWTAuthMiddleware())
  router.POST("/edit-profile", controller.EditProfile)
  router.POST("/edit-avatar", controller.EditAvatar)
  router.POST("/entry", controller.AddEntry)
  router.GET("/entry", controller.GetAllEntries)
}

