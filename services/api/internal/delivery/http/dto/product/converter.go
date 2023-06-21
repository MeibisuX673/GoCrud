package product

import (

	"github.com/MeibisuX673/GoCrud/services/api/internal/delivery/http/dto/user"
	"github.com/MeibisuX673/GoCrud/services/api/internal/domain/product"
	
)

func ConvertProductToResponse(product *product.Product) *ProductJsonResponse{

	responseProduct := &ProductJsonResponse{
		Id: product.Id,
		Name: product.Name,
		Price: product.Price,
		Quantity: product.Quantity,
		CreatedAt: product.CreatedAt,
		User: *user.ConvertUserToResponse(&product.User),
		
	}

	return responseProduct

}