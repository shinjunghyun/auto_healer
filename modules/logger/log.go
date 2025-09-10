package log

import (
	"context"
	"fmt"
	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"sync"
	"time"
)

type LoggerStyle string

const (
	consoleTimeFormatMilli = "2006-01-02T15:04:05.000Z07:00"

	JsonStyle    LoggerStyle = "JSON"
	ConsoleStyle LoggerStyle = "CONSOLE"
)

type FileOptions struct {
	FileName    string
	MaxSizeMB   int
	MaxBackups  int
	MaxAgeDays  int
	Compression bool
}

type LoggerOptions struct {
	Level      LoggerLevel
	Style      LoggerStyle
	FileOption *FileOptions
}

// GCP severity 기준 정의
const (
	DisableLevel   = iota
	EmergencyLevel //Emergency 일 경우 zero log 의 Panic Level 로  표시
	AlertLevel     //Alert 일 경우 zero log 의 Panic Level 로 표시
	CriticalLevel  //Critical 일 경우 zero log 의 Fatal Level 로 표시
	ErrorLevel
	WarningLevel
	NoticeLevel //Notice 일 경우 zero log 의 WARING Level 로 표시
	InfoLevel
	DebugLevel
	DefaultLevel //Default 일 경우 zero log 의 Debug Level 로 표시
	TraceLevel
)

type LoggerLevel int

type logger struct {
	logger   *zerolog.Logger
	logLevel LoggerLevel
	mu       *sync.Mutex
}

var log *logger

var LevelToString = [...]string{
	"DISABLE",
	"EMERGENCY",
	"ALERT",
	"CRITICAL",
	"ERROR",
	"WARNING",
	"NOTICE",
	"INFO",
	"DEBUG",
	"DEFAULT",
	"TRACE",
}

func (l LoggerLevel) Name() string { return LevelToString[l] }

func GetLogLevelFromString(name string) LoggerLevel {
	level := DisableLevel
	for i, v := range LevelToString {
		if v == name {
			level = i
			break
		}
	}
	return LoggerLevel(level)
}

func checkDefaultInit() {
	if log == nil {
		Init(LoggerOptions{
			Level: TraceLevel,
			Style: ConsoleStyle,
		})
	}
}

