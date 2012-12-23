DIRS=\
	gomysql\
	net/websocket\
	npclib\
	puhelper\
	pulogic\
	pulogic/pokemon\
	pulogic/netmsg\
	puserver/pubattle
	
BUILD_DIRS=\
	puserver\
	orgasm
	
install: 
	go install $(DIRS)
	cp puserver/puserver.exe ../bin/puserver.exe
	
build: 
	go build $(BUILD_DIRS)
	cp puserver/puserver.exe ../bin/puserver.exe
	cp orgasm/orgasm.exe ../bin/orgasm.exe

clean:
	go clean $(DIRS)
	go clean $(BUILD_DIRS)
	
echo-dirs:
	@echo $(DIRS)