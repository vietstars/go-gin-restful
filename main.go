package main

import (
  "go-gin-restful/controller"
  "go-gin-restful/database"
  "go-gin-restful/middleware"
  "go-gin-restful/model"
  "strconv"
  "fmt"
  "log"
  "os"

  "github.com/gin-contrib/sessions"
  "github.com/gin-contrib/sessions/cookie"
  "github.com/gin-gonic/gin"
  "github.com/joho/godotenv"
)

func main() {
  loadEnv()
  loadDatabase()
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

  if database.DB.Migrator().HasTable(model.User{}) {
    database.DB.Migrator().DropTable(model.User{})
  }

  database.DB.AutoMigrate(&model.User{})

  if database.DBS.Migrator().HasTable(model.Entry{}) {
    database.DBS.Migrator().DropTable(model.Entry{})
  }

  database.DBS.AutoMigrate(&model.Entry{})
}

func serveApplication() {
  router := gin.Default()
  router.ForwardedByClientIP = true
  router.SetTrustedProxies([]string{os.Getenv("BASE_HOST")})

  sessionExp, _ := strconv.Atoi(os.Getenv("SESSION_EXP"))
  store := cookie.NewStore([]byte(os.Getenv("SESSION_SECRET")))
  store.Options(sessions.Options{MaxAge: sessionExp})
  router.Use(sessions.Sessions(os.Getenv("SESSION_KEY"), store))

  router.GET("/incr", controller.Incr)

  publicRoutes := router.Group("/auth")
  publicRoutes.POST("/register", controller.Register)
  publicRoutes.POST("/login", controller.Login)

  protectedRoutes := router.Group("/api")
  protectedRoutes.Use(middleware.JWTAuthMiddleware())
  protectedRoutes.POST("/entry", controller.AddEntry)
  protectedRoutes.GET("/entry", controller.GetAllEntries)
  protectedRoutes.POST("/edit-profile", controller.EditProfile)

  router.Run(":8008")
  fmt.Println("Server running on port 8008")
}