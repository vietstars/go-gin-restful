package controller

import (
  "go-gin-restful/helper"
  "go-gin-restful/model"
  "go-gin-restful/email"
  "net/http"

  "github.com/gin-contrib/sessions"
  "github.com/gin-gonic/gin"
  // "time"
  // "fmt"
)

func Register(context *gin.Context) {
  var input model.AuthenticationInput

  if err := context.ShouldBindJSON(&input); err != nil {

    context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  user := model.User{
    Username: input.Username,
    Password: input.Password,
    Email: input.Email,
  }

  savedUser, err := user.Save()

  if err != nil {
    context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  err = email.SendEmail(user)

  if err != nil {
    context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  context.JSON(http.StatusCreated, gin.H{"user": savedUser})
}

func Login(context *gin.Context) {
  var input model.AuthenticationInput

  if err := context.ShouldBindJSON(&input); err != nil {

    context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  // user, err := model.FindUserByUsername(input.Username)
  user, err := model.FindUserByEmail(input.Email)

  if err != nil {
    context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  err = user.ValidatePassword(input.Password)

  if err != nil {
    context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  jwt, err := helper.GenerateJWT(user)
  if err != nil {
    context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  session := sessions.Default(context)
  session.Set("jwt", jwt)
  session.Save()

  context.JSON(http.StatusOK, gin.H{"jwt": jwt})
}

func Incr(context *gin.Context) {
  session := sessions.Default(context)
  var count int
  counter := session.Get("count")
  if counter == nil {
      count = 0
  } else {
      count = counter.(int)
      count++
  }
  session.Set("count", count)
  session.Save()
  context.JSON(200, gin.H{"count": count})
}
