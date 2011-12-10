package main

import (
	"fmt"
)

var dummyPlayer *Player
var client	*POClient

func POTestClientDoIt() {
	LoadDummyPlayer()
	
	if dummyPlayer != nil {
		ConnectToPOServer()
	}
}

func LoadDummyPlayer() {
	dummyPlayer = NewPlayer("mr_dark")
	if dummyPlayer.LoadData() {
		fmt.Println("Player data loaded")
	} else {
		dummyPlayer = nil
	}
}

func ConnectToPOServer() {
	client, _ = NewPOClient(dummyPlayer)
	client.Connect()
}