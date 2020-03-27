package errors

// Bill18Error .
type Bill18Error string

func (ee Bill18Error) Error() string {
	return string(ee)
}

var (
	// ErrBadDBConfiguration .
	ErrBadDBConfiguration = Bill18Error("Не заполнены входные параметры БД")
)
