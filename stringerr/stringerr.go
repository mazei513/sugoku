package stringerr

// StringErr is an error type which is just a string
type StringErr string

func (s StringErr) Error() string {
	return string(s)
}
