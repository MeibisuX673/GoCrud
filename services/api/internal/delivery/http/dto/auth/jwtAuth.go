package auth


type UserAuth struct{

	Email string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`

}

type JwtToken struct{

	Token string `json:"token"`
	Type string `json:"type"`

}