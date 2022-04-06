package models

import "fmt"

type Lending struct {
	ID         int64      `json:"id"`
	User       User       `json:"user"`
	Book       Book       `json:"book"`
	LendDay    string     `json:"lendDay"`
	Returned   int        `json:"returned"`
	Devolution Devolution `json:"devolution"`
}

type AlreadyLendingError struct {
	Lending string
}

func (e *AlreadyLendingError) Error() string {
	return fmt.Sprintf("%s already have lending", e.Lending)
}

func ErrorAlreadyLending(lending string) *AlreadyLendingError {
	return &AlreadyLendingError{Lending: lending}
}
