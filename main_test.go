package main

import (
	"errors"
	"testing"
)

func TestDivide(*testing.T) {
	_, err := divide(100.2, 1.0)
	if err != nil {
		errors.New("err")
	}
}
