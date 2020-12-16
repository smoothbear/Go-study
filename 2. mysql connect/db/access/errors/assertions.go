package errors

import (
	"errors"
	"fmt"
)

const (
	InterfaceAssertionErrorMessage = "DB Error : %s"
)

var (
	UserAssertionError = errors.New(fmt.Sprintf(InterfaceAssertionErrorMessage, "*model.user"))
	BookAssertionError = errors.New(fmt.Sprintf(InterfaceAssertionErrorMessage, "*model.book"))
)