How to setup a basic text to file logger

# Introduction #

Following code shows how to create a basic text to file logger using the Logger package

```
package main

import (
    "fmt"
    "log"
    pulog "logger" // PU Logger package
)

// Create a global variable to we can use it anywhere in our application
var g_logger *log.Logger 

func main() {
    fmt.Println("Demo on how to setup a text to file logger")

    // Setup logger
    myLog, err := pulog.NewLogger("mylogfile.txt", pulog.L_CONSOLE|pulog.L_FILE) // Create PU logger constructor
    if err != nil { // Will return an error if it fails to create/open the file
   	return
    }
    g_logger = log.New(myLog, "", log.Ltime) // Create a new log.Logger using our own log io.Writer
    g_logger.Println("Writing to log here!")    
}
```

### Options ###
It's possible to regulate the output of the PU logger by defining different flags in the NewLogger function:

```
func NewLogger(_filename string, _flag int) (log *Logger, err os.Error)
```

The following options are available:
```
const (
    L_FILE	int = 1 // Write to file
    L_CONSOLE	int = 2 // Print to terminal/console
)
```

How to combine them:
```
flags := logger.L_CONSOLE|logger.L_FILE
```


---


To change the date/time format you can use one of the predefined constants in the Go [log](http://golang.org/pkg/log/#Constants) package:
```
const (
    // Bits or'ed together to control what's printed. There is no control over the
    // order they appear (the order listed here) or the format they present (as
    // described in the comments).  A colon appears after these items:
    //	2009/0123 01:23:23.123123 /a/b/c/d.go:23: message
    Ldate         = 1 << iota // the date: 2009/0123
    Ltime                     // the time: 01:23:23
    Lmicroseconds             // microsecond resolution: 01:23:23.123123.  assumes Ltime.
    Llongfile                 // full file name and line number: /a/b/c/d.go:23
    Lshortfile                // final file name element and line number: d.go:23. overrides Llongfile
)
```