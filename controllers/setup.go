package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func New() http.Handler {
  router := mux.NewRouter()

  router.HandleFunc("/quest", CreateQuest).Methods("POST") 
  router.HandleFunc("/quest/{id}",GetQuest).Methods("GET")
  router.HandleFunc("/quest/{id}", DeleteQuest).Methods("DELETE")
  router.HandleFunc("/quests", GetAllQuests).Methods("GET")
  
  
  return router
}
