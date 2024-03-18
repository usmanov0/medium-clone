package handler

import (
	"encoding/json"
	"example.com/my-medium-clone/internal/common/jwt"
	"example.com/my-medium-clone/internal/domain"
	"example.com/my-medium-clone/internal/errors"
	"example.com/my-medium-clone/internal/usecase"
	"fmt"
	"log"
	"net/http"
)

type UserHandler struct {
	UserUseCase usecase.UserUseCase
}

func NewUserHandler(userUseCase usecase.UserUseCase) *UserHandler {
	return &UserHandler{UserUseCase: userUseCase}
}

func (b *UserHandler) SignUpUser(w http.ResponseWriter, r *http.Request) {
	var newUser domain.NewUser

	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "Invalid JSON request", http.StatusBadRequest)
	}

	id, err := b.UserUseCase.SignUpUser(&newUser)
	if err != nil {
		http.Error(w, "failed to create a new user", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(id)
}

func (b *UserHandler) SignInUser(w http.ResponseWriter, r *http.Request) {
	var userReq domain.SignInUser

	err := json.NewDecoder(r.Body).Decode(&userReq)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	ok, err := b.UserUseCase.SignInUser(userReq.Email, userReq.Password)
	if err != nil {
		http.Error(w, "Wrong email or password", http.StatusInternalServerError)
	}

	if ok {
		token, err := jwt.CreateToken(userReq.Email)
		if err != nil {
			http.Error(w, "email not found: ", http.StatusInternalServerError)
			return
		}
		response := map[string]string{
			"access_token": token,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	} else {
		http.Error(w, "bad credentials", http.StatusUnauthorized)
	}
}

func (b *UserHandler) GetById(w http.ResponseWriter, r *http.Request) {
	var user *domain.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	userId := user.Id

	user, err = b.UserUseCase.GetUserById(userId)
	if err != nil {
		if err == errors.ErrUserNotFound {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Intrenal server error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func (b *UserHandler) GetByEmail(w http.ResponseWriter, r *http.Request) {
	var req domain.User

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	email := req.Email

	log.Println("email", email)
	user, err := b.UserUseCase.GetUserByEmail(email)
	if err != nil {
		if err == errors.ErrUserNotFound {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func (b *UserHandler) GetList(w http.ResponseWriter, r *http.Request) {
	var req domain.User

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	criteria := req.UserName + " " + req.Email

	users, err := b.UserUseCase.ListUsers(criteria)
	if err != nil {
		if err == errors.ErrUserNotFound {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func (b *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	var userReq domain.User
	err := json.NewDecoder(r.Body).Decode(&userReq)
	if err != nil {
		http.Error(w, "Invalid items", http.StatusBadRequest)
		return
	}

	err = b.UserUseCase.UpdateUser(userReq.Id, &userReq)
	if err != nil {
		log.Println("")
		http.Error(w, "Faield to update user", http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"message": fmt.Sprint("user successfully updated"),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (b *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	var req domain.User

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid items", http.StatusBadRequest)
		return
	}

	userID := req.Id

	err = b.UserUseCase.DeleteUserAccount(userID)
	if err != nil {
		http.Error(w, "Faield to delete user", http.StatusBadRequest)
		return
	}

	response := map[string]interface{}{
		"message": fmt.Sprint("user deleted"),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
