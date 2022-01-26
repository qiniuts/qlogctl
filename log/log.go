package log

import (
	"fmt"
	stdLog "log"
	"os"
)

type Level int

const (
	_             = iota
	VERBOSE Level = iota
	DEBUG
	INFO
	WARN
	ERROR
	NONE
)

type Logger struct {
	calldepth int
	gtlevel   Level
	debug     *stdLog.Logger
	err       *stdLog.Logger

	normal *stdLog.Logger
}

func New() Logger {
	debug := stdLog.New(os.Stderr,
		"[DEBUG]",
		stdLog.Ldate|stdLog.Ltime|stdLog.Lshortfile)

	err := stdLog.New(os.Stderr,
		"[ERROR]",
		stdLog.Ldate|stdLog.Ltime|stdLog.Lshortfile)

	normal := stdLog.New(os.Stderr, "", 0)

	return Logger{
		gtlevel:   ERROR - 1,
		calldepth: 2,
		debug:     debug,
		err:       err,
		normal:    normal}
}

func (l *Logger) SetLevel(level Level) {
	l.gtlevel = level - 1
}

func (l *Logger) GetLevel() Level {
	return l.gtlevel + 1
}

// SetCalldepth Calldepth is used to recover the PC and is
// provided for generality
func (l *Logger) SetCalldepth(calldepth int) {
	l.calldepth = calldepth
}

func (l *Logger) GetCalldepth() int {
	return l.calldepth
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	if l.gtlevel < DEBUG {
		l.debug.Output(l.calldepth, fmt.Sprintf(format, v...))
	}
}

func (l *Logger) Debugln(v ...interface{}) {
	if l.gtlevel < DEBUG {
		l.debug.Output(l.calldepth, fmt.Sprintln(v...))
	}
}

func (l *Logger) Debug(v ...interface{}) {
	if l.gtlevel < DEBUG {
		l.debug.Output(l.calldepth, fmt.Sprint(v...))
	}
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	if l.gtlevel >= ERROR {
		l.debug.Output(l.calldepth, fmt.Sprintf(format, v...))
	}
}

func (l *Logger) Errorln(v ...interface{}) {
	if l.gtlevel >= ERROR {
		l.debug.Output(l.calldepth, fmt.Sprintln(v...))
	}
}

func (l *Logger) Error(v ...interface{}) {
	if l.gtlevel >= ERROR {
		l.debug.Output(l.calldepth, fmt.Sprint(v...))
	}
}

func (l *Logger) Printf(format string, v ...interface{}) {
	l.normal.Output(l.calldepth, fmt.Sprintf(format, v...))
}

func (l *Logger) Println(v ...interface{}) {
	l.normal.Output(l.calldepth, fmt.Sprintln(v...))
}

func (l *Logger) Print(v ...interface{}) {
	l.normal.Output(l.calldepth, fmt.Sprint(v...))
}
