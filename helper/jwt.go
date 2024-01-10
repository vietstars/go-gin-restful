package helper

import (
  "go-gin-restful/model"
  "errors"
  "fmt"
  "os"
  "strings"
  "time"

  "github.com/gin-contrib/sessions"
  "github.com/gin-gonic/gin"
  "github.com/golang-jwt/jwt/v4"
)

var privateKey = []byte(os.Getenv("JWT_PRIVATE_KEY"))

func GenerateJWT(user model.User) (string, error) {
  tokenTTL, _ := time.ParseDuration(os.Getenv("TOKEN_TTL"))
  token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
    "id":  user.ID,
    "iat": time.Now().Unix(),
    "eat": time.Now().Add(tokenTTL).Unix(),
  })
  return token.SignedString(privateKey)
}

func ValidateJWT(context *gin.Context) error {
  token, err := getToken(context)

  if err != nil {
    return err
  }

  _, ok := token.Claims.(jwt.MapClaims)

  if ok && token.Valid {
    return nil
  }

  return errors.New("invalid token provided")
}


func CurrentUser(context *gin.Context) (model.User, error) {
  err := ValidateJWT(context)
  if err != nil {
    return model.User{}, err
  }

  token, _ := getToken(context)
  claims, _ := token.Claims.(jwt.MapClaims)
  userId := int64(claims["id"].(float64))

  user, err := model.FindUserById(userId)
  if err != nil {
    return model.User{}, err
  }
  return user, nil
}

func getToken(context *gin.Context) (*jwt.Token, error) {
  session := sessions.Default(context)
   tokenString, ok := session.Get("jwt").(string)

  if !ok {
    tokenString = getTokenFromRequest(context)
  }

  token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
    if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
      return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
    }

    return privateKey, nil
  })
  return token, err
}

func getTokenFromRequest(context *gin.Context) string {
  bearerToken := context.Request.Header.Get("Authorization")

  if !strings.HasPrefix(bearerToken, "Bearer ") {
    return strings.TrimPrefix(bearerToken, "Bearer ")
  }

  return ""
}