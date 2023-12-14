package responseWriter

import (
	"bytes"
	"net/http"
)

type CustomResponseWriter struct {
	W          http.ResponseWriter
	StatusCode int
	Buff       *bytes.Buffer
}

func NewCustomResponseWriter(w http.ResponseWriter) *CustomResponseWriter {
	return &CustomResponseWriter{w, 0, bytes.NewBuffer(nil)}
}

func (c *CustomResponseWriter) Header() http.Header {
	return c.W.Header()
}

func (c *CustomResponseWriter) Write(i []byte) (int, error) {
	c.Buff.Write(i)
	return c.W.Write(i)
}

func (c *CustomResponseWriter) WriteHeader(statusCode int) {
	c.StatusCode = statusCode
	c.W.WriteHeader(statusCode)
}
