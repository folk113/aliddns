package log

import (
	"log"
	"os"
)

type Logger struct {
	client *log.Logger
}

func NewLogger() *Logger {
	client := log.New(os.Stdout, "[ALIDDNS] ", log.LstdFlags)
	return &Logger{
		client: client,
	}
}

func (l *Logger) Infof(msg string, ctx ...interface{}) {
	l.client.Printf("INFO "+msg+"\n", ctx)
}

func (l *Logger) Info(ctx ...interface{}) {
	l.client.Println(ctx)
}

func (l *Logger) Errorf(msg string, ctx ...interface{}) {
	l.client.Printf("ERROR "+msg+"\n", ctx)
}

func (l *Logger) Error(ctx ...interface{}) {
	l.client.Println("ERROR ", ctx)
}
