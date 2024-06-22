package errors

import "github.com/pkg/errors"

var (
	RequestSizeExceededError = errors.New("request size exceeded")
	InvalidContext = errors.New("invalid context")
)
