/*Pokemon Universe MMORPG
Copyright (C) 2010 the Pokemon Universe Authors

This program is free software; you can redistribute it and/or
modify it under the terms of the GNU General Public License
as published by the Free Software Foundation; either version 2
of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program; if not, write to the Free Software
Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA  02110-1301, USA.*/

package logger

import (
	"fmt"
	"os"
)

// A Logger represents an object which inherits from io.Writer
// it's used in combination with log.Logger to write text to a file
type Logger struct {
	filename string
	file *os.File
}

// Create constructor for io.Writer
// No need to close the log file because it will last untill the application exists
func NewLogger(_filename string) (log *Logger, err os.Error) {
	log = &Logger { filename : _filename }
	log.file, err = os.Open(log.filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}
	
	return log, nil
}

// Use the io.Writer interface
func (l *Logger) Write(p []byte) (n int, err os.Error) {
	// Print string to terminal before writing to file	
	fmt.Printf("%v", string(p))
	
	// Write to file
	n, err2 := l.file.Write(p)
	if err2 != nil {
		return n, err
	}
	
	return len(p), nil
}
