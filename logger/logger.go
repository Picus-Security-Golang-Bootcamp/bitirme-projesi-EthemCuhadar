package logger

import (
	"fmt"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// InitLogger creates a new "Core" — which is a minimal, fast logger interface— based on the configuration
// previously set. Then, "zap.New()" constructs a new Logger from the provided core and Options.
// The only option set is "zap.AddCaller()", which configures the Logger to annotate each message with
// the filename, line number, and function name.
func InitLogger() {
	createDirectoryIfNotExist()
	writerSync := getLogWriter()
	encoder := getEncoder()

	core := zapcore.NewCore(encoder, writerSync, zapcore.DebugLevel)
	logg := zap.New(core, zap.AddCaller())
	zap.ReplaceGlobals(logg)
}

// createDirectoryIfNotExist gets where the application running. Next, it creates the file path for our
// log folder using "fmt.Sprintf()" which returns a string formatted. Finally, it checks if the folder
// does exist and create it otherwise.
func createDirectoryIfNotExist() {
	path, _ := os.Getwd()
	if _, err := os.Stat(fmt.Sprintf("%s/logs", path)); os.IsNotExist(err) {
		_ = os.Mkdir("logs", os.ModePerm)
	}
}

// getLogWriter gets the file path to the environment path that the app is running and then called "os.OpenFile()".
// Since the flags "O_APPEND" and "O_CREATE" are used, if the file exists, it's only going to write after the
// content of the file, otherwise it will create a new one.
func getLogWriter() zapcore.WriteSyncer {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	file, err := os.OpenFile(path+"/logs/logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	return zapcore.AddSync(file)
}

// The first function returns an opinionated EncoderConfig for production environments, which have several fields already
// configured. We are going to change only two, "EncodeTime" to UTC and EncodeLevel to "CapitalColorLevelEncoder".
// The later prints the log level in all capital and colored based on the level.
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.TimeEncoder(func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.UTC().Format("2006-01-02T15:04:05Z0700"))
	})
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

// Close will terminates syncronization.
func Close() {
	defer zap.L().Sync()
}
