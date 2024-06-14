package sperrors

import (
	"errors"
	"fmt"
)

const (
	CodeServerStopped         string = "server_stopped"
	CodeInternalServer        string = "internal_server"
	CodeNotAuthorized         string = "not_authorized"
	CodeNotFound              string = "not_found"
	CodeDuplicate             string = "duplicate"
	CodeResourceBusy          string = "resource_busy"
	CodeResourceAlreadyExists string = "resource_already_exists"
	CodeSymLinkAttack         string = "symlink_attack"
	CodeInvalidPassword       string = "invalid_password"
	CodePermission            string = "permission"
	CodeInvalidJWTToken       string = "invalid_jwt_token" //nolint:gosec
	CodeAlreadyExists         string = "already_exists"
	CodeOwnerRole             string = "owner_role"
	CodeTokenExpired          string = "token_expired"
	CodeAlreadyChanged        string = "already_changed"
	CodeNotHasEnoughQuota     string = "not_has_enough_quota"
	CodeStillHasUser          string = "still_has_user"
	CodeNotInPackage          string = "not_in_package"
	CodeSamePassword          string = "same_password"
	CodeDisabledUser          string = "disabled_user"
	CodeInvalidPath           string = "invalid_path"
	CodeIPIsNotAllow          string = "ip_is_not_allowed"
	CodeDeleteDefault         string = "delete_default"
	CodeInvalidCertificate    string = "invalid_certificate"
	CodeNotSupportedAlias     string = "not_supported_alias"
	CodeNotEndUser            string = "not_end_user"
	CodeHasChildUser          string = "has_child_user"
	CodeNotAllowedFunction    string = "not_allowed_function"
	CodeNotSupported          string = "not_supported"
	CodeInvalidSSHKey         string = "invalid_ssh_key"
	CodeExceedLimit           string = "exceed_limit"
	CodeTooLong               string = "too_long"
	CodeTooShort              string = "too_short"
	CodeInvalidData           string = "invalid_data"
	CodeBackupAlreadyRunning  string = "backup_already_running"
	CodeNoDefaultDestination  string = "no_default_destination"
	CodeStillInUse            string = "still_in_use"
	CodeUserDisabledByAdmin   string = "user_disabled_by_admin"
	CodePackageNotSuitable    string = "package_not_suitable"
	CodeRunning               string = "running"
	CodeFileTooLarge          string = "file_too_large"
	CodeRequestTimeout        string = "request_timeout"
	CodeTooManyRequests       string = "too_many_requests"
	CodeNotEnoughDiskOrInode  string = "not_enough_disk_or_inode"
)

