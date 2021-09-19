package entity

import "go-rest-api/ent"

type Post struct {
	Model
	Title string `json:"title" validate:"required"`
	Text  string `json:"text" validate:"required"`
	User  *User  `json:"user,omitempty"`
}

func NewPostFromSchema(s *ent.Post) *Post {
	post := &Post{
		Model: Model{
			ID:        s.ID,
			CreatedAt: s.CreatedAt,
			UpdatedAt: s.UpdatedAt,
		},
		Title: s.Title,
		Text:  s.Text,
	}

	if s.Edges.User != nil {
		userS := s.Edges.User
		post.User = NewUserFromSchema(userS)
	}

	return post
}
