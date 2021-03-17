package logging

import (
	"fmt"
	"os"

	"github.com/tinywell/fabtool/pkg/core"
)

type log struct {
}

func (l *log) Fatal(v ...interface{}) {
	fmt.Println(v...)
	os.Exit(0)
}

func (l *log) Fatalf(format string, v ...interface{}) {
	fmt.Printf(format, v...)
	os.Exit(0)
}

func (l *log) Fatalln(v ...interface{}) {
	panic("not implemented") // TODO: Implement
}

func (l *log) Panic(v ...interface{}) {
	panic("not implemented") // TODO: Implement
}

func (l *log) Panicf(format string, v ...interface{}) {
	panic("not implemented") // TODO: Implement
}

func (l *log) Panicln(v ...interface{}) {
	panic("not implemented") // TODO: Implement
}

func (l *log) Print(v ...interface{}) {
	panic("not implemented") // TODO: Implement
}

func (l *log) Printf(format string, v ...interface{}) {
	panic("not implemented") // TODO: Implement
}

func (l *log) Println(v ...interface{}) {
	panic("not implemented") // TODO: Implement
}

func (l *log) Debug(args ...interface{}) {
	fmt.Println(args...)
}

func (l *log) Debugf(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}

func (l *log) Debugln(args ...interface{}) {
	fmt.Println(args...)
}

func (l *log) Info(args ...interface{}) {
	fmt.Println(args...)
}

func (l *log) Infof(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}

func (l *log) Infoln(args ...interface{}) {
	fmt.Println(args...)
}

func (l *log) Warn(args ...interface{}) {
	panic("not implemented") // TODO: Implement
}

func (l *log) Warnf(format string, args ...interface{}) {
	panic("not implemented") // TODO: Implement
}

func (l *log) Warnln(args ...interface{}) {
	panic("not implemented") // TODO: Implement
}

func (l *log) Error(args ...interface{}) {
	panic("not implemented") // TODO: Implement
}

func (l *log) Errorf(format string, args ...interface{}) {
	panic("not implemented") // TODO: Implement
}

func (l *log) Errorln(args ...interface{}) {
	panic("not implemented") // TODO: Implement
}

// DefaultLogger ...
func DefaultLogger() core.Logger {
	return &log{}
}
