package goixc

import (
	"errors"
	"fmt"
)

var (
	ErrBad       = errors.New("invalid data")
	ErrNotFound  = errors.New("not found")
	ErrForbidden = errors.New("forbidden")
)

type ErrBadToken struct {
	Token string
}

func (e ErrBadToken) Error() string {
	return "invalid IXC token: " + e.Token
}

type ErrBadJSON struct {
	Bytes []byte
	Err   error
}

func (e ErrBadJSON) Error() string {
	return fmt.Sprintf("invalid JSON: %v: %.120q", e.Err, e.Bytes)
}

func (e ErrBadJSON) Unwrap() error {
	return e.Err
}

type ErrBadPDF struct {
	PDF []byte
}

func (e ErrBadPDF) Error() string {
	return fmt.Sprintf("invalid PDF: %.120q", e.PDF)
}

func (e ErrBadPDF) Unwrap() error {
	return ErrBad
}

type ErrLoginNotFound struct {
	Login string
}

func (e ErrLoginNotFound) Error() string {
	return fmt.Sprintf("login not found: %v", e.Login)
}

func (e ErrLoginNotFound) Unwrap() error {
	return ErrNotFound
}

type ErrContrato struct {
	ContratoID int64
	Mensagem   string
	Err        error
}

func (e ErrContrato) Error() string {
	return fmt.Sprintf("contrato ID %v: %v", e.ContratoID, e.Mensagem)
}

func (e ErrContrato) Unwrap() error {
	return e.Err
}

type ErrIXCForm struct {
	Form    string
	JSON    interface{}
	Message string
}

func (e ErrIXCForm) Error() string {
	return fmt.Sprintf("IXC error in %v(%v): %v", e.Form, e.JSON, e.Message)
}
