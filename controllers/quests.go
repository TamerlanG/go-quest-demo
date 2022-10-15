package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/tamerlang/go-quest-demo/models"
	"github.com/tamerlang/go-quest-demo/utils"
)

var validate *validator.Validate

type QuestInput struct {
  Title string `json:"title" validate:"required"`
  Description string `json:"description" validate:"required"`
  Reward int `json:"reward" validate:"required"`
}

func CreateQuest(w http.ResponseWriter, r *http.Request){
  var input QuestInput 

  body, _ := ioutil.ReadAll(r.Body)
  _ = json.Unmarshal(body, &input)

  validate = validator.New()
  err := validate.Struct(input)

  if err != nil {
    utils.RespondWithError(w, http.StatusBadRequest, "Validation Error")
    return 
  }

  quest, err := models.NewQuest(input.Title, input.Description, input.Reward)

  w.Header().Set("Content-Type", "application/json")

  json.NewEncoder(w).Encode(quest) 

}


func GetAllQuests(w http.ResponseWriter, r *http.Request){
  w.Header().Set("Content-Type", "application/json")

  var quests []models.Quest
  models.DB.Find(&quests)

  json.NewEncoder(w).Encode(quests)
}


func GetQuest(w http.ResponseWriter, r *http.Request){
  w.Header().Set("Content-Type", "application/json")

  id := mux.Vars(r)["id"]
  var quest models.Quest

  if err := models.DB.Where("id = ?", id).First(&quest).Error; err != nil{
    utils.RespondWithError(w, http.StatusNotFound, "Quest not found")
    return
  }

  json.NewEncoder(w).Encode(quest)
}


func DeleteQuest(w http.ResponseWriter, r *http.Request){
  w.Header().Set("Content-Type", "application/json")

  id := mux.Vars(r)["id"]
  var quest models.Quest

  if err := models.DB.Where("id = ?", id).First(&quest).Error; err != nil{
    utils.RespondWithError(w, http.StatusNotFound, "Quest not found")
    return
  }

  models.DB.Delete(&quest)

  w.WriteHeader(http.StatusNoContent)
  json.NewEncoder(w).Encode(quest)
}


func UpdateQuest(w http.ResponseWriter, r *http.Request){
  w.Header().Set("Content-Type", "application/json")

  id := mux.Vars(r)["id"]
  var quest models.Quest

  if err := models.DB.Where("id = ?", id).First(&quest).Error; err != nil{
    utils.RespondWithError(w, http.StatusNotFound, "Quest not found")
    return
  }

  var input QuestInput 

  body, _ := ioutil.ReadAll(r.Body)
  _ = json.Unmarshal(body, &input)

  validate = validator.New()
  err := validate.Struct(input)

  if err != nil {
    utils.RespondWithError(w, http.StatusBadRequest, "Validation Error")
    return 
  }
  
  quest.Title = input.Title
  quest.Description = input.Description
  quest.Reward = input.Reward

  models.DB.Save(&quest)

  json.NewEncoder(w).Encode(quest)
}
