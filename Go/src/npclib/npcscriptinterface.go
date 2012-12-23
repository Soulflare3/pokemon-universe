package npclib

type NpcScriptInterface interface {
	GetCreatureName(cid uint64) string
	SelfSay(message string)

	// Dialogue
	SendDialogue(cid uint64, title string, options ...string)
	HideDialogue(cid uint64)
	EndDialogue(cid uint64)	
	
	// Pokecenter
	HealParty(cid uint64)
	
	// Market
	OpenShopWindow(cid uint64)
	CloseShopWindow(cid uint64)
	
	// Quest
	GetQuestProgress(cid uint64, questId int) int
	SetQuestProgress(cid uint64, questId int, progress int)
	
	// Items
	AddItem(cid uint64, itemId int, amount int) bool
	CheckItem(cid uint64, itemId int, amount int) bool
	RemoveItem(cid uint64, itemId int, amount int) bool
	
	// Golds
	AddMoney(cid uint64, amount int)
	CheckMoney(cid uint64, amount int) bool
	RemoveMoney(cid uint64, amount int)
	
	// Battle
	StartBattle(cid uint64)
}

type NpcInteractionInterface interface {
	// Impleneted in base class
	SetScriptInterface(script NpcScriptInterface)
	
	// General dialogue answer
	OnAnswer(cid uint64, answer int)
	
	// Market stuff
	OnBuy(cid uint64, callback int)
	OnShopWindowClose(cid uint64)
	
	// Battle
	OnBattleWin(cid uint64)
	OnBattleLose(cid uint64)
}