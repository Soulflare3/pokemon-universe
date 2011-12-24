package pu_npclib

type NPC_NurseJoy struct { NpcBase }

func NewNPC_NurseJoy() *NPC_NurseJoy {
	return &NPC_NurseJoy { }
}

func (n *NPC_NurseJoy) OnAnswer(cid uint64, answer int) {
	if answer == 0 {
		n.Script.SendDialogue(cid, 
							"Hello " + n.Script.GetCreatureName(cid) + ", how can I help you?",
							"1-Could you heal my Pokemon?", 
							"2-No, thanks.")
	} else if answer == 1 {
		n.Script.HealParty(cid)
		n.Script.SendDialogue(cid, "Of course! Your Pokemon are now fully healed")
		n.Script.EndDialogue(cid)
	} else if answer == 2 {
		n.Script.SendDialogue(cid, "Good bye then!")
		n.Script.EndDialogue(cid)
	}
}