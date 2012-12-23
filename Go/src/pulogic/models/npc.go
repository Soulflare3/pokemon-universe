package models

const (
	Npc_Idnpc      string = "npc.idnpc"
	Npc_Name       string = "npc.name"
	Npc_ScriptName string = "npc.script_name"
	Npc_Position   string = "npc.position"
	Npc_Idmap      string = "npc.idmap"
)

type Npc struct {
	Idnpc      int `PK`
	Name       string
	ScriptName string
	Position   int64
	Idmap      int
}

type NpcJoinOutfitJoinEvent struct {
	Idnpc      int `PK`
	Name       string
	ScriptName string
	Position   int64
	Idmap      int
	
	Head  		int
	Nek   		int
	Upper 		int
	Lower 		int
	Feet  		int
	
	Event	string
	Initid	int
}