package exception

import (
	"github.com/sirupsen/logrus"
)

type AppError struct {
	Code    int    // HTTP статус код
	Message string // Описание ошибки
}

func Cast(code int, message string) *AppError {
	logrus.Error(message)
	return &AppError{code, message}
}

func (e *AppError) Error() string {
	return e.Message
}
