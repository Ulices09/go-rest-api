package entity

import "go-rest-api/ent"

type User struct {
	Model
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password,omitempty" validate:"required"`
	Posts    []Post `json:"posts,omitempty"`
}

func NewUserFromSchema(s *ent.User) *User {
	user := &User{
		Model: Model{
			ID:        s.ID,
			CreatedAt: s.CreatedAt,
			UpdatedAt: s.UpdatedAt,
		},
		Email:    s.Email,
		Password: s.Password,
	}

	if len(s.Edges.Posts) > 0 {
		for _, p := range s.Edges.Posts {
			post := NewPostFromSchema(p)
			user.Posts = append(user.Posts, *post)
		}
	}

	return user
}

func NewUserToLog(user User) User {
	user.Password = ""
	return user
}
