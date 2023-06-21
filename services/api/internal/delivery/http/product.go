package http

import (

	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"github.com/MeibisuX673/GoCrud/pkg/queryParametrs"
	jsonProduct "github.com/MeibisuX673/GoCrud/services/api/internal/delivery/http/dto/product"
	"github.com/MeibisuX673/GoCrud/services/api/internal/delivery/http/errorResponse"
	domainProduct "github.com/MeibisuX673/GoCrud/services/api/internal/domain/product"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"

)

func (d *Delivery) CreateProduct(w http.ResponseWriter, r *http.Request){

	w.Header().Add("Content-type", "application/json")

	var requestJsonProduct jsonProduct.CreateProductJson

	json.NewDecoder(r.Body).Decode(&requestJsonProduct)

	validate := validator.New()
	if err := validate.Struct(requestJsonProduct); err != nil{

		er := errorResponse.Error{
			Message: "Bad Request",
			StatusCode: 400,
			Desscription: err.Error(),
		}

		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(er)

		return
	}

	user, err := d.ucUser.GetById(requestJsonProduct.UserId)

	if err != nil{

		er := errorResponse.Error{
			Message: "Not Found",
			StatusCode: http.StatusNotFound,
			Desscription: err.Error(),
		}

		w.WriteHeader(http.StatusNotFound)

		json.NewEncoder(w).Encode(er)

		return
	}

	product := &domainProduct.Product{
		Name: requestJsonProduct.Name,
		Price: requestJsonProduct.Price,
		Quantity: requestJsonProduct.Quantity,
		User: *user,
	}

	newProduct, err := d.ucProduct.Create(product)

	if err != nil{

		er := errorResponse.Error{
			Message: "Error",
			StatusCode: http.StatusInternalServerError,
			Desscription: err.Error(),
		}

		w.WriteHeader(http.StatusInternalServerError)

		json.NewEncoder(w).Encode(er)

		return
	}

	response := jsonProduct.ConvertProductToResponse(newProduct)

	json.NewEncoder(w).Encode(response)

}

func (d *Delivery) GetCollectionProduct(w http.ResponseWriter, r *http.Request){

	w.Header().Add("Content-type", "application/json")

	queryParams := queryParametrs.New()

	param := r.URL.Query()
	page := param.Get("page")
	
	numberPage, err := strconv.Atoi(page)

	if err == nil{
		queryParams.Page = numberPage
	}

	products, err := d.ucProduct.GetCollection(queryParams)

	if err != nil{

		er := errorResponse.Error{
			Message: "Error",
			StatusCode: http.StatusInternalServerError,
			Desscription: err.Error(),
		}

		w.WriteHeader(http.StatusInternalServerError)

		json.NewEncoder(w).Encode(er)

		return
		
	}

	var response []*jsonProduct.ProductJsonResponse
	for _, value := range products{

		product := jsonProduct.ConvertProductToResponse(value)

		response = append(response, product)
	}

	json.NewEncoder(w).Encode(response)

}

func (d *Delivery) DeleteProduct(w http.ResponseWriter, r *http.Request){

	w.Header().Add("Content-type", "application/json")

	fmt.Println(mux.Vars(r))

	queryParam := mux.Vars(r)

	val, ok := queryParam["id"]

	if !ok{

		er := errorResponse.Error{
			Message: "Bad request",
			StatusCode: http.StatusBadRequest,
			Desscription: "Not found id",
		}

		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(er)

		return

	}

	id, err := strconv.ParseInt(val, 10, 32)

	if err != nil{
		fmt.Println(err.Error())
	}

	if _, err := d.ucProduct.GetById(int(id)); err != nil{

		er := errorResponse.Error{
			Message: "Not Found",
			StatusCode: http.StatusNotFound,
			Desscription: err.Error(),
		}

		w.WriteHeader(http.StatusNotFound)

		json.NewEncoder(w).Encode(er)

		return

	}

	if err := d.ucProduct.Delete(int(id)); err != nil{

		er := errorResponse.Error{
			Message: "Error",
			StatusCode: http.StatusInternalServerError,
			Desscription: err.Error(),
		}

		w.WriteHeader(http.StatusInternalServerError)

		json.NewEncoder(w).Encode(er)

		return

	}

	w.WriteHeader(http.StatusNoContent)

}

func (d *Delivery) UpdateProduct(w http.ResponseWriter, r *http.Request){

	w.Header().Add("Content-type", "application/json")

	queryParam := mux.Vars(r)

	val, ok := queryParam["id"]

	if !ok{

		er := errorResponse.Error{
			Message: "Bad request",
			StatusCode: http.StatusBadRequest,
			Desscription: "Not found id",
		}

		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(er)

		return

	}

	id, err := strconv.ParseInt(val, 10, 32)

	if err != nil{
		fmt.Println(err.Error())
	}

	if _, err := d.ucProduct.GetById(int(id)); err != nil{

		er := errorResponse.Error{
			Message: "Not Found",
			StatusCode: http.StatusNotFound,
			Desscription: err.Error(),
		}

		w.WriteHeader(http.StatusNotFound)

		json.NewEncoder(w).Encode(er)

		return

	}

	var updateProduct jsonProduct.UpdateJsonProduct

	if err := json.NewDecoder(r.Body).Decode(&updateProduct); err != nil{

		er := errorResponse.Error{
			Message: "Error",
			StatusCode: http.StatusInternalServerError,
			Desscription: err.Error(),
		}

		w.WriteHeader(http.StatusInternalServerError)

		json.NewEncoder(w).Encode(er)

		return
	}
	

	updatedProduct, err := d.ucProduct.Update(int(id), updateProduct)

	if err != nil{

		er := errorResponse.Error{
			Message: "Error",
			StatusCode: http.StatusInternalServerError,
			Desscription: err.Error(),
		}

		w.WriteHeader(http.StatusInternalServerError)

		json.NewEncoder(w).Encode(er)

		return

	}

	w.WriteHeader(http.StatusOK)

	response := jsonProduct.ConvertProductToResponse(updatedProduct)

	json.NewEncoder(w).Encode(response)

}

func (d *Delivery) GetByProductId(w http.ResponseWriter, r *http.Request){

	w.Header().Add("Content-type", "application/json")

	queryParam := mux.Vars(r)

	val, ok := queryParam["id"]

	if !ok{

		er := errorResponse.Error{
			Message: "Bad request",
			StatusCode: http.StatusBadRequest,
			Desscription: "Not found id",
		}

		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(er)

		return

	}

	id, err := strconv.ParseInt(val, 10, 32)

	if err != nil{
		fmt.Println(err.Error())
	}

	product, err := d.ucProduct.GetById(int(id))

	if err != nil{

		er := errorResponse.Error{
			Message: "Not Found",
			StatusCode: http.StatusNotFound,
			Desscription: err.Error(),
		}

		w.WriteHeader(http.StatusNotFound)

		json.NewEncoder(w).Encode(er)

		return

	}

	response := jsonProduct.ConvertProductToResponse(product)

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(response)

}