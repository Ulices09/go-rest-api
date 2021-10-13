package entity

import "go-rest-api/ent"

type User struct {
	Model
	Email    string `json:"email"`
	Password string `json:"-"`
	Role     *Role  `json:"role,omitempty"`
	Posts    []Post `json:"posts,omitempty"`
}

func NewUserFromSchema(s *ent.User, mapPassword bool) *User {
	user := &User{
		Model: Model{
			ID:        s.ID,
			CreatedAt: s.CreatedAt,
			UpdatedAt: s.UpdatedAt,
		},
		Email: s.Email,
	}

	if mapPassword {
		user.Password = s.Password
	}

	if s.Edges.Role != nil {
		user.Role = &Role{
			Model: Model{ID: s.Edges.Role.ID},
			Name:  s.Edges.Role.Name,
		}
	}

	if len(s.Edges.Posts) > 0 {
		for _, p := range s.Edges.Posts {
			post := NewPostFromSchema(p)
			user.Posts = append(user.Posts, *post)
		}
	}

	return user
}
