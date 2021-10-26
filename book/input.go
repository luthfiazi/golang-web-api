package book

import "encoding/json"

type BookInput struct {
	Title string `json:"title" binding:"required"`//title ini harus diisi
	Price json.Number`json:"price" binding:"required,number"`
}

