package errors

// Bad given data
type ErrorData struct {
	Msg string `json:"errors"`
}

func (e ErrorData) Error() string {
	return e.Msg
}

// Server work error
type ErrorServer struct {
	Msg string `json:"errors"`
}

func (e ErrorServer) Error() string {
	return e.Msg
}

// No authentication
type ErrorNoAuth struct {
	Msg string `json:"errors"`
}

func (e ErrorNoAuth) Error() string {
	return e.Msg
}

// All is alright
// no error
type OK struct {
	Msg string `json:"message"`
}

func (e OK) Error() string {
	return e.Msg
}

// Requested object not found
type ErrorNotFound struct {
	Msg string `json:"errors"`
}

func (e ErrorNotFound) Error() string {
	return e.Msg
}

// Object created
type Created struct {
	Msg string `json:"message"`
}

func (e Created) Error() string {
	return e.Msg
}
