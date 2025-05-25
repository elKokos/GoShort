package auth

import (
	"GoShort/configs"
	"GoShort/pkg/res"
	"log"
	"net/http"
)

type AuthConfig struct {
	*configs.Config
}

type AuthHandler struct {
	*configs.Config
}

func NewAuthHandler(router *http.ServeMux, deps AuthConfig) {
	handler := &AuthHandler{deps.Config}
	router.HandleFunc("POST /auth/login", handler.Login)
	router.HandleFunc("POST /auth/register", handler.Register)
}

func (handler *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	log.Println(handler.Config.Auth.Secret)
	log.Println("Login")
	data := LoginResponse{
		Token: "123",
	}
	res.Json(w, data, http.StatusCreated)
}

func (handler *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	log.Println("Register")
}
