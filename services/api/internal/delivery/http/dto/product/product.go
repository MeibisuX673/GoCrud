package product

import "github.com/MeibisuX673/GoCrud/services/api/internal/delivery/http/dto/user"


type CreateProductJson struct{

	Name string `json:"name" validate:"required"`
	Price float32 `json:"price" validate:"required"`
	Quantity uint32 `json:"quantity" validate:"required"`
	UserId int `json:"userId" validate:"required"`

}

type ProductJsonResponse struct{

	Id int `json:"id"`
	Name string `json:"name"`
	Price float32 `json:"price"`
	Quantity uint32 `json:"quantity"`
	User user.UserJsonResponse `json:"user"`
	CreatedAt string `json:"createdAt"`

}

type UpdateJsonProduct struct{

	Name string `json:"name"`
	Price float32 `json:"price"`
	Quantity uint32 `json:"quantity"`

}