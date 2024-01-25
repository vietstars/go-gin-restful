package database

import (
  "fmt"
  "os"

  "gorm.io/driver/postgres"
  "gorm.io/plugin/dbresolver"
  "gorm.io/gorm"
)

//database master
var DB *gorm.DB
//database shading
var DBS *gorm.DB
//database replicate
var DBR *gorm.DB

func Connect() {
  var err error
  host := os.Getenv("DB_HOST")
  databaseName := os.Getenv("DB_NAME")
  username := os.Getenv("DB_USER")
  password := os.Getenv("DB_PASSWORD")
  port := os.Getenv("DB_PORT")

  slaveHost := os.Getenv("SLAVE_DB_HOST")
  slaveDBName := os.Getenv("SLAVE_DB_NAME")
  slaveDBUser := os.Getenv("SLAVE_DB_USER")
  slavePass := os.Getenv("SLAVE_DB_PASSWORD")
  slavePort := os.Getenv("SLAVE_DB_PORT")

  dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Saigon",
    host,
    username,
    password,
    databaseName,
    port,
  )

  dsnSlave := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Saigon",
    slaveHost,
    slaveDBUser,
    slavePass,
    slaveDBName,
    slavePort,
  )

  conn := postgres.Open(dsn)
  connSlave := postgres.Open(dsnSlave)

  DB, err = gorm.Open(conn, &gorm.Config{
    DisableForeignKeyConstraintWhenMigrating: true,
  })

  DBS, err = gorm.Open(conn, &gorm.Config{
    DisableForeignKeyConstraintWhenMigrating: true,
  })

  DBR, err = gorm.Open(connSlave, &gorm.Config{
    DisableForeignKeyConstraintWhenMigrating: true,
  })

  DBS.Use(dbresolver.Register(dbresolver.Config{
    Replicas: []gorm.Dialector{DBR.Dialector},
  }))

  if err != nil {
    panic(err)
  } else {
    fmt.Println("Successfully connected to the database")
  }
}