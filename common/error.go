package common

import "fmt"

type FetchCovidCaseError struct {
	errMsg string
}

func NewFetchCovidCaseError(message string) *FetchCovidCaseError {
	return &FetchCovidCaseError{errMsg: message}
}

func (e *FetchCovidCaseError) Error() string {
	return fmt.Sprintf("FetchCovidCaseError: %s", e.errMsg)
}
