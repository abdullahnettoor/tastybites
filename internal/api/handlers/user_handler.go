package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/abdullahnettoor/tastybites/internal/api/dto"
	"github.com/abdullahnettoor/tastybites/internal/auth"
	"github.com/abdullahnettoor/tastybites/internal/usecases"
	"github.com/abdullahnettoor/tastybites/internal/utils"
)

type userHandler struct {
	UserUsecase  usecases.UserIUsecase
	TableUsecase usecases.TableIUsecase
}

func NewUserHandler(userUsecase usecases.UserIUsecase, tableUsecase usecases.TableIUsecase) *userHandler {
	return &userHandler{
		UserUsecase:  userUsecase,
		TableUsecase: tableUsecase,
	}
}

func (h *userHandler) UserLogin(w http.ResponseWriter, r *http.Request) {

	var LoginReq dto.UserLoginRequest
	if err := json.NewDecoder(r.Body).Decode(&LoginReq); err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	// Validate the login request
	if LoginReq.Email == "" || LoginReq.Password == "" {
		utils.WriteErrorResponse(w, http.StatusBadRequest, "Email and password are required")
		return
	}
	// Call the usecase to handle login
	user, err := h.UserUsecase.LoginUser(r.Context(), LoginReq.Email, LoginReq.Password)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusUnauthorized, err.Error())
		return
	}

	secretKey := os.Getenv("JWT_SECRET_KEY")

	token, _, err := auth.CreateToken(secretKey, user.Role, user.ID, time.Now().UTC().Sub(time.Time{}), time.Hour*24)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Create a response object
	resp := dto.LoginResponse{
		Token: token,
		User:  user,
	}

	// Send the response
	utils.WriteJSONResponse(w, http.StatusOK, resp)
}

func (h *userHandler) UserRegister(w http.ResponseWriter, r *http.Request) {
	var userReq dto.UserRegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&userReq); err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	// Validate the registration request
	if userReq.Email == "" || userReq.Password == "" || userReq.Name == "" {
		utils.WriteErrorResponse(w, http.StatusBadRequest, "Email, password, and name are required")
		return
	}

	user := dto.ToUserModel(userReq)

	// Call the usecase to handle user registration
	userID, err := h.UserUsecase.CreateUser(r.Context(), *user)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.WriteSuccessResponse(w, http.StatusCreated, "User registered successfully", map[string]int{"userId": userID})
}

func (h *userHandler) GetAvailableTables(w http.ResponseWriter, r *http.Request) {
	allTables, err := h.TableUsecase.GetAvailableTables(r.Context())
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.WriteJSONResponse(w, http.StatusOK, allTables)
}
