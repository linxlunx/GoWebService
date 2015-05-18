package main

import (
	"os"
	"net/http"
	"log"
	"time"
	"strings"
)

var logger = log.New(os.Stdout, "[Go Web Server]", 0)

func httpLogger(fn func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
    return func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        ip_addr := strings.Split(r.RemoteAddr,":")[0]

        fn(w, r)
        logger.Printf("%s - [%v] \"%s %s\" \"%s%s\" %v", ip_addr, start, r.Method, r.URL.Path, r.Host, r.URL, time.Since(start))
    }
}