package handler

import (
	"Grocery-Shopping-Order-Module/src/app/model"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"net/http"
	"strconv"
)

func GetOrders(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	var result = authenticate(w, r)
	if result!=true{
		respondError(w, http.StatusUnauthorized,"Invalid User or Token Expired")
		return
	}
 users:=[] model.Order{}
 db.Set("gorm:auto_preload", true).Find(&users)
 respondJSON(w,http.StatusOK,users)
}
func CreateOrder(db *gorm.DB,w http.ResponseWriter, r *http.Request){
	var result = authenticate(w, r)
	if result!=true{
		respondError(w, http.StatusUnauthorized,"Invalid User or Token Expired")
		return
	}
	order:= model.Order{}
	decoder := json.NewDecoder(r.Body)
	if err:=decoder.Decode(&order); err!=nil{
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()
	err :=db.Create(&order)
	if  err.Error!=nil{
		fmt.Println(err.Error.Error())
		respondError(w, http.StatusBadRequest, err.Error.Error())
		return
	}
	respondJSON(w, http.StatusCreated,order)
}
func GetOrder(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	var result = authenticate(w, r)
	if result!=true{
		respondError(w, http.StatusUnauthorized,"Invalid User or Token Expired")
		return
	}
	vars := mux.Vars(r)

	id := vars["id"]
	i, err := strconv.ParseUint(id, 10, 64)
	if err!=nil{
		fmt.Println(err.Error())
		respondError(w, http.StatusBadRequest, err.Error())
	}
	order := getOrderOr404(db, uint(i), w, r)
	if order == nil {
		return
	}
	respondJSON(w, http.StatusOK, order)
}
func getOrderOr404(db *gorm.DB, id uint, w http.ResponseWriter, r *http.Request) *model.Order {
	authenticate(w,r)
	user := model.Order{}
	if err := db.Set("gorm:auto_preload", true).First(&user,id).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &user
}
