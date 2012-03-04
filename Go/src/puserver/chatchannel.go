package main

type IChatChannel interface {
	ChatChannelType() int
	
	GetName() string
	GetId() int
	GetOwner() uint64
	
	AddUser(_player *Player) bool
	RemoveUser(_player *Player, _sendCloseChannel bool) bool
	Talk(_fromPlayer *Player, _type int, _text string, _time int) bool
}

type ChatChannel struct {
	Id		int
	Users	ChatUsersMap
	Name	string
}

func NewChatChannel(_id int, _name string) *ChatChannel {
	return &ChatChannel{ Id: _id,
						 Name: _name,
						 Users: make(ChatUsersMap) }
}

func (c *ChatChannel) ChatChannelType() int {
	return 1;
}

func (c *ChatChannel) GetName() string {
	return c.Name
}

func (c *ChatChannel) GetOwner() uint64 {
	return 0
}

func (c *ChatChannel) GetId() int {
	return c.Id
}

func (c *ChatChannel) AddUser(_player *Player) bool {
	_, found := c.Users[_player.GetUID()]
	if !found {
		c.Users[_player.GetUID()] = _player
	}
	
	return !found
}

func (c *ChatChannel) RemoveUser(_player *Player, _sendCloseChannel bool) bool {
	_, found := c.Users[_player.GetUID()]
	if found {
		delete(c.Users, _player.GetUID())
		
		if _sendCloseChannel {
			_player.sendClosePrivateChat(c.Id)
		}
	}
	
	return !found
}

func (c *ChatChannel) Talk(_fromPlayer *Player, _type int, _text string, _time int) (success bool) {
	success = false
	
	// Can't speak to a channel you're not connected to
	if _, found := c.Users[_fromPlayer.GetUID()]; found {
		for _, value := range(c.Users) {
			value.sendToChannel(_fromPlayer, _type, _text, c.Id, _time)
		}
		
		success = true
	}
	
	return
}

func (c *ChatChannel) SendInfo(_type int, _text string, _time int) {
	for _, value := range(c.Users) {
		value.sendToChannel(nil, _type, _text, c.Id, _time)
	}
}