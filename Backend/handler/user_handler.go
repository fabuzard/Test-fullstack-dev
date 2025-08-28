package handler

import (
	"inventory_backend/dto"
	"inventory_backend/model"
	"inventory_backend/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	Service service.UserService
}

func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{Service: service}
}

func (h *UserHandler) Login(c echo.Context) error {
	var loginRequest dto.LoginRequest
	if err := c.Bind(&loginRequest); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}
	// validate input
	if err := c.Validate(&loginRequest); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	// Call service to authenticate
	user, err := h.Service.Login(loginRequest.Email, loginRequest.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid email or password"})
	}

	return c.JSON(http.StatusOK, user)

}

func (h *UserHandler) Register(c echo.Context) error {
	var registerRequest dto.RegisterRequest
	if err := c.Bind(&registerRequest); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}
	if err := c.Validate(&registerRequest); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	user := dto.RegisterResponse{}
	newUser, err := h.Service.Create(model.User{
		Fullname: registerRequest.Fullname,
		Email:    registerRequest.Email,
		Password: registerRequest.Password,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to register user"})
	}
	if newUser.ID == 0 {
		return c.JSON(http.StatusConflict, map[string]string{"error": "User already exists"})
	}
	user.Message = "User registered successfully"
	user.User.ID = newUser.ID
	user.User.Fullname = newUser.Fullname
	user.User.Email = newUser.Email
	return c.JSON(http.StatusCreated, user)
}
