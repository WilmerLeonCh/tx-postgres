package domain

import "errors"

var (
	ErrMissingName            = errors.New("missing name")
	ErrMissingVersion         = errors.New("missing version")
	ErrMissingOwner           = errors.New("missing owner")
	ErrMissingManufactureDate = errors.New("missing manufacture date")
)
