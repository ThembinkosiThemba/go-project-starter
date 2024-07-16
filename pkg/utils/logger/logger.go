package logger

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sync"
	"time"
)

type LogEntry struct {
	Timestamp string `json:"timestamp"`
	File      string `json:"file"`
	Line      int    `json:"line"`
	Message   string `json:"message"`
	Status    string `json:"status"`
}

type Logger struct {
	errorFile string
	warnFile  string
	infoFile  string
	mutex     sync.Mutex
}

var instance *Logger
var once sync.Once

func InitLogger() {
	once.Do(func() {
		instance = NewLogger(
			"logs/errors.json",
			"logs/warnings.json",
			"logs/infos.json",
		)
	})
}

func NewLogger(errorFile, warnFile, infoFile string) *Logger {
	if err := os.MkdirAll(filepath.Dir(errorFile), 0755); err != nil {
		panic(err)
	}
	return &Logger{
		errorFile: errorFile,
		warnFile:  warnFile,
		infoFile:  infoFile,
	}
}

func logToFile(filename string, entry LogEntry) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	var logEntries []LogEntry
	fileInfo, _ := file.Stat()
	if fileInfo.Size() != 0 {
		decoder := json.NewDecoder(file)
		if err := decoder.Decode(&logEntries); err != nil && err.Error() != "EOF" {
			return err
		}
	}

	logEntries = append(logEntries, entry)

	file.Seek(0, 0)
	file.Truncate(0)
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(logEntries); err != nil {
		return err
	}

	return nil
}

// getCallerInfo returns the file and line number of the caller
func getCallerInfo() (string, int) {
	_, file, line, _ := runtime.Caller(2)
	return filepath.ToSlash(file), line
}

func printLogToTerminal(entry LogEntry) {
	fmt.Printf("[%s] %s: %s\n", entry.Timestamp, entry.Status, entry.Message)
}

func Error(err error) {
	if instance == nil {
		panic("Logger not initialized.")
	}
	instance.mutex.Lock()
	defer instance.mutex.Unlock()

	file, line := getCallerInfo()
	entry := LogEntry{
		Timestamp: time.Now().Format(time.RFC3339),
		File:      file,
		Line:      line,
		Message:   err.Error(),
		Status:    "ERROR",
	}
	printLogToTerminal(entry)
	logToFile(instance.errorFile, entry)
}

func Warn(message string) {
	if instance == nil {
		panic("Logger not initialized.")
	}
	instance.mutex.Lock()
	defer instance.mutex.Unlock()

	file, line := getCallerInfo()
	entry := LogEntry{
		Timestamp: time.Now().Format(time.RFC3339),
		File:      file,
		Line:      line,
		Message:   message,
		Status:    "WARN",
	}
	printLogToTerminal(entry)
	logToFile(instance.warnFile, entry)
}

func Info(message string) {
	if instance == nil {
		panic("Logger not initialized.")
	}
	instance.mutex.Lock()
	defer instance.mutex.Unlock()

	file, line := getCallerInfo()
	entry := LogEntry{
		Timestamp: time.Now().Format(time.RFC3339),
		File:      file,
		Line:      line,
		Message:   message,
		Status:    "INFO",
	}
	printLogToTerminal(entry)
	logToFile(instance.infoFile, entry)
}
