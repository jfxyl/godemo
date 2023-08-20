package logger

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"time"
)

var (
	logDir      string
	infoLogger  *log.Logger
	debugLogger *log.Logger
	outFile     *os.File
	logLevel    int
	currentDay  int
)

const (
	DebugLevel = iota
	InfoLevel
)

func SetFile(path string) {
	var (
		err error
		now time.Time
	)
	logDir = path
	now = time.Now()
	ymd := now.Format("2006-01-02")

	if _, err = os.Stat(logDir); os.IsNotExist(err) {
		os.MkdirAll(logDir, 0755)
	}
	logPath := fmt.Sprintf("%s/%s.log", logDir, ymd)
	if outFile, err = os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644); err != nil {
		panic(err)
	}
	currentDay = now.YearDay()
	fmt.Println("currentDay:", currentDay)
	infoLogger = log.New(outFile, "[INFO] ", log.LstdFlags)
	debugLogger = log.New(outFile, "[DEBUG] ", log.LstdFlags)
}

func SetLevel(level int) {
	logLevel = level
}

func Debug(format string, v ...any) {
	if logLevel <= DebugLevel {
		checkIfDayChanged()
		debugLogger.Printf(getPrefix()+format, v...)
	}
}

func Info(format string, v ...any) {
	if logLevel <= InfoLevel {
		checkIfDayChanged()
		infoLogger.Printf(getPrefix()+format, v...)
	}
}

func checkIfDayChanged() {
	if time.Now().YearDay() != currentDay {
		outFile.Close()
		SetFile(logDir)
	}
}

func getCallTrace() (string, int, uintptr) {
	function, filename, line, ok := runtime.Caller(3)
	if ok {
		return filename, line, function
	}
	return "", 0, 0
}

func getPrefix() string {
	filename, line, function := getCallTrace()
	return fmt.Sprintf("%s:%d:%s ", filename, line, runtime.FuncForPC(function).Name())
}
