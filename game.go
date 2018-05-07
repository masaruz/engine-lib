package lib

// Form to comunicate with plugin
type Form = struct {
	Header    string
	Body      string
	Signature string
}

type ack = func(form Form)

// Game interface engine
type Game interface {
	On(name string, ack ack)
	Start() *Error
	Pause() *Error
	Finish() *Error
}
