package models

const (
	PokemonForms_Id             string = "pokemon_forms.id"
	PokemonForms_FormIdentifier string = "pokemon_forms.form_identifier"
	PokemonForms_PokemonId      string = "pokemon_forms.pokemon_id"
	PokemonForms_IsDefault      string = "pokemon_forms.is_default"
	PokemonForms_IsBattleOnly   string = "pokemon_forms.is_battle_only"
	PokemonForms_Order          string = "pokemon_forms.order"
)

type PokemonForms struct {
	Id             int `PK`
	FormIdentifier string
	PokemonId      int
	IsDefault      bool
	IsBattleOnly   bool
	//Order          int
}
