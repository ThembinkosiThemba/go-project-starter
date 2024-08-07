package handlers

import (
	usecase "github.com/ThembinkosiThemba/go-project-starter/internal/application/usecases/user"
	entity "github.com/ThembinkosiThemba/go-project-starter/internal/entity/user"
	"github.com/ThembinkosiThemba/go-project-starter/pkg/dto"
	httpRes "github.com/ThembinkosiThemba/go-project-starter/pkg/http"
	"github.com/ThembinkosiThemba/go-project-starter/pkg/validate"

	"net/http"

	"github.com/gin-gonic/gin"
)

// UserHandler handles HTTP requests related to user operations.
type UserHandler struct {
	useCase *usecase.UserUsecase
}

// NewUserHandler creates a new UserHandler instance.
func NewUserHandler(useCase *usecase.UserUsecase) *UserHandler {
	return &UserHandler{useCase: useCase}
}

// Register handles the user registration process.
// It binds the JSON request to a user struct, adds the user to the system,
// and tracks the signup event.
func (h *UserHandler) Register(c *gin.Context) {
	var ctx, cancel = httpRes.Context()
	defer cancel()
	var user entity.USER
	if err := validate.BindDataToJson(c, &user); err != nil {
		return
	}

	if err := h.useCase.AddUser(ctx, &user); err != nil {
		httpRes.WriteJSON(c, http.StatusInternalServerError, 0, nil, err.Error())
		return
	}

	httpRes.WriteJSON(c, http.StatusCreated, 1, user, "OK")
}

// Login handles the user login process.
// It binds the JSON request to a login DTO, retrieves the user,
// and tracks the login event.
func (h *UserHandler) Login(c *gin.Context) {
	var ctx, cancel = httpRes.Context()
	defer cancel()
	var request dto.Login
	if err := validate.BindDataToJson(c, &request); err != nil {
		return
	}

	user, err := h.useCase.GetUser(ctx, request.Email, request.Password)
	if err != nil {
		httpRes.WriteJSON(c, http.StatusInternalServerError, 0, nil, err.Error())
		return
	}

	httpRes.WriteJSON(c, http.StatusOK, 1, user, "OK")
}

// GetAllUsers retrieves all users from the system.
func (h *UserHandler) GetAllUsers(c *gin.Context) {
	var ctx, cancel = httpRes.Context()
	defer cancel()
	users, err := h.useCase.GetAllUsers(ctx)
	if err != nil {
		httpRes.WriteJSON(c, http.StatusInternalServerError, 0, nil, err.Error())
		return
	}

	res := gin.H{"data": users, "count": len(users)}
	httpRes.WriteJSON(c, http.StatusOK, 1, res, "OK")
}

// Delete handles the user deletion process.
// It binds the JSON request to an email DTO, deletes the user,
// and tracks the account deletion event.
func (h *UserHandler) Delete(c *gin.Context) {
	var ctx, cancel = httpRes.Context()
	defer cancel()
	var request dto.Email
	if err := validate.BindDataToJson(c, &request); err != nil {
		return
	}

	if err := h.useCase.Delete(ctx, request.Email); err != nil {
		httpRes.WriteJSON(c, http.StatusInternalServerError, 0, nil, err.Error())
		return
	}

	httpRes.WriteJSON(c, http.StatusOK, 1, nil, "user deleted successfully")
}
