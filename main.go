package main

import (
  "go-gin-restful/controller"
  "go-gin-restful/database"
  "go-gin-restful/migration"
  "go-gin-restful/route"
  "strconv"
  "errors"
  "fmt"
  "log"
  "os"

  "github.com/gin-contrib/sessions"
  "github.com/gin-contrib/cors"
  "github.com/gin-contrib/sessions/cookie"
  "github.com/gin-contrib/static"
  "github.com/gin-gonic/gin"
  "github.com/joho/godotenv"
  "time"
)

func main() {
  loadEnv()
  loadDatabase()
  initImageLibrary()
  serveApplication()
}

func loadEnv() {
  err := godotenv.Load(".env.local")
  if err != nil {
    log.Fatal("Error loading .env file")
  }
}

func loadDatabase() {
  database.Connect()
  migration.Run()
}

func initImageLibrary() {
  if _, err := os.Stat("./web-app/public/images"); errors.Is(err, os.ErrNotExist) {
    err := os.MkdirAll("./web-app/public/images", os.ModePerm)
    if err != nil {
      log.Println(err)
    }
  }
}

func serveApplication() {
  router := gin.Default()
  router.ForwardedByClientIP = true
  router.SetTrustedProxies([]string{os.Getenv("BASE_HOST")})

  sessionExp, _ := strconv.Atoi(os.Getenv("SESSION_EXP"))
  store := cookie.NewStore([]byte(os.Getenv("SESSION_SECRET")))
  store.Options(sessions.Options{MaxAge: sessionExp})
  router.Use(cors.New(cors.Config{
    AllowOrigins:   []string{"https://localhost:8008/"},
    AllowMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
    AllowHeaders:   []string{"Origin"},
    ExposeHeaders:  []string{"Content-Length"},
    AllowCredentials: true,
    AllowOriginFunc: func(origin string) bool {
      return origin == "https://127.0.0.1:8008"
    },
    MaxAge: 12 * time.Hour,
  }))

  router.Use(sessions.Sessions(os.Getenv("SESSION_KEY"), store))

  router.GET("/incr", controller.Incr)

  router.Use(static.Serve("/", static.LocalFile("./web-app/dist/", false)))

  route.AuthRoute(router.Group("/auth"))
  route.PublicRoute(router.Group("/pub"))
  route.ProtectedRoutes(router.Group("/api"))

  router.Run(":8008")
  fmt.Println("Server running on port 8008")
}