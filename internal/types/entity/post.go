package entity

import "go-rest-api/ent"

type Post struct {
	Model
	Title string `json:"title" validate:"required"`
	Text  string `json:"text" validate:"required"`
	User  *User  `json:"user,omitempty"`
}

func (p *Post) MapFromSchema(s *ent.Post) {
	p.ID = s.ID
	p.Title = s.Title
	p.Text = s.Text
	p.CreatedAt = s.CreatedAt
	p.UpdatedAt = s.UpdatedAt

	if s.Edges.User != nil {
		user := s.Edges.User
		p.User = &User{}
		p.User.MapFromSchema(user)
	}
}
