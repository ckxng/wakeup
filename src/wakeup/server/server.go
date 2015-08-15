// Copyright (c) 2015 Cameron King. All rights reserved.
// License: BSD 2-clause.
// Website: https://github.com/ckxng/wakeup

package server

import (
    "log"
	"os"
	"runtime"
	"fmt"
    "time"
	"net/http"
	"wakeup/config"
)

var logger *log.Logger = log.New(os.Stdout, "[server] ", log.Lshortfile)

func Go(cfg *config.Config, c chan int) {
	logger.Println("http.HandleFunc handler")
    http.HandleFunc("/", handler)
	
    listen_at := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
    logger.Printf("http.ListenAndServe %s\n", listen_at)
    if err := http.ListenAndServe(listen_at, nil); err != nil {
		logger.Printf("error %s\n", err)
		c <- 1
		return
	}
	
	logger.Println("shutdown")
	c <- 0
}

func handler(w http.ResponseWriter, req *http.Request) {
	logger.Println("serve")
    w.Header().Set("Content-Type", "text/html; charset=utf-8")
    fmt.Fprintf(w, "This is Go server talking.<br>")
    fmt.Fprintf(w, "Time on the server: %v<br>",
            time.Now().Format("Jan 2, 2006 at 3:04pm (MST)"))
    fmt.Fprintf(w, "Go version: %v<br>", runtime.Version())
    fmt.Fprintf(w, "<br>")
    if req.URL.Path == "/" {
        fmt.Fprintf(w, "Try <a href=/test.go>/test.go</a><br>")
    } else if req.URL.Path == "/test.go" {
        fmt.Fprintf(w, "<b>You accessed /test.go</b><br>")
    }
}