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
    "encoding/json"
    "github.com/gorilla/mux"
    "github.com/codegangsta/negroni"
	"wakeup/config"
)

var logger *log.Logger = log.New(os.Stdout, "[server] ", log.Lshortfile)

// a specialized version of http.HandlerFunc that returns (interface{}, error)
// and is intended to be wrapped by JSONDecorator
type JSONHandlerFunc func(http.ResponseWriter, *http.Request) (interface{}, error)

func Go(cfg *config.Config, c chan int) {
	logger.Println("router = mux.NewRouter")
    router := mux.NewRouter().StrictSlash(false)
	
	logger.Println("router: adding controllers")
    router.HandleFunc("/", handler)
    router.HandleFunc("/json", JSONDecorator(jsonHandler))
    // add additional controller routes
	
	logger.Println("router: adding static file server")
    router.PathPrefix("/").Handler(http.FileServer(http.Dir("./assets/")))

	logger.Println("stack = negroni.New")
    stack := negroni.New()
	
	logger.Println("stack: adding middleware")
    // add additional middleware
	
	logger.Println("stack.UseHandler router")
    stack.UseHandler(router)

    http.Handle("/", stack)
	
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
	fmt.Fprintf(w, "Try <a href=/json>/json</a><br>")
}

func jsonHandler(http.ResponseWriter, *http.Request) (interface{}, error) {
	return "ok", nil
}

// wraps JSONHanderFunc functions and provide JSON Content-type, marshaling,
// and HTTP error code functionality
func JSONDecorator(handler JSONHandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        data, err := handler(w,r)

        payload := struct{
            Data    interface{}         `json:"data"`
            Err     interface{}         `json:"err"`
        } {
            Data: data,
            Err: err,
        }
        
        if err != nil {
            payload.Err = err.Error()
        }
        
        w.Header().Set("Content-Type", "application/json")
        if bPayload, merr := json.MarshalIndent(payload, "", "  "); merr == nil {
            if err != nil {
                w.WriteHeader(500)
            }
            fmt.Fprint(w, string(bPayload))
        } else {
            w.WriteHeader(500)
            fmt.Fprintf(w, "{\"data\":\"\",\"err\":\"Marshal error: %s\"}", err)
        }
    }
}