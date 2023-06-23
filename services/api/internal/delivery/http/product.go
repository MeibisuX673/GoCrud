package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	_ "github.com/MeibisuX673/GoCrud/services/api/internal/delivery/http/docs"
	jsonProduct "github.com/MeibisuX673/GoCrud/services/api/internal/delivery/http/dto/product"
	"github.com/MeibisuX673/GoCrud/services/api/internal/delivery/http/errorResponse"
	domainProduct "github.com/MeibisuX673/GoCrud/services/api/internal/domain/product"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"

)

// ListProducts
//  @Summary		Create Product
//	@Description	Create Product
// 	@Security ApiKeyAuth
//	@Tags			products
//	@Accept			json
//	@Produce		json
//  @Param 	body 	body jsonProduct.CreateProductJson true "body"
//	@Success		200	{object}	    jsonProduct.ProductJsonResponse
//	@Failure		400	{object}	errorResponse.Error
//	@Failure		404	{object}	errorResponse.Error
//	@Failure		500	{object}	errorResponse.Error
//	@Router			/products [post]
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

// ListProducts
//  @Summary		Get Collection Products
//	@Description	Get Collection Products
// 	@Security ApiKeyAuth
//	@Tags			products
//	@Accept			json
//	@Produce		json
//  @Param filter[name][eq] query string false "name"
//  @Param filter[name][includes] query string false "name"
//  @Param order[name] query string false "DESC/ASC"
//	@Success		200	{array}	    jsonProduct.ProductJsonResponse
//	@Failure		400	{object}	errorResponse.Error
//	@Failure		404	{object}	errorResponse.Error
//	@Failure		500	{object}	errorResponse.Error
//	@Router			/products [get]
func (d *Delivery) GetCollectionProduct(w http.ResponseWriter, r *http.Request){

	w.Header().Add("Content-type", "application/json")

	products, err := d.ucProduct.GetCollection(r)


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

// Delete Product
//  @Summary		Delete Product
//	@Description	Delete Product
// 	@Security ApiKeyAuth
//	@Tags			products
//	@Accept			json
//	@Produce		json
//  @Param id path int true "id"
//	@Success		204
//	@Failure		400	{object}	errorResponse.Error
//	@Failure		404	{object}	errorResponse.Error
//	@Failure		500	{object}	errorResponse.Error
//	@Router			/products/{id} [delete]
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

// UpdateProduct
//  @Summary		Update Product
//	@Description	Update Product
// 	@Security ApiKeyAuth
//	@Tags			products
//	@Accept			json
//	@Produce		json
//  @Param 	id 	path int true "id"
//  @Param 	body 	body jsonProduct.UpdateJsonProduct true "body"
//	@Success		200	{object}	    jsonProduct.ProductJsonResponse
//	@Failure		400	{object}	errorResponse.Error
//	@Failure		404	{object}	errorResponse.Error
//	@Failure		500	{object}	errorResponse.Error
//	@Router			/products/{id} [put]
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

// GetOneProduct
//  @Summary		Get One Product
// 	@Security ApiKeyAuth
//	@Description	Get One Product
//	@Tags			products
//	@Accept			json
//	@Produce		json
//  @Param id path int true "id"
//	@Success		200	{array}	    jsonProduct.ProductJsonResponse
//	@Failure		400	{object}	errorResponse.Error
//	@Failure		404	{object}	errorResponse.Error
//	@Failure		500	{object}	errorResponse.Error
//	@Router			/products/{id} [get]
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

		er := errorResponse.Error{
			Message: "Error",
			StatusCode: http.StatusInternalServerError,
			Desscription: err.Error(),
		}

		w.WriteHeader(http.StatusInternalServerError)

		json.NewEncoder(w).Encode(er)

		return
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