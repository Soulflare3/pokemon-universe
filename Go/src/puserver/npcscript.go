package main

type NpcScript struct {
	self *Npc
}

func NewNpcScript(_npc *Npc) *NpcScript {
	return &NpcScript{ self: _npc }
}

func (s *NpcScript) GetCreatureName(cid uint64) (name string) {
	player, ok := g_game.GetPlayerByGuid(cid)
	if ok {
		name = player.GetName()
	} else {
		name = "Unknown"
	}
	return
}

func (s *NpcScript) SelfSay(message string) {
	s.self.SelfSay(message)
}

// Dialogue
func (s *NpcScript) SendDialogue(cid uint64, title string, options ...string) {

}

func (s *NpcScript) HideDialogue(cid uint64) {

}

func (s *NpcScript) EndDialogue(cid uint64) {

}

// Pokecenter
func (s *NpcScript) HealParty(cid uint64) {
	player, ok := g_game.GetPlayerByGuid(cid)
	if ok {
		player.HealParty()
	}
}

// Market
func (s *NpcScript) OpenShopWindow(cid uint64) {
}

func (s *NpcScript) CloseShopWindow(cid uint64) {
}

// Quest
func (s *NpcScript) GetQuestProgress(cid uint64, questId int) int {
	return 0
}

func (s *NpcScript) SetQuestProgress(cid uint64, questId int, progress int) {

}

// Items
func (s *NpcScript) AddItem(cid uint64, itemId int, amount int) {
}

func (s *NpcScript) CheckItem(cid uint64, itemId, amount int) bool {
	return false
}

func (s *NpcScript) RemoveItem(cid uint64, itemId int, amount int) {

}

// Golds
func (s *NpcScript) AddMoney(cid uint64, amount int) {
	player, ok := g_game.GetPlayerByGuid(cid)
	if ok {
		player.SetMoney(amount)
	}
}

func (s *NpcScript) CheckMoney(cid uint64, amount int) bool {
	player, ok := g_game.GetPlayerByGuid(cid)
	if ok && player.Money >= amount {
		return true
	}
	
	return false
}

func (s *NpcScript) RemoveMoney(cid uint64, amount int) {
	player, ok := g_game.GetPlayerByGuid(cid)
	if ok {
		player.SetMoney(-amount)
	}
}