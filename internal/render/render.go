package render

import (
	"context"
	"errors"
	"net/http"

	"gate-way/pkg/shortcut"

	"github.com/labstack/echo/v4"
)

const (
	CodeBadRequest          = "BAD_REQUEST"
	CodeValidationError     = "VALIDATION_ERROR"
	CodeUnauthorized        = "UNAUTHORIZED"
	CodeForbidden           = "FORBIDDEN"
	CodeNotFound            = "NOT_FOUND"
	CodeConflict            = "CONFLICT"
	CodeUnprocessableEntity = "UNPROCESSABLE_ENTITY"
	CodeTooManyRequests     = "TOO_MANY_REQUESTS"
	CodeInternalError       = "INTERNAL_ERROR"
	CodeServiceUnavailable  = "SERVICE_UNAVAILABLE"
)

const (
	MsgBadRequest          = "Bad request"
	MsgValidationError     = "Validation failed"
	MsgUnauthorized        = "Unauthorized"
	MsgForbidden           = "Forbidden"
	MsgNotFound            = "Resource not found"
	MsgConflict            = "Conflict"
	MsgUnprocessableEntity = "Unprocessable entity"
	MsgTooManyRequests     = "Too many requests"
	MsgInternalError       = "Internal server error"
	MsgServiceUnavailable  = "Service unavailable"
)

type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}

type HTTPError struct {
	Status    int
	Code      string
	Message   string
	Err       error
	ExposeErr bool
}

func (e *HTTPError) Error() string {
	if e.Err != nil {
		return e.Err.Error()
	}

	return e.Message
}

func (e *HTTPError) Unwrap() error {
	return e.Err
}

func JSON(c echo.Context, status int, body any) error {
	return c.JSON(status, body)
}

func New(
	status int,
	code string,
	message string,
	err error,
	exposeErr bool,
) error {
	return &HTTPError{
		Status:    status,
		Code:      code,
		Message:   message,
		Err:       err,
		ExposeErr: exposeErr,
	}
}

func NewBadRequest(err error) error {
	return New(http.StatusBadRequest, CodeBadRequest, MsgBadRequest, err, true)
}

func NewValidationError(err error) error {
	return New(http.StatusBadRequest, CodeValidationError, MsgValidationError, err, true)
}

func NewUnauthorized(err error) error {
	return New(http.StatusUnauthorized, CodeUnauthorized, MsgUnauthorized, err, false)
}

func NewForbidden(err error) error {
	return New(http.StatusForbidden, CodeForbidden, MsgForbidden, err, false)
}

func NewNotFound(err error) error {
	return New(http.StatusNotFound, CodeNotFound, MsgNotFound, err, false)
}

func NewConflict(err error) error {
	return New(http.StatusConflict, CodeConflict, MsgConflict, err, true)
}

func NewUnprocessableEntity(err error) error {
	return New(http.StatusUnprocessableEntity, CodeUnprocessableEntity, MsgUnprocessableEntity, err, true)
}

func NewTooManyRequests(err error) error {
	return New(http.StatusTooManyRequests, CodeTooManyRequests, MsgTooManyRequests, err, false)
}

func NewInternal(err error) error {
	return New(http.StatusInternalServerError, CodeInternalError, MsgInternalError, err, false)
}

func NewServiceUnavailable(err error) error {
	return New(http.StatusServiceUnavailable, CodeServiceUnavailable, MsgServiceUnavailable, err, false)
}

func BadRequest(c echo.Context, err error) error {
	return writeError(c, http.StatusBadRequest, CodeBadRequest, MsgBadRequest, err, true)
}

func ValidationError(c echo.Context, err error) error {
	return writeError(c, http.StatusBadRequest, CodeValidationError, MsgValidationError, err, true)
}

func Unauthorized(c echo.Context, err error) error {
	return writeError(c, http.StatusUnauthorized, CodeUnauthorized, MsgUnauthorized, err, false)
}

func Forbidden(c echo.Context, err error) error {
	return writeError(c, http.StatusForbidden, CodeForbidden, MsgForbidden, err, false)
}

func NotFound(c echo.Context, err error) error {
	return writeError(c, http.StatusNotFound, CodeNotFound, MsgNotFound, err, false)
}

func Conflict(c echo.Context, err error) error {
	return writeError(c, http.StatusConflict, CodeConflict, MsgConflict, err, true)
}

func UnprocessableEntity(c echo.Context, err error) error {
	return writeError(c, http.StatusUnprocessableEntity, CodeUnprocessableEntity, MsgUnprocessableEntity, err, true)
}

