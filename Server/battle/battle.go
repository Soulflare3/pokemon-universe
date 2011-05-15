package main

import (
	"container/list"
	"time"
	"fmt"
	"strings"
	pnet "network"
)

const (
	STATUSMESSAGE_PARALYZED = "%s is paralyzed! It may be unable to move!"
	STATUSMESSAGE_ASLEEP = "%s fell asleep!"
	STATUSMESSAGE_FROZEN = "%s was frozen solid!"
	STATUSMESSAGE_BURNED = "%s was burned!"
	STATUSMESSAGE_POISONED = "%s was poisoned!"
	STATUSMESSAGE_POISONEDBADLY = "%s was badly poisoned!"
)

type Battle struct {
	battleId	int32
	id1, id2	int32
	delayed		int
	
	delayedCommands	*list.List
	started			bool
	battleEnded		bool
	
	info			*BattleInfo
	conf			*BattleConfiguration
	
	owner	*POClient
	
	usePokemonNames	bool
	
	statusChangeMessages []string
}

func NewBattle(_owner *POClient, _battleId int32, _me *PlayerInfo, _opponent *PlayerInfo, _team *TeamBattle, _conf *BattleConfiguration) *Battle {
	battle := &Battle{ owner: _owner,
						battleId: _battleId,
						id1: _me.id, 
						id2: _opponent.id,
						started: false,
						delayedCommands: list.New(),
						usePokemonNames: true,
						statusChangeMessages: make([]string, 6) }
	
	battle.conf = _conf
	battle.info = NewBattleInfo(_team, _me, _opponent, _conf.mode, _conf.spot(_me.id), _conf.spot(_opponent.id))
	battle.info.gen = _conf.gen
	
	battle.statusChangeMessages[0] = STATUSMESSAGE_PARALYZED
	battle.statusChangeMessages[1] = STATUSMESSAGE_ASLEEP
	battle.statusChangeMessages[2] = STATUSMESSAGE_FROZEN
	battle.statusChangeMessages[3] = STATUSMESSAGE_BURNED
	battle.statusChangeMessages[4] = STATUSMESSAGE_POISONED
	battle.statusChangeMessages[5] = STATUSMESSAGE_POISONEDBADLY
	
	return battle
}

// Default forceDelay is TRUE
func (b *Battle) Delay(msec int64, forceDelay bool) {
	b.delayed += 1
	
	if !forceDelay && b.delayed > 1 {
		b.delayed = 1;
	}
	
	if msec != 0 {
		delayNs := msec * 1e6
		time.AfterFunc(delayNs, func() { b.Undelay() })
	}
}

func (b *Battle) Undelay() {
	if b.delayed > 0 {
		b.delayed -= 1
	} else {
		return
	}
	
	for (b.delayed == 0 && b.delayedCommands.Len() > 0) {
		element := b.delayedCommands.Front()
		command := element.Value.(*pnet.QTPacket)
		b.ReceiveInfo(command)
		b.delayedCommands.Remove(element)
	}
}

func (b *Battle) Name(_spot int8) string {
	return b.info.Name(_spot)
}

func (b *Battle) Nick(_player int8) string {
	playerName := b.Name(b.info.Player(_player))
	pokemonName := b.RNick(_player)
	
	return fmt.Sprintf("%s's %s", playerName, pokemonName)
}

func (b *Battle) RNick(_player int8) string {
	pokenum := b.info.CurrentShallow(_player).num
	return g_PokemonInfo.GetPokemonName(pokenum)
}

func (b *Battle) Player(_spot int8) int8 {
	return b.info.Player(_spot)
}

func (b *Battle) Opponent(_spot int8) (ret int8) {
	ret = 1
	if _spot == 1 {
		ret = 0
	}
	return
}

func (b *Battle) ReceiveInfo(_packet *pnet.QTPacket) {
	commandPeek := _packet.Buffer[_packet.ReadPos]
	if b.delayed > 0 && commandPeek != BattleCommand_BattleChat && commandPeek != BattleCommand_SpectatorChat &&
					commandPeek != BattleCommand_ClockStart && commandPeek != BattleCommand_ClockStop &&
					commandPeek != BattleCommand_Spectating {
		b.delayedCommands.PushBack(_packet)
		return
	}
	
	// At the start of the battle we wait 700ms, to prevent misclicks
	// when wanting to do something else
	if !b.started && commandPeek == BattleCommand_OfferChoice {
		b.started = true
		b.Delay(700, true)
	}
		
	command := uint8(_packet.ReadUint8())
	player := int8(_packet.ReadUint8())
	
	b.DealWithCommandInfo(_packet, command, player, player)
}

