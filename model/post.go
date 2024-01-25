package model

import (
  "go-gin-restful/database"
  "gorm.io/gorm"
  "time"
)

type Post struct {
  gorm.Model
  ID int64 `gorm:"primarykey"`
  Description string `gorm:"size:255" json:"description"`
  SourceID int64
  SourceType string
  Banner Photo `gorm:"polymorphic:Media;"`
  CreatedAt time.Time `gorm:"autoCreateTime:true" json:"CreatedAt"`
  UpdatedAt time.Time `gorm:"autoUpdateTime:true" json:"UpdatedAt"`
  DeletedAt *time.Time `json:"-"`
}

func (post *Post) Save() (*Post, error) {
  err := database.DBS.Create(&post).Error
  if err != nil {
    return &Post{}, err
  }
  return post, nil
}

