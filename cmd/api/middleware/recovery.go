package middleware

import (
	"fmt"
	"net/http"

	"github.com/alitdarmaputra/belanja-project/bussiness"
	"github.com/alitdarmaputra/belanja-project/cmd/api/common/response"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ErrorHandler(ctx *gin.Context, err any) {
	if notFoundError(ctx, err) {
		return
	}

	if validationError(ctx, err) {
		return
	}

	if unauthorizedError(ctx, err) {
		return
	}

	if duplicateEntryError(ctx, err) {
		return
	}

	internalServerError(ctx, err)
}

func notFoundError(ctx *gin.Context, err any) bool {
	execption, ok := err.(*bussiness.NotFoundError)
	if ok {
		response.JsonBasicData(ctx, http.StatusNotFound, "Not found", execption.Error())
		return true
	} else {
		return false
	}
}

func validationError(ctx *gin.Context, err any) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		var messages []string
		for _, fieldErr := range exception {
			messages = append(
				messages,
				fmt.Sprintf(
					"Validation error for field %s on tag %s",
					fieldErr.Field(),
					fieldErr.Tag(),
				),
			)
		}
		response.JsonBasicData(ctx, http.StatusBadRequest, "Bad request", messages)
		return true
	} else {
		return false
	}
}

func unauthorizedError(ctx *gin.Context, err any) bool {
	exception, ok := err.(*bussiness.UnauthorizedError)
	if ok {
		response.JsonBasicData(ctx, http.StatusUnauthorized, "Unauthorized", exception.Error())
		return true
	} else {
		return false
	}
}

func internalServerError(ctx *gin.Context, err any) {
	// TODO: Custom logger
	response.JsonBasicResponse(ctx, http.StatusInternalServerError, "Internal server error")
}

func duplicateEntryError(ctx *gin.Context, err any) bool {
	exception, ok := err.(*bussiness.DuplicateEntryError)
	if ok {
		response.JsonBasicData(ctx, http.StatusConflict, "Conflict", exception.Error())
		return true
	} else {
		return false
	}
}
