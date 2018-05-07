package lib

// Error exported to communicate between modules
type Error = struct {
	Code    int
	Message string
}
