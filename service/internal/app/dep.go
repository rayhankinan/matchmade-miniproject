package app

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
)

type Dep struct {
	Logger *logrus.Logger
}

func NewDep() (dep *Dep) {
	// Create a new Dep struct and set the Logger field to a new logrus.Logger instance
	l := logrus.New()
	l.SetReportCaller(true)
	l.SetFormatter(&logrus.TextFormatter{
		CallerPrettyfier: func(frame *runtime.Frame) (funcName string, fileName string) {
			_, after, _ := strings.Cut(frame.File, "service/")
			fileName = fmt.Sprintf("%s:%d", after, frame.Line)

			s := strings.Split(frame.Function, ".")
			funcName = s[len(s)-1]

			return
		},
	})
	l.SetLevel(logrus.DebugLevel)

	// Set the Logger field to the new logrus.Logger instance
	dep = &Dep{
		Logger: l,
	}

	return
}
