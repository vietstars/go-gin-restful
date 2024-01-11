package controller

import (
  "go-gin-restful/helper"
  "go-gin-restful/model"
  "net/http"

  "github.com/gin-gonic/gin"
  "path/filepath"
  "strings"
  "time"
  "fmt"
  "log"
  "io"
  "os"
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

func EditAvatar(context *gin.Context) {
  user, err := helper.CurrentUser(context)

  if err != nil {
    context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    fmt.Println("error -2")
    return
  }

  file, header, err := context.Request.FormFile("image")
  if err != nil {
    context.String(http.StatusBadRequest, fmt.Sprintf("file err : %s", err.Error()))
    fmt.Println("error -1")
    return
  }

  fileExt := filepath.Ext(header.Filename)
  originalFileName := strings.TrimSuffix(filepath.Base(header.Filename), filepath.Ext(header.Filename))
  now := time.Now()
  filename := strings.ReplaceAll(strings.ToLower(originalFileName), " ", "-") + "-" + fmt.Sprintf("%v", now.Unix()) + fileExt
  filePath := fmt.Sprintf("%s/%s/%s", os.Getenv("BASE_URL"), os.Getenv("IMG_DIR"), filename)

  out, err := os.Create("./web-app/public/images/" + filename)
  if err != nil {
    log.Fatal(err)
    context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    fmt.Println("error 1")
    return
  }
  defer out.Close()

  _, err = io.Copy(out, file)
  if err != nil {
    log.Fatal(err)
    context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    fmt.Println("error 2")
    return
  }

  user.Avatar = filePath

  updatedUser, err := user.EditUser()

  if err != nil {
    context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    fmt.Println("error 3")
    return
  }

  context.JSON(http.StatusOK, gin.H{"data": updatedUser})
}
