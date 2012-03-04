package npclib

type NpcLib struct {
	npcList map[string]NpcInteractionInterface
}

func NewNpcLib() *NpcLib {	
	npcLib := NpcLib { npcList: make(map[string]NpcInteractionInterface) }
	npcLib.init()
	
	return &npcLib
}

func (lib *NpcLib) init() {
	// Load Npc Scripts
	lib.npcList["NurseJoy"] = NewNPC_NurseJoy()
}

func (lib *NpcLib) GetNpcScript(_name string) (NpcInteractionInterface, bool) {
	npc, found := lib.npcList[_name]
	return npc, found
}