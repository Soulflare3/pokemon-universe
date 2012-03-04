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
	
build: 
	go build $(DIRS)

clean:
	go clean $(DIRS)
	
echo-dirs:
	@echo $(DIRS)