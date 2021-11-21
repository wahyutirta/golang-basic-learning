package book

type BookInput struct {
	Title  string      `json:"title" binding:"required"`
	Price  interface{} `json:"price" binding:"required,numeric"` // prevent interface conversion: error is *json.SyntaxError, not validator.ValidationErrors
	Author string      `json:"author" binding:"required"`
}
