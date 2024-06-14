package response

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/lechuhuuha/oneshield/pkg/sperrors"
)

func HandleCommonError(ctx *gin.Context, err error) {
	unwrapable, ok := err.(interface{ Unwrap() []error })
	if ok {
		handleUnwrapError(ctx, unwrapable.Unwrap())
		return
	}

	var spErr *sperrors.AppError
	if errors.As(err, &spErr) {
		responseError(ctx, spErr, err.Error())
		return
	} else {
		InternalServerErrorResponse(ctx, err)
		return
	}
}

func handleUnwrapError(ctx *gin.Context, errs []error) {
	returnErr := &ErrorsResponse{
		Errors: make([]*ErrorResponse, 0, len(errs)),
	}

	for _, e := range errs {
		var spErr *sperrors.AppError
		if errors.As(e, &spErr) {
			returnErr.Errors = append(returnErr.Errors, &ErrorResponse{
				Source: Source{
					Field:     spErr.Field,
					Value:     spErr.Value,
					ErrorCode: spErr.Code,
				},
				Detail: spErr.String(),
			})
		} else {
			returnErr.Errors = append(returnErr.Errors,
				FormatError(
					"ERROR",
					"Internal Server Error",
					e.Error(),
					"",
					sperrors.CodeInternalServer,
					spErr.Value,
				),
			)
		}
	}

	HandleErrors(ctx, http.StatusBadRequest, returnErr)
}

func responseError(ctx *gin.Context, spErr *sperrors.AppError, detail string) {
	switch spErr.Code {
	case sperrors.CodeNotFound:
		NotFoundResponse(ctx)
	case sperrors.CodeNotAuthorized:
		HandleError(
			ctx,
			http.StatusUnauthorized,
			FormatError(
				"ERROR",
				"Not Authorized",
				detail,
				spErr.Field,
				spErr.Code,
				spErr.Value,
			),
		)
	case sperrors.CodePermission:
		HandleError(
			ctx,
			http.StatusForbidden,
			FormatError(
				"ERROR",
				"Forbidden",
				detail,
				spErr.Field,
				spErr.Code,
				spErr.Value,
			),
		)
	case sperrors.CodeDuplicate,
		sperrors.CodeAlreadyExists,
		sperrors.CodeStillInUse,
		sperrors.CodeExceedLimit,
		sperrors.CodeAlreadyChanged,
		sperrors.CodeRunning,
		sperrors.CodeDisabledUser,
		sperrors.CodeInvalidData,
		sperrors.CodeTokenExpired:
		HandleError(
			ctx,
			http.StatusBadRequest,
			FormatError2(
				"ERROR",
				"Bad Request",
				detail,
				spErr.Field,
				spErr.Code,
				spErr.Value,
				spErr.IsWarn,
			),
		)
	case sperrors.CodeNotHasEnoughQuota:
		HandleError(
			ctx,
			http.StatusBadRequest,
			FormatError(
				"ERROR",
				"Don't have enough quota for this user",
				detail,
				spErr.Field,
				sperrors.CodeNotHasEnoughQuota,
				spErr.Value,
			),
		)
	case sperrors.CodeInvalidPassword:
		HandleError(ctx, http.StatusUnauthorized,
			FormatError(
				"ERROR",
				"Username or password is not correct",
				"",
				"BODY",
				"invalid_username_or_password",
				spErr.Value,
			),
		)
	default:
		HandleError(
			ctx,
			http.StatusInternalServerError,
			FormatError(
				"ERROR",
				"Internal Server Error",
				detail,
				spErr.Field,
				sperrors.CodeInternalServer,
				spErr.Value,
			),
		)
	}
}
