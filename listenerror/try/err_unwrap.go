package try

import (
	"errors"
	"fmt"
)

type WrapError struct {
	msg string
	err   error
}

func (e *WrapError) Unwrap() error {
	return e.err
}

func (e *WrapError) Error() string {
	return  e.err.Error()
}

var rootError = errors.New("root error")

func ProduceErr() error {
	return &WrapError{err: rootError, msg: "wrap it"}
}

var e *WrapError

func first() error {
	err := &WrapError{err: ProduceErr(), msg: "wrap again"}
	if errors.As(err, &rootError) {
		fmt.Printf("error is: %+v", err)
	}
	return err
}

func StartErrUnwrap() {
	err := first()
	fmt.Errorf("err: %+v", err)
}