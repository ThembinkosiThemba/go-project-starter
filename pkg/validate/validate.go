package validate

import (
	"errors"
	"regexp"

	entity "github.com/ThembinkosiThemba/go-project-starter/internal/entity/user"
	"github.com/ThembinkosiThemba/go-project-starter/pkg/utils/logger"
	"github.com/gin-gonic/gin"

	"github.com/go-playground/validator/v10"
)

var ErrInternalServerError = errors.New("internal server error, something went wrong")

var validate = validator.New()

func IsEmailValid(email string) error {
	var rxEmail = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]{1,64}@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	if len(email) < 3 || len(email) > 30 || !rxEmail.MatchString(email) {
		return errors.New("email is not valid")
	}

	return nil
}

func ValidateUser(user *entity.USER) error {
	if err := validate.Struct(user); err != nil {
		return errors.New("please fill all values")
	}

	if err := IsEmailValid(user.Email); err != nil {
		return err
	}

	return nil
}

func BindDataToJson(c *gin.Context, data interface{}) error {
	if err := c.ShouldBindJSON(data); err != nil {
		logger.Error(err)
		return ErrInternalServerError
	}
	return nil
}
