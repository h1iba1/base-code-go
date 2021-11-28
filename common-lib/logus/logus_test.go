package logus

import (
	"bytes"
	"github.com/sirupsen/logrus"
	"io"
	"log"
	"os"
	"testing"
)

func TestLogLevel(t *testing.T) {
	logrus.SetLevel(logrus.TraceLevel)

	logrus.Trace("trace msg")
	logrus.Debug("debug msg")
	logrus.Info("info msg")
	logrus.Warn("warn msg")
	logrus.Error("error msg")
	logrus.Fatal("fatal msg")
	logrus.Panic("panic msg")
}

func TestSetReportCaller(t *testing.T) {
	logrus.SetReportCaller(true)

	logrus.Info("info msg")
}

func TestWithFields(t *testing.T) {
	logrus.WithFields(logrus.Fields{
		"name": "dj",
		"age": 18,
	}).Info("info msg")
}

func TestRequestLogger(t *testing.T) {
	requestLogger := logrus.WithFields(logrus.Fields{
		"user_id": 10010,
		"ip":      "192.168.32.15",
	})

	requestLogger.Info("info msg")
	requestLogger.Error("error msg")
}

func TestSetOutput(t *testing.T) {
	writer1 := &bytes.Buffer{}
	writer2 := os.Stdout
	writer3, err := os.OpenFile("log.txt", os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		log.Fatalf("create file log.txt failed: %v", err)
	}

	//同时将日志写到bytes.Buffer、标准输出和文件中：
	logrus.SetOutput(io.MultiWriter(writer1, writer2, writer3))
	logrus.Info("info msg")
}

func TestJSONFormatter(t *testing.T) {
	logrus.SetLevel(logrus.TraceLevel)
	logrus.SetFormatter(&logrus.JSONFormatter{})

	logrus.Trace("trace msg")
	logrus.Debug("debug msg")
	logrus.Info("info msg")
	logrus.Warn("warn msg")
	logrus.Error("error msg")
	logrus.Fatal("fatal msg")
	logrus.Panic("panic msg")

}
