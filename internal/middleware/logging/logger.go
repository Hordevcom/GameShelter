package logging

import (
	"net/http"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

type Logger struct {
	*logrus.Logger
}

type ResponseData struct {
	status int
	size   int
}

type loggingResponceWriter struct {
	http.ResponseWriter
	responceData *ResponseData
}

func NewLogger() *Logger {
	log := logrus.New()

	log.SetOutput(os.Stdout)
	log.SetLevel(logrus.DebugLevel)

	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	return &Logger{Logger: log}
}

func (r *loggingResponceWriter) Write(b []byte) (int, error) {
	size, err := r.ResponseWriter.Write(b)
	r.responceData.size += size
	return size, err
}

func (l *Logger) WithLogging(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		responseData := &ResponseData{
			status: 0,
			size:   0,
		}
		lw := loggingResponceWriter{
			ResponseWriter: w,
			responceData:   responseData,
		}

		h.ServeHTTP(&lw, r)

		duration := time.Since(start)
		l.Logger.Infoln(
			"uri", r.RequestURI,
			"method", r.Method,
			"duration", duration,
			"status", responseData.status,
			"size", responseData.size,
		)
	})
}
