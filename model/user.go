package model

import (
  "go-gin-restful/database"
  "html"
  "strings"
  "time"
  // "fmt"

  "golang.org/x/crypto/bcrypt"
  "gorm.io/plugin/dbresolver"
  "gorm.io/sharding"
  "gorm.io/gorm"
)

type User struct {
  gorm.Model
  ID int64 `gorm:"primarykey"`
  Username string `gorm:"size:255;not null;unique" json:"username"`
  Password string `gorm:"size:255;not null;" json:"-"`
  Email string `gorm:"size:255;not null;unique" json:"email"`
  Avatar Photo `gorm:"foreignKey:UserID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
  Entries  []Entry `gorm:"foreignKey:UserID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
  CreatedAt time.Time `gorm:"autoCreateTime:true" json:"CreatedAt"`
  UpdatedAt time.Time `gorm:"autoUpdateTime:true" json:"UpdatedAt"`
  DeletedAt *time.Time `json:"-"`
}

func (user *User) Save() (*User, error) {
  err := database.DB.Create(&user).Error
  if err != nil {
    return &User{}, err
  }
  return user, nil
}

func (user *User) BeforeSave(*gorm.DB) error {
  passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
  if err != nil {
    return err
  }
  user.Password = string(passwordHash)
  user.Username = html.EscapeString(strings.TrimSpace(user.Username))
  return nil
}

func (user *User) ValidatePassword(password string) error {
  return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}

func FindUserByUsername(username string) (User, error) {
  var user User
  err := database.DB.Preload("Entries", func(dbs *gorm.DB) *gorm.DB {

    dbs.Use(dbresolver.Register(dbresolver.Config{
      Replicas: []gorm.Dialector{database.DBR.Dialector},
    }))

    dbs.Use(sharding.Register(sharding.Config{
      ShardingKey: "user_id",
      NumberOfShards: 26,
      PrimaryKeyGenerator: sharding.PKSnowflake,
    }, Entry{}))

    return dbs.Raw("SELECT * FROM entries WHERE user_id = ? order by updated_at DESC", user.ID)
  }).Where("username", username).Find(&user).Error

  user.Avatar, err = FindAvatarByUserID(user.ID)

  if err != nil {
    return User{}, err
  }
  return user, nil
}

// func FindUserByUsername(username string) (User, error) {
//   var user User
//   err := database.DB.Preload("Avatar", func(dbs *gorm.DB) *gorm.DB {

//     dbs.Use(dbresolver.Register(dbresolver.Config{
//       Replicas: []gorm.Dialector{database.DBR.Dialector},
//     }))

//     dbs.Use(sharding.Register(sharding.Config{
//       ShardingKey: "user_id",
//       NumberOfShards: 26,
//       PrimaryKeyGenerator: sharding.PKSnowflake,
//     }, Photo{}))

//     return dbs.Raw("SELECT * FROM photos WHERE user_id = ? order by updated_at DESC", user.ID)
//   }).Where("username", username).Find(&user).Error

//   user.Entries, err = FindEntriesByUserID(user.ID)

//   if err != nil {
//     return User{}, err
//   }
//   return user, nil
// }

func FindUserByEmail(email string) (User, error) {
  var user User
  err := database.DB.Preload("Entries", func(dbs *gorm.DB) *gorm.DB {

    dbs.Use(dbresolver.Register(dbresolver.Config{
      Replicas: []gorm.Dialector{database.DBR.Dialector},
    }))

    dbs.Use(sharding.Register(sharding.Config{
      ShardingKey: "user_id",
      NumberOfShards: 26,
      PrimaryKeyGenerator: sharding.PKSnowflake,
    }, Entry{}))

    return dbs.Raw("SELECT * FROM entries WHERE user_id = ? order by updated_at DESC", user.ID)
  }).Where("email", email).Find(&user).Error

  // user.Avatar, err = FindAvatarByUserID(user.ID)

  if err != nil {
    return User{}, err
  }
  return user, nil
}

// func FindUserByEmail(email string) (User, error) {
//   var user User
//   err := database.DB.Preload("Avatar", func(dbs *gorm.DB) *gorm.DB {

//     dbs.Use(dbresolver.Register(dbresolver.Config{
//       Replicas: []gorm.Dialector{database.DBR.Dialector},
//     }))

