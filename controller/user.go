package controller

import (
  "go-gin-restful/helper"
  "go-gin-restful/model"
  "net/http"

  "github.com/gin-gonic/gin"
  // "time"
  // "fmt"
)

func EditProfile(context *gin.Context) {
  var input model.User
  if err := context.ShouldBindJSON(&input); err != nil {

    context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  user, err := helper.CurrentUser(context)

  if err != nil {
    context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  user.Email = input.Email
  user.Username = input.Username
  user.Avatar = input.Avatar

  updatedUser, err := user.EditUser()

  if err != nil {
    context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  context.JSON(http.StatusCreated, gin.H{"data": updatedUser})
}
