package proxy

type Logger interface {
	Log(str string)
}

type RemoteLogger struct{}

func (*RemoteLogger) Log(str string) {

}

type FileLogger struct{}

func (*FileLogger) Log(str string) {
}

type LoggerProxy struct {
	logger Logger
}

func (lg *LoggerProxy) Log(str string) {
	lg.logger.Log(str)
}
