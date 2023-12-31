package useCase

import (
	"github.com/MeibisuX673/GoCrud/pkg/queryParametrs"
	"github.com/MeibisuX673/GoCrud/services/api/internal/delivery/http/dto/auth"
	"github.com/MeibisuX673/GoCrud/services/api/internal/domain/product"
	"github.com/MeibisuX673/GoCrud/services/api/internal/domain/user"
	jsonProduct "github.com/MeibisuX673/GoCrud/services/api/internal/delivery/http/dto/product"
)


type User interface{

	GetJwtToken(userAuth *auth.UserAuth) (string, error)
	Create(user *user.User) (*user.User, error)
	GetCollection(queryParams *queryParametrs.QueryParams) ([]*user.User, error)
	GetById(id int) (*user.User, error)
	Delete(id int) error
	Update(id int, updateUser user.User) (*user.User, error)

}

type Product interface{

	Create(product *product.Product) (*product.Product, error)
	GetCollection(queryParams *queryParametrs.QueryParams) ([]*product.Product, error)
	GetById(id int) (*product.Product, error)
	Delete(id int) error
	Update(id int, updateProduct jsonProduct.UpdateJsonProduct) (*product.Product, error)

}