func Init(options LoggerOptions) {
	level := zerolog.InfoLevel
	switch options.Level {
	case DisableLevel:
		level = zerolog.Disabled
	case EmergencyLevel:
		fallthrough
	case AlertLevel:
		level = zerolog.PanicLevel
	case CriticalLevel:
		level = zerolog.FatalLevel
	case ErrorLevel:
		level = zerolog.ErrorLevel
	case WarningLevel:
		fallthrough
	case NoticeLevel:
		level = zerolog.WarnLevel
	case InfoLevel:
		level = zerolog.InfoLevel
	case DebugLevel:
		fallthrough
	case DefaultLevel:
		level = zerolog.DebugLevel
	case TraceLevel:
		level = zerolog.TraceLevel
	default:
	}

	//GCP Level
	zerolog.LevelFieldName = "severity"

	zerolog.LevelPanicValue = "EMERGENCY"
	//"ALERT"
	zerolog.LevelFatalValue = "CRITICAL"
	zerolog.LevelErrorValue = "ERROR"
	zerolog.LevelWarnValue = "WARNING"
	//"NOTICE"
	zerolog.LevelInfoValue = "INFO"
	zerolog.LevelDebugValue = "DEBUG"
	//"DEFAULT"

	zerolog.LevelTraceValue = "TRACE"

	zerolog.SetGlobalLevel(level)

	zerolog.TimeFieldFormat = time.RFC3339Nano

	var writer io.Writer = os.Stdout

	if options.FileOption != nil {
		// 기본 값 셋팅
		if options.FileOption.FileName == "" {
			if runtime.GOOS == "windows" {
				options.FileOption.FileName = "unknown_filename"
			} else {
				options.FileOption.FileName = fmt.Sprintf("unknown_filename_%d", time.Now().UnixNano())
			}
		}
		if options.FileOption.MaxSizeMB == 0 {
			options.FileOption.MaxSizeMB = 10
		}
		if options.FileOption.MaxBackups == 0 {
			options.FileOption.MaxBackups = 10
		}
		if options.FileOption.MaxAgeDays == 0 {
			options.FileOption.MaxAgeDays = 14
		}

		logFilePath := fmt.Sprintf("/static/logs/%s.log", options.FileOption.FileName)
		if runtime.GOOS == "windows" { // windows 환경에서는 상대경로로 로그파일을 생성
			logFilePath = fmt.Sprintf("logs/%s.log", options.FileOption.FileName)
		}

		fileLogger := &lumberjack.Logger{
			Filename:   logFilePath,
			MaxSize:    options.FileOption.MaxSizeMB, // megabytes
			MaxBackups: options.FileOption.MaxBackups,
			MaxAge:     options.FileOption.MaxAgeDays, // days
			LocalTime:  true,
			Compress:   options.FileOption.Compression,
		}

		if err := os.MkdirAll(filepath.Dir(logFilePath), os.ModePerm); err != nil {
			fmt.Printf("failed to create log directory: %s\n", err.Error())
		} else if options.Style == JsonStyle {
			writer = zerolog.MultiLevelWriter(writer, fileLogger)
		} else {
			writer = zerolog.MultiLevelWriter(zerolog.ConsoleWriter{Out: writer, TimeFormat: consoleTimeFormatMilli}, zerolog.ConsoleWriter{Out: fileLogger, TimeFormat: consoleTimeFormatMilli, NoColor: true})
		}
	} else {
		if options.Style == JsonStyle {
			writer = zerolog.MultiLevelWriter(writer)
		} else {
			writer = zerolog.MultiLevelWriter(zerolog.ConsoleWriter{Out: writer, TimeFormat: consoleTimeFormatMilli})
		}
	}

	zerlogLogger := zerolog.New(writer).With().Timestamp().Caller().Logger()

	log = &logger{
		logger:   &zerlogLogger,
		logLevel: options.Level,
		mu:       &sync.Mutex{},
	}
}

// Output duplicates the global log and sets w as its output.
func Output(w io.Writer) zerolog.Logger {
	checkDefaultInit()
	log.mu.Lock()
	defer log.mu.Unlock()
	return log.logger.Output(w)
}

// With creates a child log with the field added to its context.
func With() zerolog.Context {
	checkDefaultInit()
	log.mu.Lock()
	defer log.mu.Unlock()
	return log.logger.With()
}

// Level creates a child log with the minimum accepted level set to level.
func Level(level zerolog.Level) zerolog.Logger {
	checkDefaultInit()
	log.mu.Lock()
	defer log.mu.Unlock()
	return log.logger.Level(level)
}

// Sample returns a log with the s sampler.
func Sample(s zerolog.Sampler) zerolog.Logger {
	checkDefaultInit()
	log.mu.Lock()
	defer log.mu.Unlock()
	return log.logger.Sample(s)
}

// Hook returns a log with the h Hook.
func Hook(h zerolog.Hook) zerolog.Logger {
	checkDefaultInit()
	log.mu.Lock()
	defer log.mu.Unlock()
	return log.logger.Hook(h)
}

// Err starts a new message with perror level with err as a field if not nil or
// with info level if err is nil.
//
// You must call Msg on the returned event in order to send the event.
func Err(err error) *zerolog.Event {
	checkDefaultInit()
	log.mu.Lock()
	defer log.mu.Unlock()
	return log.logger.Err(err)
}

// Trace starts a new message with trace level.
//
// You must call Msg on the returned event in order to send the event.
func Trace() *zerolog.Event {
	checkDefaultInit()
	log.mu.Lock()
	defer log.mu.Unlock()
	return log.logger.Trace()
}

