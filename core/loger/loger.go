package loger

import (
	"github.com/Sirupsen/logrus"
)

var Log = logrus.New()

func init() {

	// in future add logs file

	Log.Formatter = &logrus.JSONFormatter{}
}