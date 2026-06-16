package shortcut

import (
	"errors"
	"log"
)

var (
	ErrNotFound                   = errors.New("not found")
	ErrValidation                 = errors.New("validation error")
	ErrDuplicateKey               = errors.New("duplicate key")
	ErrForeignKeyViolation        = errors.New("foreign key violation")
	ErrCheckViolation             = errors.New("check violation")
	ErrExclusionViolation         = errors.New("exclusion violation")
	ErrUnspecifiedRequest         = errors.New("unspecified error")
	ErrInvalidRequest             = errors.New("invalid response")
	ErrUnspecifiedResponseGetUser = errors.New("unspecified response, nil user")
	ErrNotAllowedAge              = errors.New("not allowed")
	ErrFieldNotFilledName         = errors.New("field not filled name")
	ErrFieldNotFilledMail         = errors.New("field not filled mail")
	ErrFieldNotFilledPassword     = errors.New("field not filled password")
	ErrTokenNotFilled             = errors.New("token not filled")
)

func FatalIfErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
