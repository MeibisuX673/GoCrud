package http

import (

	"encoding/json"
	"strconv"
	"net/http"
	"github.com/MeibisuX673/GoCrud/pkg/queryParametrs"
	userJson "github.com/MeibisuX673/GoCrud/services/api/internal/delivery/http/dto/user"
	"github.com/MeibisuX673/GoCrud/services/api/internal/delivery/http/errorResponse"
	domainUser "github.com/MeibisuX673/GoCrud/services/api/internal/domain/user"
	validator "github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	
)



func (d *Delivery) CreateUser(w http.ResponseWriter, r *http.Request){

	var userJsonRegistration userJson.UserJsonRegistration

	json.NewDecoder(r.Body).Decode(&userJsonRegistration)

	w.Header().Add("Content-type", "application/json")

	validate := validator.New()

	if err := validate.Struct(userJsonRegistration); err != nil{
		
		er := errorResponse.Error{
			Message: "Bad Request",
			StatusCode: http.StatusBadRequest,
			Desscription: err.Error(),
		}

		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(er)

		return

	}

	newUser := &domainUser.User{

		FirstName: userJsonRegistration.FirstName,
		LastName: userJsonRegistration.LastName,
		Phone: userJsonRegistration.Phone,
		Email: userJsonRegistration.Email,
		Password: userJsonRegistration.Password,

	}

	newUser, err := d.ucUser.Create(newUser)

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

	response := userJson.ConvertUserToResponse(newUser)	

	json.NewEncoder(w).Encode(response)

}


func (d *Delivery) GetCollectionUser(w http.ResponseWriter, r *http.Request){

	w.Header().Add("Content-type", "application/json")

	queryParams := queryParametrs.New()

	param := r.URL.Query()
	page := param.Get("page")
	
	numberPage, err := strconv.Atoi(page)

	if err == nil{
		queryParams.Page = numberPage
	}

	users, err := d.ucUser.GetCollection(queryParams)

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

	var responseUsers []*userJson.UserJsonResponse

	for _, element := range users{

		responseUser := userJson.ConvertUserToResponse(element)

		responseUsers = append(responseUsers, responseUser)

	}

	json.NewEncoder(w).Encode(responseUsers)

}

func (d *Delivery) GetByUserId(w http.ResponseWriter, r *http.Request){

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
			Message: "Bad request",
			StatusCode: http.StatusBadRequest,
			Desscription: err.Error(),
		}

		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(er)

		return

	}

	user, err := d.ucUser.GetById(int(id))

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

	response := userJson.ConvertUserToResponse(user)

	json.NewEncoder(w).Encode(response)

}

func (d *Delivery) DeleteUser(w http.ResponseWriter, r *http.Request){

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
			Message: "Bad request",
			StatusCode: http.StatusBadRequest,
			Desscription: err.Error(),
		}

		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(er)

		return

	}

	if _, err := d.ucUser.GetById(int(id)); err != nil{

		er := errorResponse.Error{
			Message: "Not Found",
			StatusCode: http.StatusNotFound,
			Desscription: err.Error(),
		}

		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(er)

		return

	}

	if err := d.ucUser.Delete(int(id)); err != nil{

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

func (d *Delivery) UpdateUser(w http.ResponseWriter, r *http.Request){

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
			Message: "Bad request",
			StatusCode: http.StatusBadRequest,
			Desscription: err.Error(),
		}

		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(er)

		return

	}

	if _, err := d.ucUser.GetById(int(id)); err != nil{

		er := errorResponse.Error{
			Message: "Not Found",
			StatusCode: http.StatusNotFound,
			Desscription: err.Error(),
		}

		w.WriteHeader(http.StatusNotFound)

		json.NewEncoder(w).Encode(er)

		return

	}

	var userJsonUpdate userJson.UpdatedUser

	json.NewDecoder(r.Body).Decode(&userJsonUpdate)

	validate := validator.New()

	if err := validate.Struct(userJsonUpdate); err != nil{

		er := errorResponse.Error{
			Message: "Bad Request",
			StatusCode: http.StatusBadRequest,
			Desscription: err.Error(),
		}

		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(er)

		return

	}

	updatedUser := domainUser.User{

		FirstName: userJsonUpdate.FirstName,
		LastName: userJsonUpdate.LastName,
		Email: userJsonUpdate.Email,
		Phone: userJsonUpdate.Phone,

	} 

	result, err := d.ucUser.Update(int(id), updatedUser)

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

	response := userJson.ConvertUserToResponse(result)

	json.NewEncoder(w).Encode(response)

}