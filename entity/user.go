package entity

import "go-rest-api/ent"

type User struct {
	Model
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password,omitempty" validate:"required"`
	Posts    []Post `json:"posts,omitempty"`
}

func (u *User) MapFromSchema(s *ent.User) {
	u.ID = s.ID
	u.Email = s.Email
	u.Password = s.Password
	u.CreatedAt = s.CreatedAt
	u.UpdatedAt = s.UpdatedAt

	if len(s.Edges.Posts) > 0 {
		for _, p := range s.Edges.Posts {
			post := &Post{}
			post.MapFromSchema(p)
			u.Posts = append(u.Posts, *post)
		}
	}
}
