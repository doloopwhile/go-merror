# merror
--
    import "github.com/doloopwhile/go-merror"

Package merror provide MultipleError which help to collecting together multiple errors

## Usage

#### type MultipleError

```go
type MultipleError struct {
	Errors []error
}
```

MultipleError is collection of errors

#### func  Of

```go
func Of(errs []error) *MultipleError
```
Of is a factory of MultipleError. If errs contains non-nil errors, returns a
MultipleError of the errors. If errs contains nil's only or errs is nil, returns
nil

#### func (*MultipleError) Error

```go
func (m *MultipleError) Error() string
```
Error returns sub error messages they have been joined with semi-colons.

## License
MIT License. See LICENSE file.

## Contribution
I am looking for you pull request.

## Author
(doloopwhile)[https://github.com/doloopwhile/]
