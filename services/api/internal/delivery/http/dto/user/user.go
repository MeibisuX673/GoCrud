package user

import _ "github.com/go-playground/validator/v10"

type UserJsonRegistration struct{

	ID int `json:"id"`
	FirstName string `json:"firstName" validate:"required"`
	LastName string `json:"lastName" validate:"required"`
	Phone string `json:"phone" validate:"required"`
	Password string `json:"password" validate:"required"`
	Email string `json:"email" validate:"required,email"`

}

type UserJsonResponse struct{

	ID int `json:"id"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Phone string `json:"phone"`
	Email string `json:"email"`
	CreatedAt string `json:"createdAt"`

}

type UpdatedUser struct{

	FirstName string `json:"firstName" validate:"required"`
	LastName string `json:"lastName" validate:"required"`
	Phone string `json:"phone" validate:"required"`
	Email string `json:"email" validate:"required,email"`

}
