package auth

import (
	"GoShort/configs"
	"GoShort/pkg/req"
	"GoShort/pkg/res"
	"fmt"
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
	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /auth/register", handler.Register())
}

func (handler *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[LoginRequest](&w, r)
		if err != nil {
			return
		}
		fmt.Println(body)
		data := LoginResponse{
			Token: "123",
		}
		res.Json(w, data, http.StatusOK)
	}
}

func (handler *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[RegisterRequest](&w, r)
		if err != nil {
			return
		}
		fmt.Println(body)
		data := RegisterResponse{
			Token: "123",
		}
		res.Json(w, data, http.StatusOK)
	}
}
