package main

const (
    COMMAND_WhatAreYou int = iota // = 0
    COMMAND_WhoAreYou
    COMMAND_Login
    COMMAND_Logout
    COMMAND_SendMessage
    COMMAND_PlayersList
    COMMAND_SendTeam
    COMMAND_ChallengeStuff
    COMMAND_EngageBattle
    COMMAND_BattleFinished
    COMMAND_BattleMessage // = 10
    COMMAND_BattleChat
    COMMAND_KeepAlive // obsolete since we use a native Qt option now
    COMMAND_AskForPass
    COMMAND_Register
    COMMAND_PlayerKick
    COMMAND_PlayerBan
    COMMAND_ServNumChange
    COMMAND_ServDescChange
    COMMAND_ServNameChange
    COMMAND_SendPM // = 20
    COMMAND_Away
    COMMAND_GetUserInfo
    COMMAND_GetUserAlias
    COMMAND_GetBanList
    COMMAND_CPBan
    COMMAND_CPUnban
    COMMAND_SpectateBattle
    COMMAND_SpectatingBattleMessage
    COMMAND_SpectatingBattleChat
    COMMAND_SpectatingBattleFinished // = 30
    COMMAND_LadderChange
    COMMAND_ShowTeamChange
    COMMAND_VersionControl
    COMMAND_TierSelection
    COMMAND_ServMaxChange
    COMMAND_FindBattle
    COMMAND_ShowRankings
    COMMAND_Announcement
    COMMAND_CPTBan
    COMMAND_CPTUnban // = 40
    COMMAND_PlayerTBan
    COMMAND_GetTBanList
    COMMAND_BattleList
    COMMAND_ChannelsList
    COMMAND_ChannelPlayers
    COMMAND_JoinChannel
    COMMAND_LeaveChannel
    COMMAND_ChannelBattle
    COMMAND_RemoveChannel
    COMMAND_AddChannel // = 50
    COMMAND_ChannelMessage
    COMMAND_ChanNameChange
    COMMAND_HtmlMessage
    COMMAND_HtmlChannel
    COMMAND_ServerName
    COMMAND_SpecialPass
    COMMAND_ServerListEnd	// Indicates end of transmission for registry.
    COMMAND_SetIP			// Indicates that a proxy server sends the real ip of client
)