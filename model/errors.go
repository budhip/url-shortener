package model

import "errors"

var (
	// ErrInternalServerError will throw if any the Internal Server Error happen
	ErrInternalServerError = errors.New("internal server error")
	// ErrUrlNotPresent will throw if the requested item is not exists
	ErrUrlNotPresent = errors.New("url is not present")
	// ErrSlugAlreadyUse will throw if the requested item is already use
	ErrSlugAlreadyUse = errors.New("the the desired shortcode is already in use. Shortcodes are case-sensitive")
	// ErrSlugNotMatch will throw if the requested item is not match with the regexp
	ErrSlugNotMatch = errors.New("the shortcode fails to meet the following regexp: ^[0-9a-zA-Z_]{6}$")
	// ErrSlugNotFound will throw if the requested item is not found
	ErrSlugNotFound = errors.New("the shortcode cannot be found in the system")
)

// ResponseError represent the reseponse error struct
type ResponseError struct {
	Message string `json:"msg"`
}
