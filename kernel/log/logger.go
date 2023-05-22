package log

import (
	"github.com/gogf/gf/v2/os/glog"
)

// Logger
type Logger struct {
	*glog.Logger
	// LogPath specifies the directory for storing logging files.
	LogPath string

	// LogStdout specifies whether printing logging content to stdout.
	LogStdout bool

	// ErrorStack specifies whether logging stack information when error.
	ErrorStack bool

	// ErrorLogEnabled enables error logging content to files.
	ErrorLogEnabled bool

	// ErrorLogPattern specifies the error log file pattern like: error-{Ymd}.log
	ErrorLogPattern string

	// AccessLogEnabled enables access logging content to files.
	AccessLogEnabled bool

	// AccessLogPattern specifies the error log file pattern like: access-{Ymd}.log
	AccessLogPattern string
}

func New() *Logger {
	return &Logger{
		Logger:           glog.New(),
		LogStdout:        true,
		ErrorStack:       true,
		ErrorLogEnabled:  true,
		ErrorLogPattern:  "error-{Ymd}.log",
		AccessLogEnabled: false,
		AccessLogPattern: "access-{Ymd}.log",
	}
}

func NewWithLog(log *glog.Logger) *Logger {
	return &Logger{
		Logger:           log,
		LogStdout:        true,
		ErrorStack:       true,
		ErrorLogEnabled:  true,
		ErrorLogPattern:  "error-{Ymd}.log",
		AccessLogEnabled: false,
		AccessLogPattern: "access-{Ymd}.log",
	}
}
