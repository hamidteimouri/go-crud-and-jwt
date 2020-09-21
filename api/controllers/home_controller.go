package controllers

import (
	"net/http"

	"github.com/hamidteimouri/go-crud-and-jwt/api/response"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	response.JSON(w, http.StatusOK, "Welcome To This Awesome API")

}
