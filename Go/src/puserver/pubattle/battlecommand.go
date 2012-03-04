/*Pokemon Universe MMORPG
Copyright (C) 2010 the Pokemon Universe Authors

This program is free software; you can redistribute it and/or
modify it under the terms of the GNU General Public License
as published by the Free Software Foundation; either version 2
of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program; if not, write to the Free Software
Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA  02110-1301, USA.*/
package pubattle

const (
	BattleCommand_SendOut int = iota // 0
	BattleCommand_SendBack
	BattleCommand_UseAttack
	BattleCommand_OfferChoice
	BattleCommand_BeginTurn
	BattleCommand_ChangePP
	BattleCommand_ChangeHp
	BattleCommand_Ko
	BattleCommand_Effective // to tell how a move is effective
	BattleCommand_Miss
	BattleCommand_CriticalHit // 10
	BattleCommand_Hit         // for moves like fury double kick etc. 
	BattleCommand_StatChange
	BattleCommand_StatusChange
	BattleCommand_StatusMessage
	BattleCommand_Failed
	BattleCommand_BattleChat // Obsolete, we use our own chat
	BattleCommand_MoveMessage
	BattleCommand_ItemMessage
	BattleCommand_NoOpponent
	BattleCommand_Flinch // 20
	BattleCommand_Recoil
	BattleCommand_WeatherMessage
	BattleCommand_StraightDamage
	BattleCommand_AbilityMessage
	BattleCommand_AbsStatusChange
	BattleCommand_Substitute
	BattleCommand_BattleEnd
	BattleCommand_BlankMessage
	BattleCommand_CancelMove
	BattleCommand_Clause       // 30
	BattleCommand_DynamicInfo  // 31
	BattleCommand_DynamicStats // 32
	BattleCommand_Spectating
	BattleCommand_SpectatorChat
	BattleCommand_AlreadyStatusMessage
	BattleCommand_TempPokeChange
	BattleCommand_ClockStart // 37
	BattleCommand_ClockStop  // 38
	BattleCommand_Rated
	BattleCommand_TierSection // 40
	BattleCommand_EndMessage
	BattleCommand_PointEstimate
	BattleCommand_MakeYourChoice
	BattleCommand_Avoid
	BattleCommand_RearrangeTeam
	BattleCommand_SpotShifts
)