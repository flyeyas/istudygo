package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"strings"
)

func main() {
	fmt.Println("Starting http server...")
	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/healthz", healthz)
	err := http.ListenAndServe(":80", mux)
	if err != nil {
		fmt.Printf("Fatal: %s", err)
	}

}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("entering root handler")

	os.Setenv("VERSION", "0.0.1")

	version := os.Getenv("VERSION")

	fmt.Printf("VERSION %s\n", version)

	w.Header().Add("VERSION", version)

	for k, v := range r.Header {
		w.Header().Add(k, v[0])
	}

	clientIP := ClientIP(r)
	httpCode := http.StatusOK
	w.WriteHeader(httpCode)
	fmt.Printf("clentIP: %s, responseHttpCode: %d\n", clientIP, httpCode)

}

func healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	io.WriteString(w, "200")
}

func ClientIP(r *http.Request) string {
	xForwardedFor := r.Header.Get("X-Forwarded-For")
	ip := strings.TrimSpace(strings.Split(xForwardedFor, ",")[0])
	if ip != "" {
		return ip
	}

	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" {
		return ip
	}

	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		return ip
	}
	return ""
}
