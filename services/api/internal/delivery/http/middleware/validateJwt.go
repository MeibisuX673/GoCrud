package middleware

import (

	"encoding/json"
	"net/http"
	"os"
	"strings"
	"github.com/MeibisuX673/GoCrud/services/api/internal/delivery/http/errorResponse"
	"github.com/golang-jwt/jwt"

)


func ValidateJwt(next http.HandlerFunc) http.HandlerFunc{

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){

		if r.Header["Authorization"] != nil{

			jwtToken := strings.TrimLeft(r.Header["Authorization"][0], "Bearer") 
			jwtToken = strings.TrimLeft(jwtToken, " ")

			token, err := jwt.Parse(jwtToken, func(t *jwt.Token) (interface{}, error){

				_, ok := t.Method.(*jwt.SigningMethodHMAC)

				if !ok{

					er := errorResponse.Error{
						Message: "No authorized",
						StatusCode: http.StatusUnauthorized,
						Desscription: "No authorized",
					}
			
					w.Header().Add("Content-type", "application/json")
					w.WriteHeader(http.StatusUnauthorized)
			
					json.NewEncoder(w).Encode(er)
			
				}

				return []byte(os.Getenv("SECRET")), nil

			})

			if err != nil{

				er := errorResponse.Error{
					Message: "No authorized",
					StatusCode: http.StatusUnauthorized,
					Desscription: err.Error(),
				}
		
				w.Header().Add("Content-type", "application/json")
				w.WriteHeader(http.StatusUnauthorized)
		
				json.NewEncoder(w).Encode(er)
		
				return

			}

			if token.Valid{
				next(w, r)
			}

		}else{
			
			er := errorResponse.Error{
				Message: "No authorized",
				StatusCode: http.StatusUnauthorized,
				Desscription: "No authorized",
			}

			w.Header().Add("Content-type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
	
			json.NewEncoder(w).Encode(er)
	
			return

		}

	})
}