var (
	ErrServerStopped         = NewAppError(CodeServerStopped, "server stopped")
	ErrNotFound              = NewAppError(CodeNotFound, "not found")
	ErrDuplicate             = NewAppError(CodeDuplicate, "duplicate")
	ErrResourceBusy          = NewAppError(CodeResourceBusy, "resource busy")
	ErrResourceAlreadyExists = NewAppError(CodeResourceAlreadyExists, "resource already exists")
	ErrSymLinkAttack         = NewAppError(CodeSymLinkAttack, "symbolic link attack")
	ErrInvalidPassword       = NewAppError(CodeInvalidPassword, "invalid password")
	ErrPermission            = NewAppError(CodePermission, "permission denied")
	ErrInvalidJWTToken       = NewAppError(CodeInvalidJWTToken, "invalid jwt token")
	ErrAlreadyExists         = NewAppError(CodeAlreadyExists, "already exists")
	ErrOwnerRole             = NewAppError(CodeOwnerRole, "owner role is not allowed")
	ErrTokenExpired          = NewAppError(CodeTokenExpired, "token expired")
	ErrAlreadyChanged        = NewAppError(CodeAlreadyChanged, "already changed")
	ErrNotHasEnoughQuota     = NewAppError(CodeNotHasEnoughQuota, "not has enough quota")
	ErrStillHasUser          = NewAppError(CodeStillHasUser, "still has user")
	ErrNotInPackage          = NewAppError(CodeNotInPackage, "not in package")
	ErrSamePassword          = NewAppError(CodeSamePassword, "same password")
	ErrDisabledUser          = NewAppError(CodeDisabledUser, "deactivated user")
	ErrInvalidPath           = NewAppError(CodeInvalidPath, "invalid path")
	ErrDeleteDefault         = NewAppError(CodeDeleteDefault, "can not delete default package")
	ErrInvalidCertificate    = NewAppError(CodeInvalidCertificate, "invalid certificate")
	ErrNotEndUser            = NewAppError(CodeNotEndUser, "not end user")
	ErrHasChildUser          = NewAppError(CodeHasChildUser, "has child user")
	ErrNotAllowedFunction    = NewAppError(CodeNotAllowedFunction, "not allowed function")
	ErrNotSupported          = NewAppError(CodeNotSupported, "not supported")
	ErrInvalidSSHKey         = NewAppError(CodeInvalidSSHKey, "invalid ssh key")
	ErrExceedLimit           = NewAppError(CodeExceedLimit, "exceed limit")
	ErrTooLong               = NewAppError(CodeTooLong, "too long")
	ErrTooShort              = NewAppError(CodeTooShort, "too short")
	ErrInvalidData           = NewAppError(CodeInvalidData, "invalid data")
	ErrBackupAlreadyRunning  = NewAppError(CodeBackupAlreadyRunning, "backup already running")
	ErrNoDefaultDestination  = NewAppError(CodeNoDefaultDestination, "no default destination")
	ErrStillInUse            = NewAppError(CodeStillInUse, "still in use")
	ErrUserDisabledByAdmin   = NewAppError(CodeUserDisabledByAdmin, "disabled by admin")
	ErrPackageNotSuitable    = NewAppError(
		CodePackageNotSuitable,
		"Your package is not suitable for this function. Please contact admin to upgrade your package.",
	)
	ErrNotAuthorized        = NewAppError(CodeNotAuthorized, "not authorized")
	ErrRunning              = NewAppError(CodeRunning, "running")
	ErrFileTooLarge         = NewAppError(CodeFileTooLarge, "file too large")
	ErrNotEnoughDiskOrInode = NewAppError(CodeNotEnoughDiskOrInode, "not enough disk or inode")
)

type AppError struct {
	Code    string
	Message string
	Field   string
	Value   any
	IsWarn  bool
}

func NewAppError(code, msg string) *AppError {
	return &AppError{Code: code, Message: msg}
}

// Is checks if the given error is a apperror with the given code.
func (s *AppError) Is(err error) bool {
	var serr *AppError
	if errors.As(err, &serr) {
		if serr.Code == s.Code {
			return true
		}
	}
	return false
}

func (s *AppError) String() string {
	return s.Error()
}

func (s *AppError) Error() string {
	if s.Field == "" {
		return s.Message
	}

	if s.Value == nil {
		return fmt.Sprintf("[%s] where %s", s.Message, s.Field)
	}

	if s.Field == "message" {
		return fmt.Sprintf("%v", s.Value)
	}

	return fmt.Sprintf("[%s] where %s is %v", s.Message, s.Field, s.Value)
}

func (s *AppError) RuntimeError() {
}

// WithField returns a new error with the given field set.
func (s *AppError) WithField(field string, args ...any) *AppError {
	ss := *s

	if len(args) > 0 {
		ss.Value = args[0]
	}
	ss.Field = field
	return &ss
}

// WithMsg returns a new error with the given message.
func (s *AppError) WithMsg(msg string) *AppError {
	return s.WithField("message", msg)
}

// / WithMsgf returns a new error with the given message.
func (s *AppError) WithMsgf(format string, args ...any) *AppError {
	return s.WithMsg(fmt.Sprintf(format, args...))
}

func (s *AppError) Warn() *AppError {
	ss := *s

	ss.IsWarn = true

	return &ss
}
