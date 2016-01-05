package helpers

import (
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

var logger = log.New(os.Stdout, "[Go Web Server]", 0)

func HttpLogger(fn func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		ip_addr := strings.Split(r.RemoteAddr, ":")[0]

		fn(w, r)
		logger.Printf("%s - [%v] \"%s %s\" \"%s%s\" %v 200", ip_addr, start, r.Method, r.URL.Path, r.Host, r.URL, time.Since(start))
	}
}

func NotFoundLogger(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	ip_addr := strings.Split(r.RemoteAddr, ":")[0]

	resp := ShowResponse(404, "", map[string]interface{}{})
	WriteResponse(resp, w, r)

	logger.Printf("%s - [%v] \"%s %s\" \"%s%s\" %v 404", ip_addr, start, r.Method, r.URL.Path, r.Host, r.URL, time.Since(start))
}
