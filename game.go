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
	Init() error
	Start() error
	Update(msg string) error
	Pause() error
	Finish() error
	Shutdown() error
}
