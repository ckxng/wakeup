// Copyright (c) 2015 Cameron King. All rights reserved.
// License: BSD 2-clause.
// Website: https://github.com/ckxng/wakeup

package window

import (
    "log"
	"os"
	"cef"
	"wakeup/config"
)

var logger *log.Logger = log.New(os.Stdout, "[window] ", log.Lshortfile)

func Go(cfg *config.Config, c chan int) {
	logger.Println("executeCEFSubprocess")
    executeCEFSubprocess()
	
	logger.Println("initializeCEFSettings")
    initializeCEFSettings(cfg)
    
	logger.Println("createWindow")
	createWindow(cfg)
	
	logger.Println("cef.RunMessageLoop")
    cef.RunMessageLoop()
	
	logger.Println("cef.Shutdown")
    cef.Shutdown()
    c <- 0
}

func initializeCEFSettings(cfg *config.Config) {
	settings := cef.Settings{}
	settings.CachePath = cfg.CachePath
	settings.LogSeverity = cef.LOGSEVERITY_DEFAULT
	cef.Initialize(settings)
}