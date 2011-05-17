/*Pokemon Universe MMORPG
Copyright (C) 2010 the Pokemon Universe Authors

This program is free software; you can redistribute it and/or
modify it under the terms of the GNU General Public License
as published by the Free Software Foundation; either version 2
of the License or (at your option) any later version.

This program is distributed in the hope that it will be useful
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program; if not write to the Free Software
Foundation Inc. 51 Franklin Street Fifth Floor Boston MA  02110-1301 USA.*/
package main

const (
	WhatAreYou uint8 = iota // 0
	WhoAreYou
	Login // PU Obsolete
	Logout
	SendMessage
	PlayersList // PU Obsolete
	SendTeam
	ChallengeStuff
	EngageBattle
	BattleFinished
	BattleMessage // 10
	BattleChat
	KeepAlive // obsolete since we use a native Qt option now
	AskForPass
	Register
	PlayerKick
	PlayerBan
	ServNumChange
	ServDescChange
	ServNameChange
	SendPM // 20
	Away
	GetUserInfo
	GetUserAlias
	GetBanList
	CPBan
	CPUnban
	SpectateBattle
	SpectatingBattleMessage
	SpectatingBattleChat
	SpectatingBattleFinished // 30
	LadderChange
	ShowTeamChange
	VersionControl // PU Obsolete
	TierSelection // PU Obsolete
	ServMaxChange
	FindBattle
	ShowRankings
	Announcement
	CPTBan
	CPTUnban // 40
	PlayerTBan
	GetTBanList
	BattleList // PU Obsolete
	ChannelsList // PU Obsolete
	ChannelPlayers // PU Obsolete
	JoinChannel // PU Obsolete
	LeaveChannel
	ChannelBattle
	RemoveChannel
	AddChannel // 50
	ChannelMessage // PU Obsolete
	ChanNameChange
	HtmlMessage // PU Obsolete
	ChannelHtml
)

const (
	BattleCommand_SendOut uint8 = iota // 0
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
	BattleCommand_Hit // for moves like fury double kick etc. 
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
	BattleCommand_Clause // 30
	BattleCommand_DynamicInfo // 31
	BattleCommand_DynamicStats // 32
	BattleCommand_Spectating
	BattleCommand_SpectatorChat
	BattleCommand_AlreadyStatusMessage
	BattleCommand_TempPokeChange
	BattleCommand_ClockStart // 37
	BattleCommand_ClockStop // 38
	BattleCommand_Rated
	BattleCommand_TierSection // 40
	BattleCommand_EndMessage
	BattleCommand_PointEstimate
	BattleCommand_MakeYourChoice
	BattleCommand_Avoid
	BattleCommand_RearrangeTeam
	BattleCommand_SpotShifts
)

const (
	BattleResult_Forfeit int8 = iota
	BattleResult_Win
	BattleResult_Tie
	BattleResult_Close
)

const (
	ChallengeInfo_Singles uint8 = iota
	ChallengeInfo_Doubles
	ChallengeInfo_Triples
	ChallengeInfo_Rotation
	ChallengeInfo_ModeFirst = ChallengeInfo_Singles
	ChallengeInfo_ModeLast = ChallengeInfo_Rotation
)

