package models

const (
	NpcPokemon_IdnpcPokemon  string = "npc_pokemon.idnpc_pokemon"
	NpcPokemon_Idpokemon     string = "npc_pokemon.idpokemon"
	NpcPokemon_Idnpc         string = "npc_pokemon.idnpc"
	NpcPokemon_IvHp          string = "npc_pokemon.iv_hp"
	NpcPokemon_IvAttack      string = "npc_pokemon.iv_attack"
	NpcPokemon_IvAttackSpec  string = "npc_pokemon.iv_attack_spec"
	NpcPokemon_IvDefence     string = "npc_pokemon.iv_defence"
	NpcPokemon_IvDefenceSpec string = "npc_pokemon.iv_defence_spec"
	NpcPokemon_IvSpeed       string = "npc_pokemon.iv_speed"
	NpcPokemon_Gender        string = "npc_pokemon.gender"
	NpcPokemon_HeldItem      string = "npc_pokemon.held_item"
)

type NpcPokemon struct {
	IdnpcPokemon  int `PK`
	Idpokemon     int
	Idnpc         int
	IvHp          int
	IvAttack      int
	IvAttackSpec  int
	IvDefence     int
	IvDefenceSpec int
	IvSpeed       int
	Gender        int
	HeldItem      int
}
