package progressbar

import "errors"

var (
	ErrNilBar       = errors.New("go-progressbar:progressbar is nil")
	ErrInvalidTotal = errors.New("go-progressbar:total must be greater than 0")
)
