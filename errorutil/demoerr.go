package main

import (
	"errors"
	"fmt"
)

var ErrBoom = errors.New("boom!")

type FancyError struct {
	Code int
	Err  error
}

func (e *FancyError) Error() string {
	return fmt.Sprintf("code %d: %v", e.Code, e.Err)
}

func (e *FancyError) Unwrap() error {
	return e.Err
}

func customfailed() error {
	return &FancyError{Code: 404, Err: fmt.Errorf("not found: %w", ErrBoom)}
}

func fail() error {
	return fmt.Errorf("something went wrong: %w", ErrBoom)
}

func main() {
	err := fail()
	fmt.Println("err:", err)
	if errors.Is(err, ErrBoom) {
		fmt.Println("errors.Is: matched ErrBoom!")
	} else {
		fmt.Println("errors.Is: did not match.")
	}

	unwrapped := errors.Unwrap(err)
	fmt.Println("Unwrapped:", unwrapped)

	err2 := customfailed()
	var fe *FancyError
	if errors.As(err2, &fe) {
		fmt.Println("errors.As: matched FancyError with code", fe.Code)
	} else {
		fmt.Println("errors.As: did not match.")
	}
}
