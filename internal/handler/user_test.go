package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"user-registration/internal/model"
	"user-registration/internal/service"
)

func TestRegisterUser(t *testing.T) {
	userService := service.NewUserService()
	userHandler := NewUserHandler(userService)

	// Test case
	user := model.User{
		Username: "testuser",
		Password: "password123",
		Email:    "test@example.com",
	}

	body, _ := json.Marshal(user)
	req, err := http.NewRequest("POST", "/register", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(userHandler.RegisterUser)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

	var createdUser model.User
	json.NewDecoder(rr.Body).Decode(&createdUser)

	if createdUser.Username != user.Username {
		t.Errorf("handler returned unexpected body: got %v want %v",
			createdUser.Username, user.Username)
	}
}
