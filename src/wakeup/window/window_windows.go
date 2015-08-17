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
	
    hwnd, _ := createWindowSize(cfg.Title, wndproc, cfg.WindowX, cfg.WindowY, cfg.WindowWidth, cfg.WindowHeight)
	
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

func createWindowSize(title string, wndproc uintptr, x int32, y int32, width int32, height int32) (hwnd syscall.Handle, err error) {
	logger.Println("wingui.GetModuleHandle")
	var mh syscall.Handle
	if mh, err = wingui.GetModuleHandle(nil); err != nil {
		return hwnd, err
	}
	
	logger.Println("wingui.LoadIcon IDI_APPLICATION")
	var myicon syscall.Handle
    if myicon, err = wingui.LoadIcon(0, wingui.IDI_APPLICATION); err != nil {
        return hwnd, err
    }
	
	logger.Println("wingui.LoadCursor IDC_ARROW")
	var mycursor syscall.Handle
    if mycursor, err = wingui.LoadCursor(0, wingui.IDC_ARROW); err != nil {
		return hwnd, err
	}
	
	logger.Println("wingui.RegisterClassEx myWindowClass")
	wcname := syscall.StringToUTF16Ptr("myWindowClass")
    wc := wingui.Wndclassex {
    	WndProc: 	wndproc,
    	Instance:	mh,
    	Icon:		myicon,
    	Cursor:		mycursor,
    	Background:	wingui.COLOR_BTNFACE + 1,
    	MenuName:	nil,
    	ClassName:	wcname,
    	IconSm:		myicon,
	}
    wc.Size = uint32(unsafe.Sizeof(wc))
	if _, err = wingui.RegisterClassEx(&wc); err != nil {
		return hwnd, err
	}
	
	logger.Println("wingui.CreateWindowEx")
	if hwnd, err = wingui.CreateWindowEx(
        0,
        wcname,
        syscall.StringToUTF16Ptr(title),
        wingui.WS_OVERLAPPEDWINDOW,
        x, y, width, height,
        0, 0, mh, 0); err != nil {
			return hwnd, err
	}
	
	logger.Println("wingui.ShowWindow")
	wingui.ShowWindow(hwnd, wingui.SW_SHOWDEFAULT)
	
	logger.Println("wingui.UpdateWindow")
	if err = wingui.UpdateWindow(hwnd); err != nil {
		return hwnd, err
	}
	
	return hwnd, nil
}