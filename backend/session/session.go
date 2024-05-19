package session

import (
	"errors"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/sichangSun/course-enrollment-manager/domain/model"
)

// JWT Claims
type AccountClaims struct {
	ID string `json:"id"`
	jwt.RegisteredClaims
}

var ErrFailedCreateToken = errors.New("invalid password")

// JWT CreateToken
func CreateToken(ctx echo.Context, student *model.Student) (string, error) {
	// Create Claims to create JWT
	id := strconv.Itoa(student.ID)
	claims := &AccountClaims{
		ID: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 1)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", ErrFailedCreateToken
	}
	return t, nil
}
