package middleware

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/alitdarmaputra/belanja-project/bussiness"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type token struct {
	Id  int `json:"id"`
	Exp int `json:"exp"`
}

type Authetication interface {
	ExtractJWTUser(ctx *gin.Context) (*token, error)
}

type AutheticationImpl struct {
	secretKey string
}

func NewAuthentication(secretKey string) Authetication {
	return &AutheticationImpl{
		secretKey: secretKey,
	}
}

func (authentication *AutheticationImpl) ExtractJWTUser(ctx *gin.Context) (*token, error) {
	user, ok := ctx.Get("user")
	if !ok {
		return nil, bussiness.NewUnauthorizedError("User token not found")
	}

	if _, ok := user.(*jwt.Token); !ok {
		return nil, bussiness.NewUnauthorizedError("User token not found")
	}

	claims := user.(*jwt.Token).Claims.(jwt.MapClaims)

	res := new(token)
	buff := new(bytes.Buffer)
	json.NewEncoder(buff).Encode(&claims)
	json.NewDecoder(buff).Decode(res)

	return res, nil
}

func JWTMiddlewareAuth(jwtSecretKey string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := strings.Replace(
			ctx.GetHeader("Authorization"),
			"Bearer ",
			"",
			1,
		)

		if token = strings.TrimSpace(token); token == "" {
			ctx.AbortWithStatusJSON(
				http.StatusBadRequest,
				ErrResponse{
					Message:     "Bad Request",
					Status:      http.StatusBadRequest,
					Description: "Invalid Token",
				},
			)
			return
		}

		res, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
			if method, ok := t.Method.(*jwt.SigningMethodHMAC); !ok ||
				method != jwt.SigningMethodHS256 {
				return nil, errors.New("signing method invalid")
			}

			return []byte(jwtSecretKey), nil
		})

		if err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusBadRequest,
				ErrResponse{
					Message:     "Unauthorized",
					Status:      http.StatusUnauthorized,
					Description: err.Error(),
				},
			)
			return
		}

		ctx.Set("user", res)
		ctx.Next()
	}
}
