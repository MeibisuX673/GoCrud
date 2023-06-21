package user

import (
	"fmt"

	"github.com/MeibisuX673/GoCrud/pkg/jwt"
	"github.com/MeibisuX673/GoCrud/pkg/queryParametrs"
	"github.com/MeibisuX673/GoCrud/services/api/internal/delivery/http/dto/auth"
	"github.com/MeibisuX673/GoCrud/services/api/internal/domain/user"
)


func (uc *UseCase) Create(user *user.User) (*user.User, error){

	newUser, err := uc.userRepository.CreateUser(user)

	if err != nil{
		return nil, err
	}

	return newUser, nil
}

func (uc *UseCase) GetCollection(queryParams *queryParametrs.QueryParams) ([]*user.User, error){

	users, err := uc.userRepository.GetCollectionUser(queryParams)

	if err != nil{
		return nil, err
	}

	return users, nil

}

func (uc *UseCase) GetById(id int) (*user.User, error){

	user, err := uc.userRepository.GetByUserId(id)

	if err != nil{
		return nil, err
	}

	return user, nil

}

func(uc *UseCase) Delete(id int) error{

	if err := uc.userRepository.DeleteUser(id); err != nil{
		return err
	}

	return nil
	
}

func (uc *UseCase) Update(id int, updateUser user.User) (*user.User, error){

	UdatedUser, err := uc.userRepository.UpdateUser(id, updateUser)

	if err != nil{
		return nil, err
	}

	return UdatedUser, nil

}

func (uc *UseCase) GetJwtToken(userAuth *auth.UserAuth) (string, error){

	user, _ := uc.userRepository.GetUserBy(map[string]string{
		"email": userAuth.Email,
		"password": userAuth.Password,
	})
	fmt.Println(user)
	if user == nil{
		return "", nil
	}

	token, err := jwt.CreateJWToken()

	if err != nil{
		return "", err
	}

	return token, nil

}