//     dbs.Use(sharding.Register(sharding.Config{
//       ShardingKey: "user_id",
//       NumberOfShards: 26,
//       PrimaryKeyGenerator: sharding.PKSnowflake,
//     }, Photo{}))

//     return dbs.Raw("SELECT * FROM photos WHERE user_id = ? order by updated_at DESC", user.ID)
//   }).Where("email", email).Find(&user).Error

//   user.Entries, err = FindEntriesByUserID(user.ID)

//   if err != nil {
//     return User{}, err
//   }
//   return user, nil
// }

func FindUserById(id int64) (User, error) {
  var user User
  err := database.DB.Preload("Entries", func(dbs *gorm.DB) *gorm.DB {

    dbs.Use(dbresolver.Register(dbresolver.Config{
      Replicas: []gorm.Dialector{database.DBR.Dialector},
    }))

    dbs.Use(sharding.Register(sharding.Config{
      ShardingKey: "user_id",
      NumberOfShards: 26,
      PrimaryKeyGenerator: sharding.PKSnowflake,
    }, Entry{}))

    return dbs.Raw("SELECT * FROM entries WHERE user_id = ? order by updated_at DESC", id)
  }).Where("id", id).Find(&user).Error

  // user.Avatar, err = FindAvatarByUserID(user.ID)

  if err != nil {
    return User{}, err
  }
  return user, nil
}

// func FindUserById(id int64) (User, error) {
//   var user User
//   err := database.DB.Preload("Avatar", func(dbs *gorm.DB) *gorm.DB {

//     dbs.Use(dbresolver.Register(dbresolver.Config{
//       Replicas: []gorm.Dialector{database.DBR.Dialector},
//     }))

//     dbs.Use(sharding.Register(sharding.Config{
//       ShardingKey: "user_id",
//       NumberOfShards: 26,
//       PrimaryKeyGenerator: sharding.PKSnowflake,
//     }, Photo{}))

//     return dbs.Raw("SELECT * FROM photos WHERE user_id = ? order by updated_at DESC", user.ID)
//   }).Where("id", id).Find(&user).Error

//   user.Entries, err = FindEntriesByUserID(user.ID)

//   if err != nil {
//     return User{}, err
//   }
//   return user, nil
// }

func FindUserEntriesById(id int64) (User, error) {
  var user User
  err := database.DB.Preload("Entries", func(dbs *gorm.DB) *gorm.DB {

    dbs.Use(dbresolver.Register(dbresolver.Config{
      Replicas: []gorm.Dialector{database.DBR.Dialector},
    }))

    dbs.Use(sharding.Register(sharding.Config{
      ShardingKey: "user_id",
      NumberOfShards: 26,
      PrimaryKeyGenerator: sharding.PKSnowflake,
    }, Entry{}))

    return dbs.Raw("SELECT * FROM entries WHERE user_id = ? order by updated_at DESC", id)
  }).Where("id", id).Find(&user).Error

  // user.Avatar, err = FindAvatarByUserID(user.ID)

  if err != nil {
    return User{}, err
  }
  return user, nil
}

func (user *User) EditUser() (*User, error) {
  err := database.DB.Save(user).Error

  user.Avatar, err = FindAvatarByUserID(user.ID)

  if err != nil {
    return &User{}, err
  }
  return user, nil
}

func (user *User) GetAvatar() (*User, error) {
  err := database.DB.Preload("Avatar", func(dbs *gorm.DB) *gorm.DB {

    dbs.Use(dbresolver.Register(dbresolver.Config{
      Replicas: []gorm.Dialector{database.DBR.Dialector},
    }))

    dbs.Use(sharding.Register(sharding.Config{
      ShardingKey: "user_id",
      NumberOfShards: 26,
      PrimaryKeyGenerator: sharding.PKSnowflake,
    }, Photo{}))

    return dbs.Raw("SELECT * FROM photos WHERE user_id = ? order by updated_at DESC", user.ID)
  }).Find(&user).Error

  if err != nil {
    return &User{}, err
  }

  return user, nil
}

func (user *User) AddAvatar(filePath string) (*User, error) {

  media := &Photo{
    UserID: user.ID,
    Path: filePath,
  }

  err := database.DBS.Create(&media).Error

  if err != nil {
    return &User{}, err
  }

  user.Avatar = *media

  if err != nil {
    return &User{}, err
  }

  return user, nil
}