const (
	PokemonName_NoPoke = iota
	PokemonName_Bulbasaur
	PokemonName_Ivysaur
	PokemonName_Venusaur
	PokemonName_Charmander
	PokemonName_Charmeleon
	PokemonName_Charizard
	PokemonName_Squirtle
	PokemonName_Wartortle
	PokemonName_Blastoise
	PokemonName_Caterpie
	PokemonName_Metapod
	PokemonName_Butterfree
	PokemonName_Weedle
	PokemonName_Kakuna
	PokemonName_Beedrill
	PokemonName_Pidgey
	PokemonName_Pidgeotto
	PokemonName_Pidgeot
	PokemonName_Rattata
	PokemonName_Raticate
	PokemonName_Spearow
	PokemonName_Fearow
	PokemonName_Ekans
	PokemonName_Arbok
	PokemonName_Pikachu
	PokemonName_Raichu
	PokemonName_Sandshrew
	PokemonName_Sandslash
	PokemonName_Nidoran_F
	PokemonName_Nidorina
	PokemonName_Nidoqueen
	PokemonName_Nidoran_M
	PokemonName_Nidorino
	PokemonName_Nidoking
	PokemonName_Clefairy
	PokemonName_Clefable
	PokemonName_Vulpix
	PokemonName_Ninetales
	PokemonName_Jigglypuff
	PokemonName_Wigglytuff
	PokemonName_Zubat
	PokemonName_Golbat
	PokemonName_Oddish
	PokemonName_Gloom
	PokemonName_Vileplume
	PokemonName_Paras
	PokemonName_Parasect
	PokemonName_Venonat
	PokemonName_Venomoth
	PokemonName_Diglett
	PokemonName_Dugtrio
	PokemonName_Meowth
	PokemonName_Persian
	PokemonName_Psyduck
	PokemonName_Golduck
	PokemonName_Mankey
	PokemonName_Primeape
	PokemonName_Growlithe
	PokemonName_Arcanine
	PokemonName_Poliwag
	PokemonName_Poliwhirl
	PokemonName_Poliwrath
	PokemonName_Abra
	PokemonName_Kadabra
	PokemonName_Alakazam
	PokemonName_Machop
	PokemonName_Machoke
	PokemonName_Machamp
	PokemonName_Bellsprout
	PokemonName_Weepinbell
	PokemonName_Victreebel
	PokemonName_Tentacool
	PokemonName_Tentacruel
	PokemonName_Geodude
	PokemonName_Graveler
	PokemonName_Golem
	PokemonName_Ponyta
	PokemonName_Rapidash
	PokemonName_Slowpoke
	PokemonName_Slowbro
	PokemonName_Magnemite
	PokemonName_Magneton
	PokemonName_Farfetchd
	PokemonName_Doduo
	PokemonName_Dodrio
	PokemonName_Seel
	PokemonName_Dewgong
	PokemonName_Grimer
	PokemonName_Muk
	PokemonName_Shellder
	PokemonName_Cloyster
	PokemonName_Gastly
	PokemonName_Haunter
	PokemonName_Gengar
	PokemonName_Onix
	PokemonName_Drowzee
	PokemonName_Hypno
	PokemonName_Krabby
	PokemonName_Kingler
	PokemonName_Voltorb
	PokemonName_Electrode
	PokemonName_Exeggcute
	PokemonName_Exeggutor
	PokemonName_Cubone
	PokemonName_Marowak
	PokemonName_Hitmonlee
	PokemonName_Hitmonchan
	PokemonName_Lickitung
	PokemonName_Koffing
	PokemonName_Weezing
	PokemonName_Rhyhorn
	PokemonName_Rhydon
	PokemonName_Chansey
	PokemonName_Tangela
	PokemonName_Kangaskhan
	PokemonName_Horsea
	PokemonName_Seadra
	PokemonName_Goldeen
	PokemonName_Seaking
	PokemonName_Staryu
	PokemonName_Starmie
	PokemonName_MrMime
	PokemonName_Scyther
	PokemonName_Jynx
	PokemonName_Electabuzz
	PokemonName_Magmar
	PokemonName_Pinsir
	PokemonName_Tauros
	PokemonName_Magikarp
	PokemonName_Gyarados
	PokemonName_Lapras
	PokemonName_Ditto
	PokemonName_Eevee
	PokemonName_Vaporeon
	PokemonName_Jolteon
	PokemonName_Flareon
	PokemonName_Porygon
	PokemonName_Omanyte
	PokemonName_Omastar
	PokemonName_Kabuto
	PokemonName_Kabutops
	PokemonName_Aerodactyl
	PokemonName_Snorlax
	PokemonName_Articuno
	PokemonName_Zapdos
	PokemonName_Moltres
	PokemonName_Dratini
	PokemonName_Dragonair
	PokemonName_Dragonite
	PokemonName_Mewtwo
	PokemonName_Mew
	PokemonName_Chikorita
	PokemonName_Bayleef
	PokemonName_Meganium
	PokemonName_Cyndaquil
	PokemonName_Quilava
	PokemonName_Typhlosion
	PokemonName_Totodile
	PokemonName_Croconaw
	PokemonName_Feraligatr
	PokemonName_Sentret
	PokemonName_Furret
	PokemonName_Hoothoot
	PokemonName_Noctowl
	PokemonName_Ledyba
	PokemonName_Ledian
	PokemonName_Spinarak
	PokemonName_Ariados
	PokemonName_Crobat
	PokemonName_Chinchou
	PokemonName_Lanturn
	PokemonName_Pichu
	PokemonName_Cleffa
	PokemonName_Igglybuff
	PokemonName_Togepi
	PokemonName_Togetic
	PokemonName_Natu
	PokemonName_Xatu
	PokemonName_Mareep
	PokemonName_Flaaffy
	PokemonName_Ampharos
	PokemonName_Bellossom
	PokemonName_Marill
	PokemonName_Azumarill
	PokemonName_Sudowoodo
	PokemonName_Politoed
	PokemonName_Hoppip
	PokemonName_Skiploom
	PokemonName_Jumpluff
	PokemonName_Aipom
	PokemonName_Sunkern
	PokemonName_Sunflora
	PokemonName_Yanma
	PokemonName_Wooper
	PokemonName_Quagsire
	PokemonName_Espeon
	PokemonName_Umbreon
	PokemonName_Murkrow
	PokemonName_Slowking
	PokemonName_Misdreavus
	PokemonName_Unown
	PokemonName_Wobbuffet
	PokemonName_Girafarig
	PokemonName_Pineco
	PokemonName_Forretress
	PokemonName_Dunsparce
	PokemonName_Gligar
	PokemonName_Steelix
	PokemonName_Snubbull
	PokemonName_Granbull
	PokemonName_Qwilfish
	PokemonName_Scizor
	PokemonName_Shuckle
	PokemonName_Heracross
	PokemonName_Sneasel
	PokemonName_Teddiursa
	PokemonName_Ursaring
	PokemonName_Slugma
	PokemonName_Magcargo
	PokemonName_Swinub
	PokemonName_Piloswine
	PokemonName_Corsola
	PokemonName_Remoraid
	PokemonName_Octillery
	PokemonName_Delibird
	PokemonName_Mantine
	PokemonName_Skarmory
	PokemonName_Houndour
	PokemonName_Houndoom
	PokemonName_Kingdra
	PokemonName_Phanpy
	PokemonName_Donphan
	PokemonName_Porygon2
	PokemonName_Stantler
	PokemonName_Smeargle
	PokemonName_Tyrogue
	PokemonName_Hitmontop
	PokemonName_Smoochum
	PokemonName_Elekid
	PokemonName_Magby
	PokemonName_Miltank
	PokemonName_Blissey
	PokemonName_Raikou
	PokemonName_Entei
	PokemonName_Suicune
	PokemonName_Larvitar
	PokemonName_Pupitar
	PokemonName_Tyranitar
	PokemonName_Lugia
	PokemonName_Ho_Oh
	PokemonName_Celebi
	PokemonName_Treecko
	PokemonName_Grovyle
	PokemonName_Sceptile
	PokemonName_Torchic
	PokemonName_Combusken
	PokemonName_Blaziken
	PokemonName_Mudkip
	PokemonName_Marshtomp
	PokemonName_Swampert
	PokemonName_Poochyena
	PokemonName_Mightyena
	PokemonName_Zigzagoon
	PokemonName_Linoone
	PokemonName_Wurmple
	PokemonName_Silcoon
	PokemonName_Beautifly
	PokemonName_Cascoon
	PokemonName_Dustox
	PokemonName_Lotad
	PokemonName_Lombre
	PokemonName_Ludicolo
	PokemonName_Seedot
	PokemonName_Nuzleaf
	PokemonName_Shiftry
	PokemonName_Taillow
	PokemonName_Swellow
	PokemonName_Wingull
	PokemonName_Pelipper
	PokemonName_Ralts
	PokemonName_Kirlia
	PokemonName_Gardevoir
	PokemonName_Surskit
	PokemonName_Masquerain
	PokemonName_Shroomish
	PokemonName_Breloom
	PokemonName_Slakoth
	PokemonName_Vigoroth
	PokemonName_Slaking
	PokemonName_Nincada
	PokemonName_Ninjask
	PokemonName_Shedinja
	PokemonName_Whismur
	PokemonName_Loudred
	PokemonName_Exploud
	PokemonName_Makuhita
	PokemonName_Hariyama
	PokemonName_Azurill
	PokemonName_Nosepass
	PokemonName_Skitty
	PokemonName_Delcatty
	PokemonName_Sableye
	PokemonName_Mawile
	PokemonName_Aron
	PokemonName_Lairon
	PokemonName_Aggron
	PokemonName_Meditite
	PokemonName_Medicham
	PokemonName_Electrike
	PokemonName_Manectric
	PokemonName_Plusle
	PokemonName_Minun
	PokemonName_Volbeat
	PokemonName_Illumise
	PokemonName_Roselia
	PokemonName_Gulpin
	PokemonName_Swalot
	PokemonName_Carvanha
	PokemonName_Sharpedo
	PokemonName_Wailmer
	PokemonName_Wailord
	PokemonName_Numel
	PokemonName_Camerupt
	PokemonName_Torkoal
	PokemonName_Spoink
	PokemonName_Grumpig
	PokemonName_Spinda
	PokemonName_Trapinch
	PokemonName_Vibrava
	PokemonName_Flygon
	PokemonName_Cacnea
	PokemonName_Cacturne
	PokemonName_Swablu
	PokemonName_Altaria
	PokemonName_Zangoose
	PokemonName_Seviper
	PokemonName_Lunatone
	PokemonName_Solrock
	PokemonName_Barboach
	PokemonName_Whiscash
	PokemonName_Corphish
	PokemonName_Crawdaunt
	PokemonName_Baltoy
	PokemonName_Claydol
	PokemonName_Lileep
	PokemonName_Cradily
	PokemonName_Anorith
	PokemonName_Armaldo
	PokemonName_Feebas
	PokemonName_Milotic
	PokemonName_Castform
	PokemonName_Kecleon
	PokemonName_Shuppet
	PokemonName_Banette
	PokemonName_Duskull
	PokemonName_Dusclops
	PokemonName_Tropius
	PokemonName_Chimecho
	PokemonName_Absol
	PokemonName_Wynaut
	PokemonName_Snorunt
	PokemonName_Glalie
	PokemonName_Spheal
	PokemonName_Sealeo
	PokemonName_Walrein
	PokemonName_Clamperl
	PokemonName_Huntail
	PokemonName_Gorebyss
	PokemonName_Relicanth
	PokemonName_Luvdisc
	PokemonName_Bagon
	PokemonName_Shelgon
	PokemonName_Salamence
	PokemonName_Beldum
	PokemonName_Metang
	PokemonName_Metagross
	PokemonName_Regirock
	PokemonName_Regice
	PokemonName_Registeel
	PokemonName_Latias
	PokemonName_Latios
	PokemonName_Kyogre
	PokemonName_Groudon
	PokemonName_Rayquaza
	PokemonName_Jirachi
	PokemonName_Deoxys
	PokemonName_Turtwig
	PokemonName_Grotle
	PokemonName_Torterra
	PokemonName_Chimchar
	PokemonName_Monferno
	PokemonName_Infernape
	PokemonName_Piplup
	PokemonName_Prinplup
	PokemonName_Empoleon
	PokemonName_Starly
	PokemonName_Staravia
	PokemonName_Staraptor
	PokemonName_Bidoof
	PokemonName_Bibarel
	PokemonName_Kricketot
	PokemonName_Kricketune
	PokemonName_Shinx
	PokemonName_Luxio
	PokemonName_Luxray
	PokemonName_Budew
	PokemonName_Roserade
	PokemonName_Cranidos
	PokemonName_Rampardos
	PokemonName_Shieldon
	PokemonName_Bastiodon
	PokemonName_Burmy
	PokemonName_Wormadam
	PokemonName_Mothim
	PokemonName_Combee
	PokemonName_Vespiquen
	PokemonName_Pachirisu
	PokemonName_Buizel
	PokemonName_Floatzel
	PokemonName_Cherubi
	PokemonName_Cherrim
	PokemonName_Shellos
	PokemonName_Gastrodon
	PokemonName_Ambipom
	PokemonName_Drifloon
	PokemonName_Drifblim
	PokemonName_Buneary
	PokemonName_Lopunny
	PokemonName_Mismagius
	PokemonName_Honchkrow
	PokemonName_Glameow
	PokemonName_Purugly
	PokemonName_Chingling
	PokemonName_Stunky
	PokemonName_Skuntank
	PokemonName_Bronzor
	PokemonName_Bronzong
	PokemonName_Bonsly
	PokemonName_MimeJr
	PokemonName_Happiny
	PokemonName_Chatot
	PokemonName_Spiritomb
	PokemonName_Gible
	PokemonName_Gabite
	PokemonName_Garchomp
	PokemonName_Munchlax
	PokemonName_Riolu
	PokemonName_Lucario
	PokemonName_Hippopotas
	PokemonName_Hippowdon
	PokemonName_Skorupi
	PokemonName_Drapion
	PokemonName_Croagunk
	PokemonName_Toxicroak
	PokemonName_Carnivine
	PokemonName_Finneon
	PokemonName_Lumineon
	PokemonName_Mantyke
	PokemonName_Snover
	PokemonName_Abomasnow
	PokemonName_Weavile
	PokemonName_Magnezone
	PokemonName_Lickilicky
	PokemonName_Rhyperior
	PokemonName_Tangrowth
	PokemonName_Electivire
	PokemonName_Magmortar
	PokemonName_Togekiss
	PokemonName_Yanmega
	PokemonName_Leafeon
	PokemonName_Glaceon
	PokemonName_Gliscor
	PokemonName_Mamoswine
	PokemonName_Porygon_Z
	PokemonName_Gallade
	PokemonName_Probopass
	PokemonName_Dusknoir
	PokemonName_Froslass
	PokemonName_Rotom
	PokemonName_Uxie
	PokemonName_Mesprit
	PokemonName_Azelf
	PokemonName_Dialga
	PokemonName_Palkia
	PokemonName_Heatran
	PokemonName_Regigigas
	PokemonName_Giratina
	PokemonName_Cresselia
	PokemonName_Phione
	PokemonName_Manaphy
	PokemonName_Darkrai
	PokemonName_Shaymin
	PokemonName_Arceus
	PokemonName_Victini
	PokemonName_Tsutarja
	PokemonName_Janobii
	PokemonName_Jaroda
	PokemonName_Pokabu
	PokemonName_Chaobuu
	PokemonName_Enbuoo
	PokemonName_Mijumaru
	PokemonName_Futachimaru
	PokemonName_Daikenki
	PokemonName_Minezumi
	PokemonName_Miruhoggu
	PokemonName_Yooterii
	PokemonName_Haderia
	PokemonName_Muurando
	PokemonName_Choroneko
	PokemonName_Leperasudu
	PokemonName_Yanappu
	PokemonName_Yanakki
	PokemonName_Boappu
	PokemonName_Baokki
	PokemonName_Hiyappu
	PokemonName_Hiyakki
	PokemonName_Munna
	PokemonName_Musharna
	PokemonName_Mamepato
	PokemonName_Hatoopoo
	PokemonName_Kenhorou
	PokemonName_Shimama
	PokemonName_Zeburaika
	PokemonName_Dangoro
	PokemonName_Gantoru
	PokemonName_Gigaiasu
	PokemonName_Koromori
	PokemonName_Kokoromori
	PokemonName_Moguryuu
	PokemonName_Doryuuzu
	PokemonName_Tabunne
	PokemonName_Dokkora
	PokemonName_Dotekkotsu
	PokemonName_Roopushin
	PokemonName_Otamaru
	PokemonName_Gamagaru
	PokemonName_Gamageroge
	PokemonName_Nageki
	PokemonName_Dageki
	PokemonName_Kurumiru
	PokemonName_Kurumayu
	PokemonName_Hahakurimo
	PokemonName_Futsude
	PokemonName_Hoiiga
	PokemonName_Pendoraa
	PokemonName_Monmon
	PokemonName_Erufuun
	PokemonName_Churine
	PokemonName_Doreida
	PokemonName_Basurao
	PokemonName_Meguroko
	PokemonName_Warubiru
	PokemonName_Warubiaru
	PokemonName_Darumakka
	PokemonName_Hihidaruma
	PokemonName_Marakacchi
	PokemonName_Ishizumai
	PokemonName_Iwaparesu
	PokemonName_Zuruggu
	PokemonName_Zuruzukin
	PokemonName_Shinpora
	PokemonName_Desumasu
	PokemonName_Desukan
	PokemonName_Purotooga
	PokemonName_Abagoora
	PokemonName_Aaken
	PokemonName_Aakeosu
	PokemonName_Yabakuron
	PokemonName_Dasutodasu
	PokemonName_Zorua
	PokemonName_Zoroark
	PokemonName_Chillarmy
	PokemonName_Chirachiino
	PokemonName_Gochimu
	PokemonName_Gochimiru
	PokemonName_Gochiruzeru
	PokemonName_Yuniran
	PokemonName_Daburan
	PokemonName_Rankurusu
	PokemonName_Koaruhii
	PokemonName_Swanna
	PokemonName_Banipucchi
	PokemonName_Baniricchi
	PokemonName_Baibanira
	PokemonName_Shikijika
	PokemonName_Mebukijka
	PokemonName_Emonga
	PokemonName_Kaburuchi
	PokemonName_Shubarugo
	PokemonName_Tamagetake
	PokemonName_Morobareru
	PokemonName_Pururiru
	PokemonName_Burunkeru
	PokemonName_Mamanbou
	PokemonName_Bachuru
	PokemonName_Denchura
	PokemonName_Tesshiido
	PokemonName_Nattorei
	PokemonName_Gear
	PokemonName_Gigear
	PokemonName_Gigigear
	PokemonName_Shibishirasu
	PokemonName_Shibibiiru
	PokemonName_Shibirudon
	PokemonName_Riguree
	PokemonName_Oobemu
	PokemonName_Hitomoshi
	PokemonName_Ranpuraa
	PokemonName_Shanderaa
	PokemonName_Kibago
	PokemonName_Onondo
	PokemonName_Ononokusu
	PokemonName_Kumashun
	PokemonName_Tsunbeaa
	PokemonName_Furiijio
	PokemonName_Chobomaki
	PokemonName_Agirudaa
	PokemonName_Maggyo
	PokemonName_Kojofuu
	PokemonName_Kojondo
	PokemonName_Kurimugan
	PokemonName_Gobitto
	PokemonName_Goruggo
	PokemonName_Komatana
	PokemonName_Kirikizan
	PokemonName_Baffuron
	PokemonName_Washibon
	PokemonName_Wargle
	PokemonName_Baruchai
	PokemonName_Barujiina
	PokemonName_Kuitaran
	PokemonName_Aianto
	PokemonName_Monozu
	PokemonName_Jiheddo
	PokemonName_Sazando
	PokemonName_Meraruba
	PokemonName_Urugamosu
	PokemonName_Kobaruon
	PokemonName_Terakion
	PokemonName_Birijion
	PokemonName_Torunerosu
	PokemonName_Borutorosu
	PokemonName_Reshiram
	PokemonName_Zekrom
	PokemonName_Randorosu
	PokemonName_Kyuremu
	PokemonName_Kerudio
	PokemonName_Meloia
	PokemonName_Insekuta
	// Base forms end here.
	PokemonName_Rotom_C = PokemonName_Rotom + (1 << 16)
	PokemonName_Rotom_H = PokemonName_Rotom + (2 << 16)
	PokemonName_Rotom_F = PokemonName_Rotom + (3 << 16)
	PokemonName_Rotom_W = PokemonName_Rotom + (4 << 16)
	PokemonName_Rotom_S = PokemonName_Rotom + (5 << 16)
	PokemonName_Deoxys_A = PokemonName_Deoxys + (1 << 16)
	PokemonName_Deoxys_D = PokemonName_Deoxys + (2 << 16)
	PokemonName_Deoxys_S = PokemonName_Deoxys + (3 << 16)
	PokemonName_Wormadam_G = PokemonName_Wormadam + (1 << 16)
	PokemonName_Wormadam_S = PokemonName_Wormadam + (2 << 16)
	PokemonName_Giratina_O = PokemonName_Giratina + (1 << 16)
	PokemonName_Shaymin_S = PokemonName_Shaymin + (1 << 16)
	PokemonName_Meloia_S = PokemonName_Meloia + (1 << 16)
)

