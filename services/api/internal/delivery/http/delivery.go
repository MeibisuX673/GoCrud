package http

import (
	"net/http"

	"github.com/MeibisuX673/GoCrud/services/api/internal/useCase"
	"github.com/gorilla/mux"
)

type Delivery struct{

	ucUser useCase.User;
	ucProduct useCase.Product

	router *mux.Router
	
}

func New(ucUser useCase.User, ucProduct useCase.Product) *Delivery{

	delivery := &Delivery{
		ucUser: ucUser,
		ucProduct: ucProduct,
	}

	delivery.router = delivery.initRouter()

	return delivery
	
}

func (delivery *Delivery) Run() error{

	return http.ListenAndServe(":8080", delivery.router)

}

