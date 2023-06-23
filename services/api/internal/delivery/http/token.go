package http

import (
	"encoding/json"
	"net/http"

	"github.com/MeibisuX673/GoCrud/services/api/internal/delivery/http/dto/auth"
	"github.com/MeibisuX673/GoCrud/services/api/internal/delivery/http/errorResponse"
	"github.com/go-playground/validator/v10"
)


// GetJWToken
//  @Summary		Get JWT
//	@Description	Get JWT
//	@Tags			token
//	@Accept			json
//	@Produce		json
//  @Param body body auth.UserAuth true "body"
//	@Success		200	{array}	    auth.JwtToken
//	@Failure		400	{object}	errorResponse.Error
//	@Failure		404	{object}	errorResponse.Error
//	@Failure		500	{object}	errorResponse.Error
//	@Router			/token [post]
func (d *Delivery) GetToken(w http.ResponseWriter, r *http.Request){

	w.Header().Add("Content-type", "application/json")

	var jsonUserAuth auth.UserAuth

	if err := json.NewDecoder(r.Body).Decode(&jsonUserAuth); err != nil{

		er := errorResponse.Error{
			Message: "Error",
			StatusCode: 500,
			Desscription: err.Error(),
		}

		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(er)

		return

	}

	validate := validator.New()

	if err := validate.Struct(jsonUserAuth); err != nil{

		er := errorResponse.Error{
			Message: "Bad Request",
			StatusCode: 400,
			Desscription: err.Error(),
		}

		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(er)

		return

	}

	token, err := d.ucUser.GetJwtToken(&jsonUserAuth)

	if err != nil{

		er := errorResponse.Error{
			Message: "Error",
			StatusCode: 500,
			Desscription: err.Error(),
		}

		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(er)

		return

	}

	w.WriteHeader(http.StatusOK)


	response := auth.JwtToken{
		Token: token,
		Type: "Bearer",
	}

	json.NewEncoder(w).Encode(response)

}



