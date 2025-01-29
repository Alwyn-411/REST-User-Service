package routes

import (
	"go-user-api/internal/handlers"
	"go-user-api/internal/middleware"

	"github.com/gorilla/mux"
)

func SetupRoutes(r *mux.Router, h *handlers.UserHandler) {
    // GET routes
    r.Handle("/person/{id}", middleware.LogRequest(middleware.ValidateObjectID(h.GetUser))).Methods("GET")
    r.Handle("/persons", middleware.LogRequest(h.GetAllUsers)).Methods("GET")

    // POST routes
    r.Handle("/person", middleware.LogRequest(middleware.ValidateRequestBody(h.CreateUser))).Methods("POST")

    // PUT routes
    r.Handle("/person/{id}", middleware.LogRequest(middleware.ValidateObjectID(middleware.ValidateRequestBody(h.UpdateUser)))).Methods("PUT")

    // DELETE routes
    r.Handle("/person/{id}", middleware.LogRequest(middleware.ValidateObjectID(h.DeleteUser))).Methods("DELETE")
}
