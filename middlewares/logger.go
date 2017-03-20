package middlewares

import (
	"log"
	"net/http"
	"time"
)

func Logger(rw http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	start := time.Now()
	next(rw, req)

	log.Printf(
		"%s\t%s\t%s",
		req.Method,
		req.RequestURI,
		time.Since(start),
	)
}