func (b *Battle) DealWithCommandInfo(_packet *pnet.QTPacket, _command uint8, _spot int8, _truespot int8) {
	switch _command {
		default:
			fmt.Printf("[Warning] Received unknown battle command %v (player %v)\n\r", _command, _spot)
		case BattleCommand_SendOut: // 0
			b.CommandSendOut(_packet, _spot)
		case BattleCommand_UseAttack: // 2
			b.CommandUseAttack(_packet, _spot)
		case BattleCommand_OfferChoice: // 3
			b.CommandOfferChoice(_packet)
		case BattleCommand_BeginTurn: // 4
			b.CommandBeginTurn(_packet) 
		case BattleCommand_ChangePP:
			b.CommandChangePP(_packet, _spot)
		case BattleCommand_ChangeHp: // 6
			b.CommandChangeHp(_packet, _spot)
		case BattleCommand_Ko: // 7
			b.CommandKo(_spot)
		case BattleCommand_Effective: // 8
			b.CommandEffective(_packet)
		case BattleCommand_CriticalHit:
			b.CommandCriticalHit()
		case BattleCommand_StatusChange: // 13
			b.CommandStatusChange(_packet, _spot)
		case BattleCommand_MoveMessage: // 17
			b.CommandMoveMessage(_packet, _spot)
		case BattleCommand_AbsStatusChange: // 25
			b.CommandAbsStatusChange(_packet, _spot)	
		case BattleCommand_StraightDamage: // 26
			b.CommandStraightDamage(_packet, _spot)
		case BattleCommand_BattleEnd: // 27
			b.CommandBattleEnd(_packet, _spot)
		case BattleCommand_BlankMessage: // 28
			 fmt.Println("")
		case BattleCommand_DynamicInfo: // 31
			b.CommandDynamicInfo(_packet, _spot)
		case BattleCommand_DynamicStats: // 32
			b.CommandDynamicStats(_packet, _spot)
		case BattleCommand_ClockStart: // 37
			b.CommandClockStart(_packet, _spot)
		case BattleCommand_ClockStop: // 38
			b.CommandClockStop(_packet, _spot)
		case BattleCommand_Rated: // 39
			fmt.Println("[Rated] Rules and clauses shii")
		case BattleCommand_TierSection: // 40
			tier := _packet.ReadString()
			fmt.Printf("[TierSelection] Tier: %v\n", tier)
		case BattleCommand_BattleChat:
			fallthrough
		case BattleCommand_EndMessage:
			b.CommandBattleChat(_packet, _spot)
		case BattleCommand_MakeYourChoice:
			b.CommandMakeYourChoice()
	}
}

func (b *Battle) CommandSendOut(_packet *pnet.QTPacket, spot int8) {
	silent := (_packet.ReadUint8() == 1)
	prevIndex := int8(_packet.ReadUint8())
	
	b.info.sub[spot] = false
	b.info.specialSprite[spot] = PokemonName_NoPoke
	
	b.info.SwitchPoke(spot, int8(prevIndex))
	shallow := NewPokeBattleFromPacket(_packet)
	b.info.SetCurrentShallow(spot, shallow)
	b.info.pokeAlive[spot] = true
	
	// TODO: Send update display to PU Client
	// mydisplay->updatePoke(spot);
	// mydisplay->updatePoke(info().player(spot), info().slotNum(spot));
	// mydisplay->updatePoke(info().player(spot), prevIndex);
	
	if !silent {
		pokename := g_PokemonInfo.GetPokemonName(shallow.num)
		othername := b.RNick(spot)
		if pokename != othername {
			fmt.Printf("%v sent out %v! (%3)\n", b.Name(b.info.Player(spot)), othername, pokename)
		} else {
			fmt.Printf("%v sent out %v!\n", b.Name(b.info.Player(spot)), pokename)
		}
	}
}

