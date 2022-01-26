package log

var logger Logger

func init() {
	logger = New()
	logger.SetCalldepth(3)
}

func SetLevel(level Level) {
	logger.SetLevel(level)
}

func GetLevel() Level {
	return logger.GetLevel()
}

func Debugf(format string, msg ...interface{}) {
	logger.Debugf(format, msg...)
}

func Debugln(msg ...interface{}) {
	logger.Debugln(msg...)
}

func Debug(msg ...interface{}) {
	logger.Debug(msg...)
}

func Errorf(format string, msg ...interface{}) {
	logger.Errorf(format, msg...)
}

func Errorln(msg ...interface{}) {
	logger.Errorln(msg...)
}

func Error(msg ...interface{}) {
	if logger.gtlevel < ERROR {
		logger.Error(msg...)
	}
}

func Printf(format string, msg ...interface{}) {
	logger.Printf(format, msg...)
}

func Println(msg ...interface{}) {
	logger.Println(msg...)
}

func Prin(msg ...interface{}) {
	logger.Print(msg...)
}
