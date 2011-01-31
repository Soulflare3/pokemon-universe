# Copyright 2009 The Go Authors.  All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

include $(GOROOT)/src/Make.inc

TARG=db
CGOFILES=\
	mysql.go\
	
GOFILES=\
	database.go\

CGO_LDFLAGS=-lmysqlclient
CLEANFILES+=db

include $(GOROOT)/src/Make.pkg

%: install %.go
	$(GC) $*.go
	$(LD) -o $@ $*.$O

