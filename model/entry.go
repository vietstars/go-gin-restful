package model

import (
  "go-gin-restful/database"
  "gorm.io/gorm"
  // "fmt"
  "time"
)

type Entry struct {
  gorm.Model
  ID int64 `gorm:"primarykey"`
  Content string `gorm:"type:text" json:"content"`
  UserID  int64
  CreatedAt time.Time `gorm:"autoCreateTime:true" json:"CreatedAt"`
  UpdatedAt time.Time `gorm:"autoUpdateTime:true" json:"UpdatedAt"`
  DeletedAt *time.Time `json:"-"`
}

func (entry *Entry) Save() (*Entry, error) {
  err := database.DBS.Create(&entry).Error
  if err != nil {
    return &Entry{}, err
  }
  return entry, nil
}

// func FindEntriesByUserID(userID int64) ([] Entry, error) {
//   var entries []Entry

//   database.DBS.Raw("SELECT * FROM entries WHERE user_id = ?", userID).Scan(&entries)
//   fmt.Printf("%#v\n", entries)

//   return entries, nil
// }
