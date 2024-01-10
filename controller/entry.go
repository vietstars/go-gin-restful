package controller

import (
  "go-gin-restful/helper"
  "go-gin-restful/model"
  "net/http"

  "github.com/gin-gonic/gin"
  // "time"
  // "fmt"
)

func AddEntry(context *gin.Context) {
  var input model.Entry
  if err := context.ShouldBindJSON(&input); err != nil {

    context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  user, err := helper.CurrentUser(context)

  if err != nil {
    context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  input.UserID = int64(user.ID)

  savedEntry, err := input.Save()

  if err != nil {
    context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  context.JSON(http.StatusCreated, gin.H{"data": savedEntry})
}

func GetAllEntries(context *gin.Context) {

  user, err := helper.CurrentUser(context)

  if err != nil {
    context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  current, err := model.FindUserById(int64(user.ID))
  if err != nil {
    context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return 
  }

  // current, err := model.FindEntriesByUserID(int64(user.ID))
  // if err != nil {
  //   context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
  //   return 
  // }

  context.JSON(http.StatusOK, gin.H{"data": current.Entries})
}