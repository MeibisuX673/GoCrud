package http

import (
	"net/http"

	"github.com/MeibisuX673/GoCrud/services/api/internal/useCase"
	"github.com/gorilla/mux"
	_ "github.com/MeibisuX673/GoCrud/services/api/internal/delivery/http/docs"
	

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


// @title Scout Collector
// @version 1.0.0
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath /api
func (delivery *Delivery) Run() error{

	return http.ListenAndServe(":8081", delivery.router)

}

