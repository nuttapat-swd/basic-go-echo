package school

import "errors"

var (
	ErrNotFound      = errors.New("record not found")
	ErrSchoolInvalid = errors.New("invalid school reference")
	ErrForbidden     = errors.New("action not allowed")
)
