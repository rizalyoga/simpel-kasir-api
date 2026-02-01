package errors

type ErrNotFound struct {
	Resource string
	ID       int
}

func (e *ErrNotFound) Error() string {
	return e.Resource + " not found"
}
