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

/*const (
	COMMAND_ZipCommand int = iota // = 0
	COMMAND_Login
	COMMAND_Reconnect
	COMMAND_Logout
	COMMAND_SendMessage
	COMMAND_PlayersList
	COMMAND_SendTeam
	COMMAND_ChallengeStuff
	COMMAND_EngageBattle
	COMMAND_BattleFinished
	COMMAND_BattleMessage //= 10
	COMMAND_BattleChat
	COMMAND_KeepAlive // obsolete since we use a native Qt option now
	COMMAND_AskForPass
	COMMAND_Register
	COMMAND_PlayerKick
	COMMAND_PlayerBan
	COMMAND_ServNumChangeServerInfoChanged // = COMMAND_ServNameChange
	COMMAND_ServDescChange
	COMMAND_ServNameChange
	COMMAND_SendPM //= 20
	COMMAND_OptionsChange
	COMMAND_GetUserInfo
	COMMAND_GetUserAlias
	COMMAND_GetBanList
	COMMAND_CPBan
	COMMAND_CPUnban
	COMMAND_SpectateBattle
	COMMAND_SpectatingBattleMessage
	COMMAND_SpectatingBattleChat
	COMMAND_Unused30 //= 30
	COMMAND_Unused31
	COMMAND_Unused32
	COMMAND_VersionControl_
	COMMAND_TierSelection
	COMMAND_ServMaxChange
	COMMAND_FindBattle
	COMMAND_ShowRankings
	COMMAND_Announcement
	COMMAND_CPTBan
	COMMAND_CPTUnban //= 40
	COMMAND_PlayerTBan
	COMMAND_Unused42
	COMMAND_BattleList
	COMMAND_ChannelsList
	COMMAND_ChannelPlayers
	COMMAND_JoinChannel
	COMMAND_LeaveChannel
	COMMAND_ChannelBattle
	COMMAND_RemoveChannel
	COMMAND_AddChannel //= 50
	COMMAND_Unused51
	COMMAND_ChanNameChange
	COMMAND_Unused53
	COMMAND_Unused54
	COMMAND_ServerName
	COMMAND_SpecialPass
	COMMAND_ServerListEnd // Indicates end of transmission for registry.
	COMMAND_SetIP         // Indicates that a proxy server sends the real ip of client
	COMMAND_ServerPass    // Prompts for the server password
)*/

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
