package middleware

import (
	"bytes"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"io"
	"lms_try/helper/responseWriter"
	"net/http"
	"os"
	"time"
)

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logrus.SetOutput(os.Stdout)
		startTime := time.Now()

		// capture RequestBody
		requestBody, err := CaptureRequestBody(r.Body)
		if err != nil {
			logrus.Error("failed to get requestBody")
		}

		r.Body = io.NopCloser(bytes.NewReader([]byte(requestBody)))

		// create custom response writer
		customW := responseWriter.NewCustomResponseWriter(w)

		// next to handler
		next.ServeHTTP(customW, r)

		// get responseBody
		responseBody, err := CaptureResponseBody(customW)

		var logFilePath string
		if os.Getenv("LOG_OS") == "linux" {
			logFilePath = "/app/log/lmsapp.log"
		} else {
			logFilePath = "./log/lmsapp.log"
		}

		file, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			logrus.Fatalf("Error opening log file: %v", err)
		}

		logrus.SetOutput(file)
		logrus.WithFields(logrus.Fields{
			"middleware":    true,
			"url":           r.URL.String(),
			"method":        r.Method,
			"request":       requestBody,
			"response_time": time.Since(startTime).Milliseconds(),
			"status_code":   customW.StatusCode,
			"response":      responseBody,
		}).Info("success hit endpoint")
	})
}

// method to get request body
func CaptureRequestBody(body io.Reader) (string, error) {
	var requestMap map[string]any
	err := json.NewDecoder(body).Decode(&requestMap)
	if err != nil {
		return "", err
	}

	requestBody, err := json.Marshal(&requestMap)
	if err != nil {
		return "", err
	}

	return string(requestBody), nil
}

// method to get response body
func CaptureResponseBody(w *responseWriter.CustomResponseWriter) (string, error) {
	var responseBody bytes.Buffer
	if err := json.Indent(&responseBody, w.Buff.Bytes(), "", ""); err != nil {
		return "", err
	}
	return responseBody.String(), nil
}
