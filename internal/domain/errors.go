package domain

type Err struct {
	Code    int
	Message string
	Cause   error
}
