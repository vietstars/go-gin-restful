package model

import (
  "go-gin-restful/database"
  "gorm.io/gorm"
  "time"
)

type Photo struct {
  gorm.Model
  ID int64 `gorm:"primarykey"`
  MediaID int64
  MediaType string
  UserID  int64
  Path string `gorm:"size:255;" json:"path"`
  CreatedAt time.Time `gorm:"autoCreateTime:true" json:"CreatedAt"`
  UpdatedAt time.Time `gorm:"autoUpdateTime:true" json:"UpdatedAt"`
  DeletedAt *time.Time `json:"-"`
}

// func (Media) TableName() string {
//   return "photos"
// }

func (media *Photo) Save() (*Photo, error) {
  err := database.DBS.Create(&media).Error
  if err != nil {
    return &Photo{}, err
  }
  return media, nil
}

func FindAvatarByUserID(userID int64) (Photo, error) {
  var avatar Photo

  database.DBS.Raw("SELECT * FROM photos WHERE user_id = ?", userID).Scan(&avatar)

  return avatar, nil
}
