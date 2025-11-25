package apperror

type NoDataError struct{}

func (e *NoDataError) Error() string {
	return "no data available"
}

func NewNoDataError() *NoDataError {
	return &NoDataError{}
}
