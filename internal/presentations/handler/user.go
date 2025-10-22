package handler

import (
	"encoding/json"
	"fmt"
	"golang-hexagonal/internal/core/domain/ports"
	"golang-hexagonal/internal/presentations/mapper"
	"net/http"
)

type UserHandler struct {
	service ports.UserServices
}

func NewUserHandler(
	service ports.UserServices,
) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (u *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {

	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "Missing user ID", http.StatusBadRequest)
		return
	}

	fmt.Printf("idStr: %v\n", idStr)

	userEntity, err := u.service.GetUser(idStr)
	if err != nil {
		http.Error(w, "An internal error occurred.", http.StatusInternalServerError)
		return
	}

	userResponse := mapper.ToHandler(userEntity)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(userResponse); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}
