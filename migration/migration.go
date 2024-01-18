package migration

import (
  "go-gin-restful/database"
  "go-gin-restful/model"
)

func Run() {
  if database.DB.Migrator().HasTable(model.User{}) {
    database.DB.Migrator().DropTable(model.User{})
  }

  database.DB.AutoMigrate(&model.User{})

  if database.DBS.Migrator().HasTable(model.Entry{}) {
    database.DBS.Migrator().DropTable(model.Entry{})
  }

  database.DBS.AutoMigrate(&model.Entry{})
}
