package middleware

import (
	"github.com/sirupsen/logrus"
)

type IMiddleware interface{}

type Middleware struct {
	Logger *logrus.Logger
}

func NewMiddleware(logger *logrus.Logger) (m IMiddleware) {
	m = Middleware{Logger: logger}
	return
}
