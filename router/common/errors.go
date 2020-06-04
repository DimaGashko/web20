package common

type AppError struct {
	Err  error
	Code int
}

func (e AppError) Error() string {
	return e.Err.Error()
}
