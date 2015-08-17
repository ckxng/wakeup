Wakeup
======
This tool is designed to be used by families with small children, or by adults
who need a little more structure in their wakeup routine.  The application
runs in a cross-platform CEF window and displays the current time as well as
the weather for the day.  It is intended to be run 24/7 as a full-screen alarm
clock.  At the time indicated in the configuration file, a wakeup routine will
begin.  This routine is a sequence of songs, each one pertaining to some
pre-defined aspect of the morning routine.  An icon will flash on-screen
indicating what the associated action is (get up, brush teeth, eat breakfast).

Platforms
---------
The primary target for Wakeup is a RaspberryPi running Raspbian ARMv6 with a 
small (3.2inch) screen.  However, to whatever extent it is possible, 
compatibility will be maintained with the then-current Linux x86 and x86_64, 
Windows x86_64, and OS X x86_64 running the most recent stable version of Go.


Getting started on Windows
--------------------------
1. Install Go 32-bit. CEF 64-bit binaries are still experimental and
   were not tested.

2. Install mingw 32-bit and add C:\MinGW\bin to PATH. You can install mingw
   using [mingw-get-setup.exe]
   (http://sourceforge.net/projects/mingw/files/Installer/).
   Select packages to install: "mingw-developer-toolkit",
   "mingw32-base", "msys-base". CEF2go was tested and works fine
   with GCC 4.8.2. You can check gcc version with "gcc --version".

3. Download CEF 3 branch 1750 revision 1590 binaries:
   [cef_binary_3.1750.1590_windows32.7z]
   (https://github.com/CzarekTomczak/cef2go/releases/download/cef3-b1750-r1590/cef_binary_3.1750.1590_windows32.7z)  
   Copy Release/* to cef2go/Release  
   Copy Resources/* to cef2go/Release  

4. Run build.bat (or "build.bat noconsole" to get rid of the console
    window when running the final executable)


Getting started on Linux
------------------------
1. These instructions work fine with Ubuntu 12.04 64-bit. 
   On Ubuntu 13/14 libudev.so.0 is missing and it is required to 
   create a symbolic link to libudev.so.1. For example on 
   Ubuntu 14.04 64-bit run this command: 
  `cd /lib/x86_64-linux-gnu/ && sudo ln -sf libudev.so.1 libudev.so.0`.

2. Install CEF dependencies:  
   `sudo apt-get install build-essential libgtk2.0-dev libgtkglext1-dev`

3. Download CEF 3 branch 1750 revision 1604 binaries:
   [cef_binary_notcmalloc_3.1750.1604_linux64.zip]
   (https://github.com/CzarekTomczak/cef2go/releases/download/cef3-b1750-r1604/cef_binary_notcmalloc_3.1750.1604_linux64.zip)  
   Copy Release/* to cef2go/Release

4. Run "make" command.


Getting started on Mac OS X
---------------------------
1. These instructions work fine with OS X 10.8 Mountain Lion.
   May also work with other versions, but were not tested.

2. Install Go 32-bit. Tested with Go 1.2-386 for OSX 10.8.
   CEF binaries for OSX 64-bit are still experimental, that's
   why we're using 32-bit. Though you can try building with
   CEF 64-bit, download binaries from [cefbuilds.com]
   (http://cefbuilds.com).

3. Install command line tools (make is required) from:  
   https://developer.apple.com/downloads/  
   (In my case command line tools for Mountain Lion from September 2013)

4. Install XCode (gcc that comes with XCode is required). 
   Use the link above. In my case it was XCode 4.6.3 from June 2013.

5. Download CEF 3 branch 1750 revision 1625 binaries for 32-bit:
   [releases/tag/v0.12]
   (https://github.com/CzarekTomczak/cef2go/releases/tag/v0.12)  
   Copy the cef2go.app directory to cef2go/Release.

6. Run "make" command.

Copying
-------
Wakeup is licensed under the 2-CLAUSE BSD LICENSE.  See LICENSE for details.

CEF2go is licensed under the 3-CLAUSE BSD LICENSE.  See cef2go.LICENSE for 
details.

AngularJS is licensed under the MIT license.  See angularjs.LICENSE for details.

Bootstrap is licensed under the MIT license.  See bootstrap.LICENSE for details.

Date.Format is licensed under the MIT license.  See 
http://blog.stevenlevithan.com/code for details.

Html5shiv is licensed under the MIT and GPL2 license.  See html5shiv.LICENSE.md
for details.

jQuery is licensed under the MIT license.  See http://jquery.org/license/ for 
details.

Respond.js is licensed under the MIT license.  See respondjs.LICENSE for 
details.

Cron for Javascript is licensed under the GPL3 license.  See 
http://www.gnu.org/licenses/ for details.