package log

import (
	"fmt"
	"os"
)

type Logger struct {
	ServiceName string
	hostname    string
}

func New() *Logger {
	hostname, _ := os.Hostname()
	return &Logger{
		ServiceName: "",
		hostname:    hostname,
	}
}

func (l *Logger) SetServiceName(name string) {
	l.ServiceName = name
}

func (l *Logger) LogInfo(msg string) {
	fmt.Printf("\n[INFO] (%s/%s) %s", l.ServiceName, l.hostname, msg)
}

func (l *Logger) LogWarn(msg string) {
	fmt.Printf("\n[WARN] (%s/%s) %s", l.ServiceName, l.hostname, msg)
}

func (l *Logger) LogErr(msg string) {
	fmt.Printf("\n[ERROR] (%s/%s) %s", l.ServiceName, l.hostname, msg)
}