func (b *Battle) CommandUseAttack(_packet *pnet.QTPacket, _spot int8) {
	attackId := _packet.ReadUint16()
	
	user := b.Nick(_spot)
	attackName := g_MoveInfo.GetMoveName(attackId)
	
	fmt.Printf("%s used %s!\n", user, attackName)
}

func (b *Battle) CommandOfferChoice(_packet *pnet.QTPacket) {
	if b.info.sent {
		b.info.sent = false
		for i := 0; i < b.info.available.Len(); i++ {
			b.info.available.Set(i, false)
			b.info.done.Set(i, false)
		}
	}
	
	c := NewBattleChoices()
	c.initFromPackage(_packet)
	b.info.choices.Set(int(c.numSlot/2), c)
	b.info.available.Set(int(c.numSlot/2), true)
}

func (b *Battle) CommandBeginTurn(_packet *pnet.QTPacket) {
	turn := _packet.ReadUint32()
	fmt.Println("")
	fmt.Printf("Start of turn %d\n", turn)
}

func (b *Battle) CommandChangePP(_packet *pnet.QTPacket, _spot int8) {
	move := _packet.ReadUint8()
	pp := _packet.ReadUint8()
	
	b.info.CurrentPoke(_spot).moves[move].pp = pp
	b.info.GetTempPoke(_spot).moves[move].pp = pp
	
	fmt.Printf("Changed move PP to %d\n", pp)
	
	// myazones[info().number(spot)]->tattacks[move]->updateAttack(info().tempPoke(spot).move(move), info().tempPoke(spot), gen());
    // mypzone->pokes[info().number(spot)]->updateToolTip();
}

func (b *Battle) CommandChangeHp(_packet *pnet.QTPacket, _spot int8) {
	goal := _packet.ReadUint16()
	b.info.CurrentShallow(_spot).lifePercent = uint8(goal)
}

func (b *Battle) CommandKo(_spot int8) {
	pokemonName := b.Nick(_spot)
	fmt.Printf("%s fainted!\n", pokemonName)
	
	b.switchToNaught(_spot)
}

func (b *Battle) CommandEffective(_packet *pnet.QTPacket) {
	eff := _packet.ReadUint8()
	
	if eff == 0 {
		fmt.Println("It had no effect!")
	} else if eff == 1 || eff == 2 {
		fmt.Println("It's not very effective...")
	} else if eff == 8 || eff == 16 {
		fmt.Println("It's super effective!")
	}
}

func (b *Battle) CommandCriticalHit() {
	fmt.Println("A critical hit!")
}

func (b *Battle) CommandStatusChange(_packet *pnet.QTPacket, _spot int8) {
	status := int8(_packet.ReadUint8())
	multipleTurns := (_packet.ReadUint8() == 1)
	if status > PokemonStatus_Fine && status <= PokemonStatus_Poisoned {
		messageId := status - 1
		if (status == PokemonStatus_Poisoned && multipleTurns) {
			messageId++
		}
		fmt.Printf(b.statusChangeMessages[messageId] + "\n", b.Nick(_spot))
	} else if status == PokemonStatus_Confused {
		fmt.Printf("%s became confused!\n", b.Nick(_spot))
	}
}

func (b *Battle) CommandMoveMessage(_packet *pnet.QTPacket, _spot int8) {
	move := _packet.ReadUint16()
	part := _packet.ReadUint8()
	moveType := int8(_packet.ReadUint8())
	foe := int8(_packet.ReadUint8())
	other := _packet.ReadUint16()
	q := _packet.ReadString()
	
	message := g_MoveInfo.GetMoveMessage(move, part)
	strings.Replace(message, "%s", b.Nick(_spot), 1)
	strings.Replace(message, "%ts", b.Name(b.Player(_spot)), 1)
	strings.Replace(message, "%tf", b.Name(b.Opponent(b.Player(_spot))), 1)
	strings.Replace(message, "%t", g_TypeInfo.GetTypeName(moveType), 1)
	strings.Replace(message, "%f", b.Nick(foe), 1)
	strings.Replace(message, "%m", g_MoveInfo.GetMoveName(other), 1)
	strings.Replace(message, "%d", string(other), 1)
	strings.Replace(message, "%q", q, 1)
	strings.Replace(message, "%i", g_ItemInfo.GetItemName(other), 1)
	strings.Replace(message, "%a", g_AbilityInfo.GetAbilityName(other), 1)
	strings.Replace(message, "%p", g_PokemonInfo.GetPokemonName(NewPokemonUniqueIdFromRef(uint32(other))), 1)
	
	fmt.Println(message)
}

