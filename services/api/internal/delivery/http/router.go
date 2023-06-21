package http

import (
	"github.com/MeibisuX673/GoCrud/services/api/internal/delivery/http/middleware"
	"github.com/gorilla/mux"
)



func (delivery *Delivery) initRouter() *mux.Router{

	m := mux.NewRouter()

	delivery.router = m

	delivery.userRoutes()
	delivery.productRoutes()
	delivery.tokenRoutes()

	return m

}

func (delivery *Delivery) userRoutes(){

	delivery.router.HandleFunc("/users", delivery.CreateUser).Methods("POST")
	delivery.router.HandleFunc("/users", middleware.ValidateJwt(delivery.GetCollectionUser)).Methods("GET")
	delivery.router.HandleFunc("/users/{id}", middleware.ValidateJwt(delivery.GetByUserId)).Methods("GET")
	delivery.router.HandleFunc("/users/{id}", middleware.ValidateJwt(delivery.DeleteUser)).Methods("DELETE")
	delivery.router.HandleFunc("/users/{id}", middleware.ValidateJwt(delivery.UpdateUser)).Methods("PUT")

}

func (delivery *Delivery) productRoutes(){

	delivery.router.HandleFunc("/products", middleware.ValidateJwt(delivery.CreateProduct)).Methods("POST")
	delivery.router.HandleFunc("/products", middleware.ValidateJwt(delivery.GetCollectionProduct)).Methods("GET")
	delivery.router.HandleFunc("/products/{id}", middleware.ValidateJwt(delivery.GetByProductId)).Methods("GET")
	delivery.router.HandleFunc("/products/{id}", middleware.ValidateJwt(delivery.DeleteProduct)).Methods("DELETE")
	delivery.router.HandleFunc("/products/{id}", middleware.ValidateJwt(delivery.UpdateProduct)).Methods("PUT")

}

func (delivery *Delivery) tokenRoutes(){

	delivery.router.HandleFunc("/token", delivery.GetToken).Methods("POST")
	
}