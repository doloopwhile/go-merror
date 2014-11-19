package merror

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultipleError(t *testing.T) {
	assert := assert.New(t)

	// returns nil for empty slices
	//
	// Returned value must be determined equal to nil by '=='.
	// It is because Golang has a non-intuitive behavior on comparison with nil.
	// seel http://golang.org/doc/faq#nil_error
	var err error
	err = Of([]error{})
	assert.True(err == nil, "%#V", err)

	err = Of([]error{nil, nil, nil})
	assert.True(err == nil, "%#V", err)

	err = Of(nil)
	assert.True(err == nil, "%#V", err)

	// returns MultipleError without nil's
	e1 := errors.New("err1")
	e2 := errors.New("err2")
	e3 := errors.New("err3")
	m, ok := Of([]error{e1, nil, e2, e3, nil}).(*MultipleError)
	if assert.True(ok) && assert.Len(m.Errors, 3) {
		assert.Equal(m.Errors[0], e1)
		assert.Equal(m.Errors[1], e2)
		assert.Equal(m.Errors[2], e3)
	}

	assert.Equal(m.Error(), "[err1;err2;err3]")
	assert.Equal(fmt.Sprint(m), "[err1;err2;err3]")

	// If nil's are contained, they are skipped
	m = &MultipleError{Errors: []error{e1, nil, e2, e3, nil}}
	assert.Equal(m.Error(), "[err1;err2;err3]")
	assert.Equal(fmt.Sprint(m), "[err1;err2;err3]")
}