const (
	PokemonStatus_Fine = 0
	PokemonStatus_Paralysed = 1
	PokemonStatus_Asleep = 2
	PokemonStatus_Frozen = 3
	PokemonStatus_Burnt = 4
	PokemonStatus_Poisoned = 5
	PokemonStatus_Confused = 6
	PokemonStatus_Attracted = 7
	PokemonStatus_Wrapped = 8
	PokemonStatus_Nightmared = 9
	PokemonStatus_Tormented = 12
	PokemonStatus_Disabled = 13
	PokemonStatus_Drowsy = 14
	PokemonStatus_HealBlocked = 15
	PokemonStatus_Sleuthed = 17
	PokemonStatus_Seeded = 18
	PokemonStatus_Embargoed = 19
	PokemonStatus_Requiemed = 20
	PokemonStatus_Rooted = 21
	PokemonStatus_Koed = 31
)

const ( 
	StatusFeeling_FeelConfusion int8 = iota
	StatusFeeling_HurtConfusion
	StatusFeeling_FreeConfusion
	StatusFeeling_PrevParalysed
	StatusFeeling_PrevFrozen
	StatusFeeling_FreeFrozen
	StatusFeeling_FeelAsleep
	StatusFeeling_FreeAsleep
	StatusFeeling_HurtBurn
	StatusFeeling_HurtPoison
)