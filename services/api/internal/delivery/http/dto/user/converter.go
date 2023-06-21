package user

import (

	"github.com/MeibisuX673/GoCrud/services/api/internal/domain/user"
	
)



func ConvertUserToResponse(user *user.User) *UserJsonResponse{

	responseUser := &UserJsonResponse{
		ID: user.ID,
		FirstName: user.FirstName,
		LastName: user.LastName,
		Phone: user.Phone,
		Email: user.Email,
		CreatedAt: user.CreatedAt,
	}

	return responseUser

}