package request

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"github.com/lechuhuuha/oneshield/pkg/sperrors"
	"github.com/lechuhuuha/oneshield/pkg/validate"
	"github.com/lechuhuuha/oneshield/transports/http/handlers/response"
)

type Payload interface {
	Validate() (bool, *response.ErrorsResponse)
}

func bodyValidate(l Payload) (bool, *response.ErrorsResponse) {
	var result = response.ErrorsResponse{
		Errors: make([]*response.ErrorResponse, 0, 2),
	}
	err := validate.Validate(l)
	if err != nil {
		result.Errors = validate.ValidationErrorsToErrorsResponse(err.(validator.ValidationErrors))
		return true, &result
	}

	return false, &result
}

// GetInt64Param returns int64 param from request.
func GetInt64Param(ctx *gin.Context, key string) (int64, error) {
	param := ctx.Param(key)
	if param == "" {
		return 0, nil
	}

	id, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		return 0, err
	}

	if id <= 0 {
		return 0, errors.New("invalid id")
	}

	return id, nil
}

// GetInt64Query returns int64 query from request.
func GetInt64Query(ctx *gin.Context, key string) (int64, error) {
	param := ctx.Query(key)
	if param == "" {
		return 0, nil
	}

	id, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		return 0, err
	}

	if id <= 0 {
		return 0, errors.New("invalid id")
	}

	return id, nil
}

// GetSearchNameQuery returns search name query from request.
func GetSearchNameQuery(ctx *gin.Context) string {
	searchName := ctx.Query("search")
	if len(searchName) > 100 {
		return searchName[:100]
	}

	return searchName
}

// GetQuery returns query from request.
func GetQuery(ctx *gin.Context, key string) string {
	value := ctx.Query(key)
	if len(value) > 100 {
		return value[:100]
	}

	return value
}

// GetBoolQuery returns bool query from request.
func GetBoolQuery(ctx *gin.Context, key string) bool {
	return ctx.Query(key) == "true"
}

var (
	ErrInvalidData = []*response.ErrorResponse{
		response.FormatError(
			"ERROR",
			"Bad request error.",
			"nproc must be greater than or equal to ep + 15",
			"nproc",
			sperrors.CodeInvalidData,
			nil,
		),
	}
)
