package main

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/tarm/serial"
)

const dir = "./"

var (
	tty = "COM9"
	// port         = ":8080"
	baud    uint = 115200
	monitor      = false
)

func main() {
	fs := http.FileServer(http.Dir(dir))
	log.Print("Serving " + dir + " on http://localhost:8080")
	http.ListenAndServe(":8080", http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		resp.Header().Add("Cache-Control", "no-cache")
		if strings.HasSuffix(req.URL.Path, ".wasm") {
			resp.Header().Set("content-type", "application/wasm")
		}
		fs.ServeHTTP(resp, req)
	}))

	fp, err := serial.OpenPort(&serial.Config{
		Name:        tty,
		Baud:        int(baud),
		ReadTimeout: time.Second,
	})
}
