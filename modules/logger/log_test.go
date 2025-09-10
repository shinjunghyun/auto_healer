package log

import (
	"sync"
	"testing"
)

func printLogs() {
	Trace().Msgf("this is a trace message")
	Debug().Msgf("this is a debug message")
	Info().Msgf("this is an info message")
	Warn().Msgf("this is a warn message")
	Error().Msgf("this is an error message")
}

func printManyLogs() {
	wg := &sync.WaitGroup{}
	const routineCount = 1000
	wg.Add(routineCount)
	for i := 0; i < routineCount; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				printLogs()
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func TestConsoleLog(t *testing.T) {
	// print logs
	printLogs()
}

func TestConsoleLogWithFile(t *testing.T) {
	// init logger
	Init(LoggerOptions{
		Level: TraceLevel,
		Style: ConsoleStyle,
		FileOption: &FileOptions{
			FileName:    "test-console-log-with-file.log",
			MaxSizeMB:   3,
			MaxBackups:  10,
			MaxAgeDays:  14,
			Compression: true,
		},
	})

	// print logs
	printLogs()
}

func TestJsonLog(t *testing.T) {
	// init logger
	Init(LoggerOptions{
		Level: TraceLevel,
		Style: JsonStyle,
	})

	// print logs
	printLogs()
}

func TestJsonLogWithFile(t *testing.T) {
	// init logger
	Init(LoggerOptions{
		Level: TraceLevel,
		Style: JsonStyle,
		FileOption: &FileOptions{
			FileName:    "test-json-log-with-file.log",
			MaxSizeMB:   3,
			MaxBackups:  10,
			MaxAgeDays:  14,
			Compression: true,
		},
	})

	// print logs
	printLogs()
}

func TestManyConsoleLogWithFile(t *testing.T) {
	// init logger
	Init(LoggerOptions{
		Level: TraceLevel,
		Style: ConsoleStyle,
		FileOption: &FileOptions{
			FileName:    "test-many-console-log-with-file.log",
			MaxSizeMB:   3,
			MaxBackups:  10,
			MaxAgeDays:  14,
			Compression: true,
		},
	})

	// print many logs
	printManyLogs()
}
