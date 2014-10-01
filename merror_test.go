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
	assert.Nil(Of([]error{}))
	assert.Nil(Of([]error{nil, nil, nil}))
	assert.Nil(Of(nil))

	// returns MultipleError without nil's
	e1 := errors.New("err1")
	e2 := errors.New("err2")
	e3 := errors.New("err3")
	m := Of([]error{e1, nil, e2, e3, nil})
	if assert.Len(m.Errors, 3) {
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
