package handlers

import (
	usecase "github.com/ThembinkosiThemba/go-project-starter/internal/application/usecases/user"
	domain "github.com/ThembinkosiThemba/go-project-starter/internal/entity/user"
	"github.com/ThembinkosiThemba/go-project-starter/pkg/dto"
	"github.com/ThembinkosiThemba/go-project-starter/pkg/events"
	httpRes "github.com/ThembinkosiThemba/go-project-starter/pkg/http"

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
	var user domain.USER
	if err := c.BindJSON(&user); err != nil {
		httpRes.WriteJSON(c, http.StatusBadRequest, 0, nil, err.Error())
		return
	}

	if err := h.useCase.AddUser(ctx, &user); err != nil {
		httpRes.WriteJSON(c, http.StatusInternalServerError, 0, nil, err.Error())
		return
	}

	events.TrackEvents("SIGNUP", user.ID, events.CreateEventProperties(user))
	events.UpdateUserProfile(user)

	httpRes.WriteJSON(c, http.StatusCreated, 1, user, "OK")
}

// Login handles the user login process.
// It binds the JSON request to a login DTO, retrieves the user,
// and tracks the login event.
func (h *UserHandler) Login(c *gin.Context) {
	var ctx, cancel = httpRes.Context()
	defer cancel()
	var request dto.Login
	if err := c.BindJSON(&request); err != nil {
		httpRes.WriteJSON(c, http.StatusBadRequest, 0, nil, err.Error())
		return
	}

	user, err := h.useCase.GetUser(ctx, request.Email, request.Password)
	if err != nil {
		httpRes.WriteJSON(c, http.StatusInternalServerError, 0, nil, err.Error())
		return
	}

	events.TrackEvents("LOGIN", user.ID, events.CreateEventProperties(user))

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
	if err := c.BindJSON(&request); err != nil {
		httpRes.WriteJSON(c, http.StatusBadRequest, 0, nil, err.Error())
		return
	}

	if err := h.useCase.Delete(ctx, request.Email); err != nil {
		httpRes.WriteJSON(c, http.StatusInternalServerError, 0, nil, err.Error())
		return
	}

	events.TrackEvents("DELETE_ACC", request.Email, nil)

	httpRes.WriteJSON(c, http.StatusOK, 1, nil, "user deleted successfully")
}
