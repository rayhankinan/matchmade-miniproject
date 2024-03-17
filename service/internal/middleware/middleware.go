package middleware

import (
	"github.com/sirupsen/logrus"
)

type IMiddleware interface{}

type Middleware struct {
	Logger *logrus.Logger
}

func NewMiddleware(logger *logrus.Logger) IMiddleware {
	return Middleware{Logger: logger}
}
