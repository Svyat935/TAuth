package utils

import (
	"log"
	"os"
)

// Log Level
const (
	DEBUG = iota
	INFO
	WARNING
	ERROR
)

type Logger struct {
	level  int
	logger *log.Logger
}

func CreateLogger(level int) *Logger {
	return &Logger{
		level:  level,
		logger: log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

func (c *Logger) Debug(message string) {
	if c.level <= DEBUG {
		c.logger.SetPrefix("[DEBUG] ")
		c.logger.Println(message)
	}
}

func (c *Logger) Info(message string) {
	if c.level <= INFO {
		c.logger.SetPrefix("[INFO] ")
		c.logger.Println(message)
	}
}

func (c *Logger) Warn(message string) {
	if c.level <= WARNING {
		c.logger.SetPrefix("[WARNING] ")
		c.logger.Println(message)
	}
}

func (c *Logger) Error(message string) {
	if c.level <= ERROR {
		c.logger.SetPrefix("[ERROR] ")
		c.logger.Println(message)
	}
}
