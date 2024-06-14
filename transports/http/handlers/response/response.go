package response

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/lechuhuuha/oneshield/pkg/sperrors"
)

type (
	Response struct {
		Data any `json:"data,omitempty"`
	}
	// ErrorsResponse - error response slice
	ErrorsResponse struct {
		Errors []*ErrorResponse `json:"errors"`
	}

	// ErrorResponse - Error's response details
	ErrorResponse struct {
		Status string `json:"status,omitempty"`
		Title  string `json:"title,omitempty"`
		Source Source `json:"source,omitempty"`
		Detail string `json:"detail"`
		IsWarn bool   `json:"is_warn,omitempty"`
	}

	// Source - A component nested inside ErrorResponse
	Source struct {
		Field       string `json:"field,omitempty"`
		Value       any    `json:"value,omitempty"`
		ErrorCode   string `json:"error_code,omitempty"`
		TechDetails string `json:"tech_details,omitempty"`
	}

	WithMessage struct {
		Message string `json:"message"`
	}

	PaginationResponse[T any] struct {
		Data       []T      `json:"data"`
		Total      uint64   `json:"total"`
		Limit      uint64   `json:"limit"`
		Offset     uint64   `json:"offset"`
		Orderables []string `json:"orderables"`
	}
)

// HandleErrors : Return error with status code embedded and abort the response to make sure no more information will be put into this
func HandleErrors(c *gin.Context, statusCode int, errs *ErrorsResponse) {
	c.JSON(statusCode, errs)
}

// HandleError : Return error with status code embedded and abort the response to make sure no more information will be put into this
func HandleError(c *gin.Context, statusCode int, err *ErrorResponse) {
	c.Abort()
	HandleErrors(c, statusCode, &ErrorsResponse{Errors: []*ErrorResponse{err}})
}

func FormatError(status, title, details, field, errcode string, value any) *ErrorResponse {
	return &ErrorResponse{
		Status: status,
		Source: Source{
			Field:     field,
			Value:     value,
			ErrorCode: errcode,
		},
		Title:  title,
		Detail: details,
	}
}

func FormatError2(status, title, details, field, errcode string, value any, isWarn bool) *ErrorResponse {
	if isWarn {
		title = "Warning"
		status = "WARN"
	}

	return &ErrorResponse{
		Status: status,
		Source: Source{
			Field:     field,
			ErrorCode: errcode,
			Value:     value,
		},
		Title:  title,
		Detail: details,
		IsWarn: isWarn,
	}
}

func UnprocessableEntityResponse(ctx *gin.Context, err error) {
	HandleError(
		ctx,
		http.StatusUnprocessableEntity,
		FormatError(
			"ERROR",
			"Unprocessable Entity",
			err.Error(),
			"",
			sperrors.CodeInvalidData,
			nil,
		),
	)
}

func BadRequestResponse(ctx *gin.Context, err string) {
	HandleError(
		ctx,
		http.StatusBadRequest,
		FormatError(
			"ERROR",
			"Bad Request",
			err,
			"",
			sperrors.CodeInvalidData,
			nil,
		),
	)
}

func InternalServerErrorResponse(ctx *gin.Context, err error) {
	var detail = "Internal Server Error"
	if err != nil {
		detail = err.Error()
	}

	HandleError(
		ctx,
		http.StatusInternalServerError,
		FormatError(
			"ERROR",
			"Internal Server Error",
			detail,
			"",
			sperrors.CodeInternalServer,
			nil,
		),
	)
}

func NotFoundResponse(ctx *gin.Context) {
	HandleError(
		ctx,
		http.StatusNotFound,
		FormatError(
			"ERROR",
			"Not Found",
			"Not Found",
			"",
			sperrors.CodeNotFound,
			nil,
		),
	)
}

func UnauthorizedResponse(ctx *gin.Context) {
	HandleError(
		ctx,
		http.StatusUnauthorized,
		FormatError(
			"ERROR",
			"Unauthorized",
			"Unauthorized",
			"",
			sperrors.CodeNotAuthorized,
			nil,
		),
	)
}

func SuccessMessage() *WithMessage {
	return &WithMessage{
		Message: "Success",
	}
}

func ForbiddenResponse(ctx *gin.Context, permission string) {
	var (
		detail = "Your account is not allowed to access this resource"
	)

	if permission != "" {
		detail = "Your account is not allowed to access this resource: " + permission
	}

	HandleError(
		ctx,
		http.StatusForbidden,
		FormatError(
			"ERROR",
			"Forbidden",
			detail,
			"",
			sperrors.CodePermission,
			nil,
		),
	)
}

func CSRFErrorResponse(ctx *gin.Context) {
	HandleError(
		ctx,
		http.StatusForbidden,
		FormatError(
			"ERROR",
			"Forbidden",
			"",
			"",
			"csrf_token_mismatch",
			nil,
		),
	)
}
