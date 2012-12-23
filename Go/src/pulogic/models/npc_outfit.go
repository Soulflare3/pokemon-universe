package models

const (
	NpcOutfit_Idnpc string = "npc_outfit.idnpc"
	NpcOutfit_Head  string = "npc_outfit.head"
	NpcOutfit_Nek   string = "npc_outfit.nek"
	NpcOutfit_Upper string = "npc_outfit.upper"
	NpcOutfit_Lower string = "npc_outfit.lower"
	NpcOutfit_Feet  string = "npc_outfit.feet"
)

type NpcOutfit struct {
	Idnpc int `PK`
	Head  int
	Nek   int
	Upper int
	Lower int
	Feet  int
}