func (b *Battle) CommandAbsStatusChange(_packet *pnet.QTPacket, _spot int8) {
	poke := int8(_packet.ReadUint8())
	status := _packet.ReadUint8()
	
	if poke < 0 || poke >= 6 {
		return
	}
	
	if status != PokemonStatus_Confused {
		// fmt.Printf("_spot: %d | poke %d\n", _spot, poke)
		if b.info.pokemons[_spot][poke] == nil {
			// fmt.Printf("pokemons[%d][%d] == nil\n", _spot, poke)
			return
		}
		b.info.pokemons[_spot][poke].ChangeStatus(status)
		if b.info.IsOut(poke) {
			// TODO: mydisplay->updatePoke(b.info.slot(_spot, poke))
		}
		// TODO: mydisplay->changeStatus(_spot, poke, status)
	}
}

func (b *Battle) CommandStraightDamage(_packet *pnet.QTPacket, _spot int8) {
	damage := int16(_packet.ReadUint16())
	pokemonName := b.Nick(_spot)
	fmt.Printf("%s lost %d of its health!\n", pokemonName, damage)
}

func (b *Battle) CommandBattleEnd(_packet *pnet.QTPacket, _spot int8) {
	res := int8(_packet.ReadUint8())
	b.battleEnded = true
	
	if res == BattleResult_Tie {
		fmt.Printf("Tie between %s and %s!\n", b.Name(b.info.myself), b.Name(b.info.opponent))
	} else {
		fmt.Printf("%s won the battle!\n", b.Name(_spot))
	}
}

func (b *Battle) CommandDynamicInfo(_packet *pnet.QTPacket, _spot int8) {
	dynamicInfo := NewBattleDynamicInfoFromPacket(_packet)
	b.info.statChanges.Insert(int(_spot), dynamicInfo)
	
	// mydisplay->updateToolTip(spot)
}

func (b *Battle) CommandDynamicStats(_packet *pnet.QTPacket, _spot int8) {
	battleStats := NewBattleStatsFromPacket(_packet)
	b.info.mystats.Insert(int(_spot), battleStats)
}

func (b *Battle) CommandClockStart(_packet *pnet.QTPacket, _spot int8) {
	b.info.time[_spot] = _packet.ReadUint16()
	b.info.startingTime[_spot] = time.Seconds()
	b.info.ticking[_spot] = true
}

func (b *Battle) CommandClockStop(_packet *pnet.QTPacket, _spot int8) {
	b.info.time[_spot] = _packet.ReadUint16()
	b.info.ticking[_spot] = false
}

func (b *Battle) CommandBattleChat(_packet *pnet.QTPacket, _spot int8) {
	message := _packet.ReadString()
	if len(message) > 0 {
		fmt.Printf("%s: %s\n", b.Name(_spot), message)
	}
}

func (b *Battle) CommandMakeYourChoice() {
	b.info.possible = true
	b.info.sent = true
	
	fmt.Println("Make a choice...")
	
	b.attackClicked(0)
}

