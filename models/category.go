package models

import "time"

type Category struct {
  ID uint `json:"id" gorm:"primary_key"`  
  Title string `json:"title"`
  Quests []Quest
  CreatedAt time.Time `json:"created_at"`
  UpdatedAt time.Time `json:"updated_at"`
}
