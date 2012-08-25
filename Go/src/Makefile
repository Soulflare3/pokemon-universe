DIRS=\
	goconf\
	gomysql\
	net/websocket\
	network\
	npclib\
	puhelper\
	pulogic\
	pulogic/pokemon\
	pulogic/netmsg\
	putools/math\
	putools/pos\
	putools/log\
	puserver/pubattle\
	puserver\
	
install: 
	go install $(DIRS)
	cp puserver/puserver.exe ../bin/puserver.exe
	
build: 
	go build $(DIRS)
	cp puserver/puserver.exe ../bin/puserver.exe

clean:
	go clean $(DIRS)
	
echo-dirs:
	@echo $(DIRS)