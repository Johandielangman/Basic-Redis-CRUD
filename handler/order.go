package handler

import (
	"fmt"
	"net/http"
)

type Order struct{}

func (o *Order) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create has been called")
}

func (o *Order) List(w http.ResponseWriter, r *http.Request) {
	fmt.Println("List has been called")
}

func (o *Order) GetByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetByID has been called")
}

func (o *Order) UpdateByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("UpdateByID has been called")
}

func (o *Order) DeleteByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DeleteByID has been called")
}
