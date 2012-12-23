package models

const (
	PlayerPokemon_IdplayerPokemon string = "player_pokemon.idplayer_pokemon"
	PlayerPokemon_Idpokemon       string = "player_pokemon.idpokemon"
	PlayerPokemon_Idplayer        string = "player_pokemon.idplayer"
	PlayerPokemon_Nickname        string = "player_pokemon.nickname"
	PlayerPokemon_Bound           string = "player_pokemon.bound"
	PlayerPokemon_Experience      string = "player_pokemon.experience"
	PlayerPokemon_IvHp            string = "player_pokemon.iv_hp"
	PlayerPokemon_IvAttack        string = "player_pokemon.iv_attack"
	PlayerPokemon_IvAttackSpec    string = "player_pokemon.iv_attack_spec"
	PlayerPokemon_IvDefence       string = "player_pokemon.iv_defence"
	PlayerPokemon_IvDefenceSpec   string = "player_pokemon.iv_defence_spec"
	PlayerPokemon_IvSpeed         string = "player_pokemon.iv_speed"
	PlayerPokemon_Happiness       string = "player_pokemon.happiness"
	PlayerPokemon_Gender          string = "player_pokemon.gender"
	PlayerPokemon_InParty         string = "player_pokemon.in_party"
	PlayerPokemon_PartySlot       string = "player_pokemon.party_slot"
	PlayerPokemon_HeldItem        string = "player_pokemon.held_item"
	PlayerPokemon_Shiny           string = "player_pokemon.shiny"
	PlayerPokemon_Idability       string = "player_pokemon.idability"
	PlayerPokemon_DamagedHp       string = "player_pokemon.damaged_hp"
)

type PlayerPokemon struct {
	IdplayerPokemon int `PK`
	Idpokemon       int
	Idplayer        int
	Nickname        string
	Bound           bool
	Experience      int64
	IvHp            int
	IvAttack        int
	IvAttackSpec    int
	IvDefence       int
	IvDefenceSpec   int
	IvSpeed         int
	Happiness       int
	Gender          int
	InParty         bool
	PartySlot       int
	HeldItem        int
	Shiny           bool
	Idability       int
	DamagedHp       int
}
