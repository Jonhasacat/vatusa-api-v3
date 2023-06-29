package constants

type AuthenticationMethod int64

const (
	NoAuth AuthenticationMethod = iota
	Controller
	APIUser
)

const (
	FieldMethod     = "Method"
	FieldController = "Controller"
	FieldAPIUser    = "APIUser"
	FieldToken      = "Token"
)
