package models

import "time"

type Quest struct {
  ID uint `json:"id" gorm:"primary_key"`  
  Title string `json:"title"`
  Description string `json:"description"`
  Reward int `json:"reward"`
  CreatedAt time.Time `json:"created_at"`
  UpdatedAt time.Time `json:"updated_at"`

}

func NewQuest(title string, description string, reward int) (quest *Quest, err error){
  quest = &Quest{
    Title: title,
    Description: description,
    Reward: reward,
  }

  DB.Create(&quest)
   
  return quest, nil
}


