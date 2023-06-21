package repositoryInterface

import (
	"github.com/MeibisuX673/GoCrud/pkg/queryParametrs"
	"github.com/MeibisuX673/GoCrud/services/api/internal/domain/product"
	"github.com/MeibisuX673/GoCrud/services/api/internal/domain/user"
)

type Storage interface {

	UserRepository
	ProductRepository

}

type UserRepository interface{

	CreateUser(user *user.User) (*user.User, error)
	GetCollectionUser(queryParams *queryParametrs.QueryParams) ([]*user.User, error)
	GetByUserId(id int) (*user.User, error)
	DeleteUser(id int) error
	UpdateUser(id int, updateUser user.User) (*user.User, error)
	GetUserBy(attribute map[string]string) (*user.User, error)

}

type ProductRepository interface{

	CreateProduct(product *product.Product) (*product.Product, error)
	GetCollectionProduct(queryParams *queryParametrs.QueryParams) ([]*product.Product, error)
	GetByProductId(id int) (*product.Product, error)
	DeleteProduct(id int) error
	UpdateProduct(id int, arguments map[string]string) (*product.Product, error)

}