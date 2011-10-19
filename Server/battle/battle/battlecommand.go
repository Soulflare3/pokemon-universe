package main

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
	BattleCommand_BattleChat
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