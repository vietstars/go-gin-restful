package email

import (
  "go-gin-restful/model"
  "gopkg.in/gomail.v2"
  "crypto/tls"
  "strconv"
  "os"
  "fmt"
)

func SendEmail(user model.User) error {
  mailHost := os.Getenv("MAIL_HOST")
  mailPort, _ := strconv.Atoi(os.Getenv("MAIL_PORT"))
  username := os.Getenv("MAIL_USERNAME")
  password := os.Getenv("MAIL_PASSWORD")
  fromEmail := os.Getenv("MAIL_FROM_ADDRESS")
  fromName := os.Getenv("MAIL_FROM_NAME")

  d := gomail.NewDialer(mailHost, mailPort, username, password)
  d.TLSConfig = &tls.Config{InsecureSkipVerify: true}


  m := gomail.NewMessage()
  m.SetAddressHeader("From", fromEmail, fromName)
  m.SetAddressHeader("To", user.Email, user.Username)
  m.SetAddressHeader("Cc", "oliver.doe@example.com", "Oliver")
  m.SetHeader("Subject", fmt.Sprintf("Hello %s !", user.Username))
  m.SetBody("text/html", fmt.Sprintf("Hello <b>%s</b> and <i>%s</i>.<br/> Need verify email to complete register account!", user.Username, user.Email))

  // m.Attach("lolcat.jpg")

  if err := d.DialAndSend(m); err != nil {
    return err
  }

  return nil
}
