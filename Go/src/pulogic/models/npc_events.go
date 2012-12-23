package models

const (
	NpcEvents_Idnpc  string = "npc_events.idnpc"
	NpcEvents_Event  string = "npc_events.event"
	NpcEvents_Initid string = "npc_events.initid"
)

type NpcEvents struct {
	Idnpc	int `PK`
	Event	string
	Initid	int
}