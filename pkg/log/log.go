package log

import (
	"github.com/mattn/go-colorable"
	"github.com/sirupsen/logrus"
)

// InitLog formats the log
func InitLog(level logrus.Level) {
	logrus.SetOutput(colorable.NewColorableStdout())
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:            true,
		PadLevelText:           true,
		FullTimestamp:          true,
		DisableLevelTruncation: true,
	})
	logrus.SetLevel(level)
}