// Debug starts a new message with debug level.
//
// You must call Msg on the returned event in order to send the event.
func Debug() *zerolog.Event {
	checkDefaultInit()
	log.mu.Lock()
	defer log.mu.Unlock()
	return log.logger.Debug()
}

// Info starts a new message with info level.
//
// You must call Msg on the returned event in order to send the event.
func Info() *zerolog.Event {
	checkDefaultInit()
	log.mu.Lock()
	defer log.mu.Unlock()
	return log.logger.Info().Ctx(context.Background())
}

// Warn starts a new message with warn level.
//
// You must call Msg on the returned event in order to send the event.
func Warn() *zerolog.Event {
	checkDefaultInit()
	log.mu.Lock()
	defer log.mu.Unlock()
	return log.logger.Warn()
}

// Error starts a new message with perror level.
//
// You must call Msg on the returned event in order to send the event.
func Error() *zerolog.Event {
	checkDefaultInit()
	log.mu.Lock()
	defer log.mu.Unlock()
	return log.logger.Error()
}

// Fatal starts a new message with fatal level. The os.Exit(1) function
// is called by the Msg method.
//
// You must call Msg on the returned event in order to send the event.
func Fatal() *zerolog.Event {
	checkDefaultInit()
	log.mu.Lock()
	defer log.mu.Unlock()
	return log.logger.Fatal()
}

// Panic starts a new message with panic level. The message is also sent
// to the panic function.
//
// You must call Msg on the returned event in order to send the event.
func Panic() *zerolog.Event {
	checkDefaultInit()
	log.mu.Lock()
	defer log.mu.Unlock()
	return log.logger.Panic()
}

// Print sends a log event using debug level and no extra field.
// Arguments are handled in the manner of fmt.Print.
func Print(v ...interface{}) {
	checkDefaultInit()
	log.mu.Lock()
	defer log.mu.Unlock()
	log.logger.Debug().CallerSkipFrame(1).Msg(fmt.Sprint(v...))
}

// Printf sends a log event using debug level and no extra field.
// Arguments are handled in the manner of fmt.Printf.
func Printf(format string, v ...interface{}) {
	checkDefaultInit()
	log.mu.Lock()
	defer log.mu.Unlock()
	log.logger.Debug().CallerSkipFrame(1).Msgf(format, v...)
}

// Emergency starts a new message with panic level. The message is also sent
// to the panic function.
//
// You must call Msg on the returned event in order to send the event.
func Emergency() *zerolog.Event {
	checkDefaultInit()
	log.mu.Lock()
	defer log.mu.Unlock()
	return log.logger.Panic()
}

// Critical starts a new message with fatal level. The os.Exit(1) function
// is called by the Msg method.
//
// You must call Msg on the returned event in order to send the event.
func Critical() *zerolog.Event {
	checkDefaultInit()
	log.mu.Lock()
	defer log.mu.Unlock()
	return log.logger.Fatal()
}

func Alert() *zerolog.Event {
	checkDefaultInit()
	log.mu.Lock()
	defer log.mu.Unlock()
	if log.logLevel >= AlertLevel {
		return log.logger.Log().Str("severity", "ALERT")
	} else {
		return log.logger.Panic()
	}
}

func Notice() *zerolog.Event {
	checkDefaultInit()
	log.mu.Lock()
	defer log.mu.Unlock()
	if log.logLevel >= NoticeLevel {
		return log.logger.Log().Str("severity", "NOTICE")
	} else {
		return log.logger.Info()
	}
}

func Default() *zerolog.Event {
	checkDefaultInit()
	log.mu.Lock()
	defer log.mu.Unlock()
	if log.logLevel >= DefaultLevel {
		return log.logger.Log().Str("severity", "DEFAULT")
	} else {
		return log.logger.Trace()
	}
}
