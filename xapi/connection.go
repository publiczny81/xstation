package xapi

import (
	"io"
)

type connection interface {
	io.Closer
	Send(msg any) error
	Receive(msg any) error
}

type Connection struct {
}
