package views

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type UserResponse struct {
	ID       int    `json:"pid"`
	Username string `json:"puser"`
	Status   int    `json:"pstatus"`
}

type UserPost struct {
	Username string `json:"puser"`
	Password string `json:"ppass"`
	Status   int    `json:"pstatus"`
}

type UserPut struct {
	Username string `json:"puser"`
	Status   int    `json:"pstatus"`
}

func (u *UserPost) Validate() error {
	return validation.ValidateStruct(u,
		validation.Field(&u.Username, validation.Required, validation.Length(5, 50)),
		validation.Field(&u.Password, validation.Required, validation.Length(8, 12)),
		validation.Field(&u.Status, validation.Required, validation.In(0, 1)),
	)
}

func (u *UserPut) Validate() error {
	return validation.ValidateStruct(u,
		validation.Field(&u.Username, validation.Required, validation.Length(5, 50)),
		validation.Field(&u.Status, validation.Required, validation.In(0, 1)),
	)
}
