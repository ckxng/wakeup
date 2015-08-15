// Copyright (c) 2015 Cameron King. All rights reserved.
// License: BSD 2-clause.
// Website: https://github.com/ckxng/wakeup

package window

import (
	"cef"
	"syscall"
	"unsafe"
    "wingui"
    "time"
	"wakeup/config"
)

func executeCEFSubprocess() {
	logger.Println("GetModuleHandle")
	hInstance, e := wingui.GetModuleHandle(nil)
    if e != nil { wingui.AbortErrNo("GetModuleHandle", e) }
    
	logger.Println("cef.ExecuteProcess")
    cef.ExecuteProcess(unsafe.Pointer(hInstance))
}

func createWindow(cfg *config.Config) {
    wndproc := syscall.NewCallback(wndProc)
    logger.Println("wingui.CreateWindow")
    hwnd := wingui.CreateWindow(cfg.Title, wndproc)
	
    browserSettings := cef.BrowserSettings{}
	url := cfg.URL()
	logger.Printf("cef.CreateBrowser %s\n", url)
    cef.CreateBrowser(unsafe.Pointer(hwnd), browserSettings, url)
	
	// It should be enough to call WindowResized after 10ms,
    // though to be sure let's extend it to 100ms.
    time.AfterFunc(time.Millisecond * 100, func(){
		logger.Println("cef.WindowResized")
        cef.WindowResized(unsafe.Pointer(hwnd))
    })
}

func wndProc(hwnd syscall.Handle, msg uint32, wparam, lparam uintptr) (rc uintptr) {
    switch msg {
    case wingui.WM_CREATE:
        rc = wingui.DefWindowProc(hwnd, msg, wparam, lparam)
    case wingui.WM_SIZE:
        cef.WindowResized(unsafe.Pointer(hwnd))
    case wingui.WM_CLOSE:
        wingui.DestroyWindow(hwnd)
    case wingui.WM_DESTROY:
        cef.QuitMessageLoop()
    default:
        rc = wingui.DefWindowProc(hwnd, msg, wparam, lparam)
    }
    return
}