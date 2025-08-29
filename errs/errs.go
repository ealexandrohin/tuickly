package errs

type ErrorMsg struct {
	Msg string
	Err error
}

func (e ErrorMsg) Error() string {
	if e.Err != nil {
		return e.Err.Error()
	}
	return e.Msg
}
