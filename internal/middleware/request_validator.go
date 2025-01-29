package middleware

import (
	"context"
	"go-user-api/pkg/utils"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ValidateObjectID(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        id := r.URL.Query().Get("id")
        objectID, err := primitive.ObjectIDFromHex(id)
        if err != nil {
            utils.JSONResponse(w, http.StatusBadRequest, map[string]string{"error": "invalid ID format"})
            return
        }
        ctx := context.WithValue(r.Context(), "objectID", objectID)
        next.ServeHTTP(w, r.WithContext(ctx))
    }
}

func ValidateRequestBody(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if r.Body == nil {
            utils.JSONResponse(w, http.StatusBadRequest, map[string]string{"error": "empty request body"})
            return
        }
        next.ServeHTTP(w, r)
    }
}
