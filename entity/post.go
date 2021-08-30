package entity

type Post struct {
	Model
	Title string `json:"title" validate:"required"`
	Text  string `json:"text" validate:"required"`
}
