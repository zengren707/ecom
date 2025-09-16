package user

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zengren707/ecom/service/auth"
	"github.com/zengren707/ecom/types"
	"github.com/zengren707/ecom/utils"
)

type Hander struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Hander {
	return &Hander{store: store}
}

func (h *Hander) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
}

func (h *Hander) handleLogin(w http.ResponseWriter, r *http.Request) {

}
func (h *Hander) handleRegister(w http.ResponseWriter, r *http.Request) {
	var payload types.RegisterUserPayload
	if err := utils.ParseJSON(r, payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}
	//check if the user exists
	_, err := h.store.GetUserByEmail(payload.Email)
	if err == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with eamil %s already exists", payload.Email))
		return
	}

	hashedPassword, _ := auth.HashedPassword(payload.Password)

	//if it doesnt we create the new user
	err = h.store.CreateUser(types.User{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Password:  hashedPassword,
	})
	if err == nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusCreated, nil)

}
