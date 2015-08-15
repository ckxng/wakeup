// Copyright (c) 2015 Cameron King. All rights reserved.
// License: BSD 2-clause.
// Website: https://github.com/ckxng/wakeup

package main

import (
    "os"
    "log"
	"wakeup/window"
	"wakeup/config"
	"wakeup/server"
)

var logger *log.Logger = log.New(os.Stdout, "[main] ", log.Lshortfile)
var main_is_running = false

func main() {
	cWindow := make(chan int)
	cServer := make(chan int)
	
	logger.Println("config.NewConfig")
	cfg := config.NewConfig()
	cfg.Title = "Wakeup"
	
	// cef2go calls itself several times to start subprocesses of the 
	// executable.  Since we are running a server, only the first executable
	// can bind the port.
	if len(os.Args) > 1 {
		logger.Println("os.Args are present, not running server")
		cfg.EnableServer = false
	}
	
	if(cfg.EnableServer) {
		logger.Println("go server.Go")
		go server.Go(cfg, cServer)
	}
	
	if(cfg.EnableWindow) {
		logger.Println("go window.Go")
		go window.Go(cfg, cWindow)
	}
	
	for {
		select {
			case iExitWindow := <- cWindow:
				logger.Printf("Window Exit: %d\n", iExitWindow)
				os.Exit(iExitWindow)
			case iExitServer := <- cServer:
				logger.Printf("Server Exit: %d\n", iExitServer)
				os.Exit(iExitServer)
		}
	}
}