// -------------------------------------------------------------------------------- //
func (b *Battle) goToNextChoice() {
	for i := 0; i < b.info.available.Len(); i++ {
		slot := b.info.Slot(b.info.myself, int8(i))
		
		if b.info.available.At(i).(bool) && !b.info.done.At(i).(bool) {
			b.enableAll()
			
			b.info.currentSlot = slot
			
			//  myswitch->setText(tr("&Switch Pokemon"));
			if b.info.choices.At(i).(*BattleChoices).attacksAllowed == false && b.info.choices.At(i).(*BattleChoices).switchAllowed == true {
				// mytab->setCurrentIndex(PokeTab)
			} else {
				b.switchTo(b.info.Number(slot), slot, false)
				if b.info.mode == ChallengeInfo_Triples && i != 1 {
					// myswitch->setText(tr("&Shift to centre"))
				}
			}
			
			// Moves first
			if b.info.pokeAlive.At(int(slot)).(bool) {
				if b.info.choices.At(i).(*BattleChoices).attacksAllowed == false {
					// myattack->setEnabled(false)
					for j := 0; j < 4; j++ {
						// myazones[info().number(slot)]->attacks[j]->setEnabled(false);
					}
				} else {
					// myattack->setEnabled(true)
					for j := 0; j < 4; j++ {
                        // myazones[info().number(slot)]->attacks[j]->setEnabled(info().choices[i].attackAllowed[i]);
                    }
					
					if b.info.choices.At(i).(*BattleChoices).Struggle() {
						// mystack->setCurrentWidget(szone);
					} else {
						// mystack->setCurrentWidget(myaxones[info().number(slot)]);
					}
				}
			}
			
			// Then Pokemon
			if b.info.choices.At(i).(*BattleChoices).switchAllowed == false {
				// myswitch->setEnabled(false)
				// mypzone->setEnabled(false)
			} else {
				// myswitch->setEnabled(true)
				for j := 0; j < 6; j++ {
                    // mypzone->pokes[i]->setEnabled(team().poke(i).num() != 0 && team().poke(i).lifePoints() > 0 && team().poke(i).status() != Pokemon::Koed);
                }
				
				if b.info.Multiples() {
					// In doubles, whatever happens, you can't switch to your partner
					for j := 0; j < (b.info.numberOfSlots / 2); j++ {
						// mypzone->pokes[i].setEnabled(false)
					}
					
					// Also you can't switch to a pokemon you've chosen before
					for j := 0; j < b.info.available.Len(); j++ {
						if b.info.available.At(j).(bool) && b.info.done.At(j).(bool) && b.info.choice.At(j).(*BattleChoice).SwitchChoice() {
							// mypzone->pokes[info().choice[j].pokeSlot()]->setEnabled(false);
						}
					}
				}
			}
			
			return
		}
	}
	
	// myattack->setEnabled(false);
    // myswitch->setEnabled(false);
	
	b.disableAll()
	
	for key, element := range b.info.available {
		if element.(bool) == true {
			b.sendChoice(b.info.choice.At(key).(*BattleChoice))
		}
	}
}

func (b *Battle) enableAll() {
	// mypzone->setEnabled(true);
    // for (int i = 0; i < 3; i++)
		// myazones[i]->setEnabled(true);
    // if (info().multiples())
       // tarZone->setEnabled(true);
}

func (b *Battle) disableAll() {
	// mypzone->setEnabled(false);
    // for (int i = 0; i < 3; i++)
		// myazones[i]->setEnabled(false);
	// if (info().multiples())
		// tarZone->setEnabled(false);
}

func (b *Battle) sendChoice(_choice *BattleChoice) {
	b.owner.SendBattleChoice(b.battleId, _choice)
	// emit battleCommand(battleId(), _choice)
	b.info.possible = false
}

func (b *Battle) switchTo(_pokezone int8, _spot int8, _forced bool) {
	snum := b.info.Number(_spot)
	
	if snum != _pokezone || _forced {
		b.info.SwitchPokeExt(_spot, _pokezone, true)
		// mypzone->pokes[snum]->changePokemon(info().myteam.poke(snum));
        // mypzone->pokes[pokezone]->changePokemon(info().myteam.poke(pokezone));
	}
	
	// mystack->setCurrentIndex(info().number(spot));
    // mytab->setCurrentIndex(MoveTab);

    // mydisplay->updatePoke(spot);
	
	// for (int i = 0; i< 4; i++) {
        // myazones[info().number(spot)]->tattacks[i]->updateAttack(info().tempPoke(spot).move(i), info().tempPoke(spot), gen());
    // }	
}

func (b* Battle) attackClicked(_zone int8) {
	slot := b.info.currentSlot
	
	if _zone != -1 { // Struggle
		b.info.lastMove[b.info.Number(slot)] = _zone
	}
	
	if b.info.possible {
		attack := AttackChoice{}
		attack.attackSlot = _zone
		attack.attackTarget = b.info.Number(b.info.opponent)
		choice := NewBattleChoiceAttack(uint8(slot), attack)
		b.info.choice[b.info.Number(slot)] = choice
		
		if !b.info.Multiples() {
			b.info.done[b.info.Number(slot)] = true
			b.goToNextChoice()
		}
	}
}

func (b *Battle) switchToNaught(_spot int8) {
	b.info.pokeAlive.Set(int(_spot), false)
	
	// mydisplay->updatePoke(spot)
}