func TooManyRequests(c echo.Context, err error) error {
	return writeError(c, http.StatusTooManyRequests, CodeTooManyRequests, MsgTooManyRequests, err, false)
}

func Internal(c echo.Context, err error) error {
	return writeError(c, http.StatusInternalServerError, CodeInternalError, MsgInternalError, err, false)
}

func ServiceUnavailable(c echo.Context, err error) error {
	return writeError(c, http.StatusServiceUnavailable, CodeServiceUnavailable, MsgServiceUnavailable, err, false)
}

func writeError(
	c echo.Context,
	status int,
	code string,
	message string,
	err error,
	exposeErr bool,
) error {
	if c.Response().Committed {
		return nil
	}

	resp := ErrorResponse{
		Code:    code,
		Message: message,
	}

	if err != nil && exposeErr {
		resp.Error = err.Error()
	}

	return c.JSON(status, resp)
}

func ErrorHandler(err error, c echo.Context) {
	if c.Response().Committed {
		return
	}

	var httpErr *HTTPError
	if errors.As(err, &httpErr) {
		_ = writeError(
			c,
			httpErr.Status,
			httpErr.Code,
			httpErr.Message,
			httpErr.Err,
			httpErr.ExposeErr,
		)
		return
	}

	var echoErr *echo.HTTPError
	if errors.As(err, &echoErr) {
		_ = writeError(
			c,
			echoErr.Code,
			CodeBadRequest,
			MsgBadRequest,
			err,
			true,
		)
		return
	}

	_ = Internal(c, err)
}

func FromError(c echo.Context, err error) error {
	if err == nil {
		return nil
	}

	switch {
	case errors.Is(err, shortcut.ErrFieldNotFilledName),
		errors.Is(err, shortcut.ErrFieldNotFilledMail),
		errors.Is(err, shortcut.ErrFieldNotFilledPassword),
		errors.Is(err, shortcut.ErrValidation),
		errors.Is(err, shortcut.ErrABValidation),
		errors.Is(err, shortcut.ErrABExperimentNameRequired),
		errors.Is(err, shortcut.ErrABExperimentLayersRequired),
		errors.Is(err, shortcut.ErrABExperimentGroupsMinCount),
		errors.Is(err, shortcut.ErrABExperimentStartDateAfterEnd),
		errors.Is(err, shortcut.ErrABExperimentRolloutOutOfRange),
		errors.Is(err, shortcut.ErrABNamespaceNameRequired),
		errors.Is(err, shortcut.ErrABLayerNameRequired),
		errors.Is(err, shortcut.ErrABLayerNamespaceRequired),
		errors.Is(err, shortcut.ErrABCustomParamNameRequired),
		errors.Is(err, shortcut.ErrABCustomParamNamespaceRequired),
		errors.Is(err, shortcut.ErrABCustomParamTypeRequired),
		errors.Is(err, shortcut.ErrABExperimentIDRequired),
		errors.Is(err, shortcut.ErrABNamespaceRequired),
		errors.Is(err, shortcut.ErrABSplitIDRequired):
		return ValidationError(c, err)

	case errors.Is(err, shortcut.ErrInvalidRequest),
		errors.Is(err, shortcut.ErrABInvalidRequest):
		return BadRequest(c, err)

	case errors.Is(err, shortcut.ErrTokenNotFilled):
		return Unauthorized(c, err)

	case errors.Is(err, shortcut.ErrNotAllowedAge):
		return Forbidden(c, err)

	case errors.Is(err, shortcut.ErrNotFound):
		return NotFound(c, err)

	case errors.Is(err, shortcut.ErrDuplicateKey):
		return Conflict(c, err)

	case errors.Is(err, shortcut.ErrForeignKeyViolation),
		errors.Is(err, shortcut.ErrCheckViolation),
		errors.Is(err, shortcut.ErrExclusionViolation):
		return UnprocessableEntity(c, err)

	case errors.Is(err, context.DeadlineExceeded),
		errors.Is(err, context.Canceled),
		errors.Is(err, shortcut.ErrABEmptyResponse):
		return ServiceUnavailable(c, err)

	case errors.Is(err, shortcut.ErrUnspecifiedRequest),
		errors.Is(err, shortcut.ErrUnspecifiedResponseGetUser),
		errors.Is(err, shortcut.ErrABUnspecified):
		return Internal(c, err)

	default:
		return Internal(c, err)
	}
}
