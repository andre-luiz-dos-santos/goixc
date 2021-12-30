package goixc

import (
	"errors"
	"fmt"
)

var (
	ErrInvalid     = errors.New("invalid data")
	ErrNotFound    = errors.New("not found")
	ErrForbidden   = errors.New("forbidden")
	ErrAlreadyDone = errors.New("already done")
)

type TokenError struct {
	Token string
}

func (e TokenError) Error() string {
	return "invalid IXC token: " + e.Token
}

type InvalidJSONError struct {
	JSON []byte
	Err  error
}

func (e InvalidJSONError) Error() string {
	return fmt.Sprintf("invalid JSON: %v: %.120q", e.Err, e.JSON)
}

func (e InvalidJSONError) Unwrap() error {
	return e.Err
}

type InvalidPDFError struct {
	PDF []byte
}

func (e InvalidPDFError) Error() string {
	return fmt.Sprintf("invalid PDF: %.120q", e.PDF)
}

func (e InvalidPDFError) Unwrap() error {
	return ErrInvalid
}

type NotFoundError struct {
	Resource string
	Login    string
}

func (e NotFoundError) Error() string {
	return fmt.Sprintf("%v not found: %v", e.Resource, e.Login)
}

func (e NotFoundError) Unwrap() error {
	return ErrNotFound
}

type ContratoError struct {
	ContratoID int64
	Mensagem   string
	Err        error
}

func (e ContratoError) Error() string {
	return fmt.Sprintf("contrato ID %v: %v: %v", e.ContratoID, e.Mensagem, e.Err)
}

func (e ContratoError) Unwrap() error {
	return e.Err
}

type IXCFormError struct {
	Form    string
	Args    interface{}
	Message string
}

func (e IXCFormError) Error() string {
	return fmt.Sprintf("IXC error in %v(%+v): %v", e.Form, e.Args, e.Message)
}
