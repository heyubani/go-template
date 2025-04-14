package interfaces

type IAppError interface {
	Error() string
	GetCode() int
	GetLabel() string
}
