package interfaces

//go:generate mockgen -destination=mocks/logger_mock.go -package=logger . LoggerInterface
type ILogger interface {
	Error(args ...interface{})
	Trace(args ...interface{})
	Debug(args ...interface{})
	Info(args ...interface{})
	Print(args ...interface{})
	Warn(args ...interface{})
	Fatal(args ...interface{})
	Infof(format string, args ...interface{})
	Debugf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Warnf(format string, args ...interface{})
}
