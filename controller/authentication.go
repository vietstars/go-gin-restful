package controller

import (
  "go-gin-restful/helper"
  "go-gin-restful/model"
  "net/http"

  "github.com/gin-contrib/sessions"
  "github.com/gin-gonic/gin"
  "gopkg.in/gomail.v2"
  "strconv"
  "os"
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
  }

  savedUser, err := user.Save()

  if err != nil {
    context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  _, err = sendEmail()

  if err != nil {
    context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  context.JSON(http.StatusCreated, gin.H{"user": savedUser})
}

func sendEmail() (bool, error) {
  mailHost := os.Getenv("MAIL_HOST")
  mailPort, _ := strconv.Atoi(os.Getenv("MAIL_PORT"))
  username := os.Getenv("MAIL_USERNAME")
  password := os.Getenv("MAIL_PASSWORD")
  fromEmail := os.Getenv("MAIL_FROM_ADDRESS")
  fromName := os.Getenv("MAIL_FROM_NAME")

  d := gomail.NewPlainDialer(mailHost, mailPort, username, password)

  m := gomail.NewMessage()
  m.SetAddressHeader("From", fromEmail, fromName)
  m.SetAddressHeader("To", "recipient@example.com", "noah.doe@example.com")
  m.SetAddressHeader("Cc", "oliver.doe@example.com", "Oliver")
  m.SetHeader("Subject", "Hello!")
  m.SetBody("text/html", "Hello <b>Kate</b> and <i>Noah</i>!")

  // m.Attach("lolcat.jpg")

  if err := d.DialAndSend(m); err != nil {
    return false, err
  }

  return true, nil
}

func Login(context *gin.Context) {
  var input model.AuthenticationInput

  if err := context.ShouldBindJSON(&input); err != nil {

    context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  user, err := model.FindUserByUsername(input.Username)

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
