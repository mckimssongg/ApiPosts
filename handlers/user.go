package handlers

import (
	"encoding/json"
	"net/http"
	"rest-wsgo/models"
	"rest-wsgo/repository"
	"rest-wsgo/server"

	"github.com/segmentio/ksuid"
)

type SignupRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignupResponse struct {
	Id    string `json:"id"`
	Email string `json:"email"`
}

func SingnUpHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var request = SignupRequest{}

		err := json.NewDecoder(r.Body).Decode(&request)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		ksuidValue, err := ksuid.NewRandom()

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		id := ksuidValue.String()

		var user = models.User{
			Id:       id,
			Email:    request.Email,
			Password: request.Password,
		}

		// la variable "implementation" ha sido correctamente inicializad
		err = repository.IntertUser(r.Context(), &user)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(SignupResponse{
			Id:    user.Id,
			Email: user.Email,
		})
	}
}
