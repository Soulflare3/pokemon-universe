package npclib

var npcList map[string]NpcInteractionInterface

func init() {
	npcList = make(map[string]NpcInteractionInterface)

	// Load Npc Scripts
	npcList["NurseJoy"] = NewNPC_NurseJoy()
}

func GetNpcScript(_name string) (NpcInteractionInterface, bool) {
	npc, found := npcList[_name]
	return npc, found
}