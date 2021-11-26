package book

import "encoding/json"

type BookRequest struct {
	Title       string      `json:"title" binding:"required"`
	Description string      `json:"description" binding:"required"`
	Author      string      `json:"author" binding:"required"`
	Rating      json.Number `json:"rating" binding:"required,numeric"`   // prevent interface conversion: error is *json.SyntaxError, not validator.ValidationErrors
	Discount    json.Number `json:"discount" binding:"required,numeric"` // prevent interface conversion: error is *json.SyntaxError, not validator.ValidationErrors
	Price       json.Number `json:"price" binding:"required,numeric"`    // prevent interface conversion: error is *json.SyntaxError, not validator.ValidationErrors
}
