package handlers

import (
	"go-user-api/internal/database"
	"go-user-api/internal/models"
	"go-user-api/pkg/utils"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserHandler struct {
    db *database.MongoDB
}

func NewUserHandler(db *database.MongoDB) *UserHandler {
    return &UserHandler{db: db}
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
    objectID := r.Context().Value("objectID").(primitive.ObjectID)
    var User models.User
    err := h.db.FindOne(r.Context(), models.Users, bson.M{"_id": objectID}, &User)
    if err != nil {
        utils.JSONResponse(w, http.StatusNotFound, map[string]string{"error": err.Error()})
        return
    }
    utils.JSONResponse(w, http.StatusOK, User)
}

func (h *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
    Users, err := h.db.FindAll(r.Context(), models.Users, bson.M{})
    if err != nil {
        utils.JSONResponse(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
        return
    }
    utils.JSONResponse(w, http.StatusOK, Users)
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
    objectID := r.Context().Value("objectID").(primitive.ObjectID)
    var updates models.User
    if err := utils.ParseJSON(r, &updates); err != nil {
        utils.JSONResponse(w, http.StatusBadRequest, map[string]string{"error": "invalid request body"})
        return
    }
    result, err := h.db.UpdateOne(r.Context(), models.Users, bson.M{"_id": objectID}, bson.M{"$set": updates})
    if err != nil {
        utils.JSONResponse(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
        return
    }
    utils.JSONResponse(w, http.StatusOK, result)
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
    objectID := r.Context().Value("objectID").(primitive.ObjectID)
    result, err := h.db.DeleteOne(r.Context(), models.Users, bson.M{"_id": objectID})
    if err != nil {
        utils.JSONResponse(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
        return
    }
    utils.JSONResponse(w, http.StatusOK, result)
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
    var person models.User
    if err := utils.ParseJSON(r, &person); err != nil {
        utils.JSONResponse(w, http.StatusBadRequest, map[string]string{"error": "invalid request body"})
        return
    }

    // Set creation time
    now := time.Now()
    person.CreatedAt = now
    person.UpdatedAt = now

    result, err := h.db.InsertOne(r.Context(), models.Users, person)
    if err != nil {
        utils.JSONResponse(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
        return
    }

    utils.JSONResponse(w, http.StatusCreated, result)
}
