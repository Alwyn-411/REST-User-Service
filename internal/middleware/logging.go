package middleware

import (
	"log"
	"net/http"
	"time"
)

func LogRequest(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        startTime := time.Now()
        
        // Log request details
        log.Printf("Started %s %s", r.Method, r.URL.Path)
        
        next.ServeHTTP(w, r)
        
        // Log completion time
        log.Printf("Completed %s %s in %v", r.Method, r.URL.Path, time.Since(startTime))
    }
}
