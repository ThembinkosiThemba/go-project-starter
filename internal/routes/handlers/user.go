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

type UserHandler struct {
	useCase *usecase.UserUsecase
}

func NewUserHandler(useCase *usecase.UserUsecase) *UserHandler {
	return &UserHandler{useCase: useCase}
}

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

func (h *UserHandler) Delete(c *gin.Context) {
	var ctx, cancel = httpRes.Context()
	defer cancel()
	var request dto.Delete
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
