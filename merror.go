// Package merror provide MultipleError which help to collecting together multiple errors
package merror

import (
	"strings"
)

// MultipleError is collection of errors
type MultipleError struct {
	Errors []error
}

// Of is a factory of MultipleError.
// If errs contains non-nil errors, returns a MultipleError of the errors.
// If errs contains nil's only or errs is nil, returns nil
func Of(errs []error) error {
	m := &MultipleError{}
	for _, err := range errs {
		if err != nil {
			m.Errors = append(m.Errors, err)
		}
	}
	if len(m.Errors) == 0 {
		return nil
	}
	return m
}

// Error returns sub error messages they have been joined with semi-colons.
func (m *MultipleError) Error() string {
	msgs := []string{}
	for _, err := range m.Errors {
		if err != nil {
			msgs = append(msgs, err.Error())
		}
	}
	return `[` + strings.Join(msgs, `;`) + `]`
}
