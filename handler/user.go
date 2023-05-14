package handler

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/tirthankarkundu17/ecommerce-price-checker/auth"
	"github.com/tirthankarkundu17/ecommerce-price-checker/model"
	"golang.org/x/crypto/bcrypt"
)

func (h *Server) Login(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	user := model.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	token, err := h.signIn(user.Email, user.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&map[string]string{
		"token": token,
	})
}

func (server *Server) signIn(email, password string) (string, error) {

	var err error

	user := model.User{}

	err = server.DB.Debug().Model(model.User{}).Where("email = ?", email).Take(&user).Error
	if err != nil {
		return "", err
	}
	err = model.VerifyPassword(user.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}
	return auth.CreateToken(user.ID)
}

func (server *Server) Register(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	user := model.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	user.Prepare()

	userCreated, err := user.SaveUser(server.DB)

	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userCreated)
}
