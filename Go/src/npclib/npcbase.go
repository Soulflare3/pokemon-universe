package npclib

type NpcBase struct {
	Script NpcScriptInterface
}

func (n *NpcBase) SetScriptInterface(_script NpcScriptInterface) {
	n.Script = _script
}

func (n *NpcBase) OnBuy(cid uint64, callback int) {
}

func (n *NpcBase) OnShopWindowClose(cid uint64) {
}

func (n *NpcBase) OnBattleWin(cid uint64) {
}

func (n *NpcBase) OnBattleLose(cid uint64) {
}