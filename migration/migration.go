package migration

import (
  "go-gin-restful/database"
  "go-gin-restful/model"
  "gorm.io/sharding"
)

func Run() {
  Sharding()

  if database.DB.Migrator().HasTable(model.User{}) {
    database.DB.Migrator().DropTable(model.User{})
  }
  if database.DB.Migrator().HasTable(model.Post{}) {
    database.DB.Migrator().DropTable(model.Post{})
  }

  database.DB.AutoMigrate(&model.User{}, &model.Post{})

  if database.DBS.Migrator().HasTable(model.Entry{}) {
    database.DBS.Migrator().DropTable(model.Entry{})
  }
  if database.DBS.Migrator().HasTable(model.Photo{}) {
    database.DBS.Migrator().DropTable(model.Photo{})
  }

  database.DBS.AutoMigrate(&model.Entry{}, &model.Photo{})
}

func Sharding() {
  database.DBS.Use(sharding.Register(sharding.Config{
    ShardingKey: "user_id",
    NumberOfShards: 26,
    PrimaryKeyGenerator: sharding.PKSnowflake,
  }, model.Entry{}, model.Photo{}))
}
