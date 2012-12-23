package models

const (
	Pokemon_Id             string = "pokemon.id"
	Pokemon_SpeciesId      string = "pokemon.species_id"
	Pokemon_Height         string = "pokemon.height"
	Pokemon_Weight         string = "pokemon.weight"
	Pokemon_BaseExperience string = "pokemon.base_experience"
	Pokemon_Order          string = "pokemon.order"
	Pokemon_IsDefault      string = "pokemon.is_default"
)

type Pokemon struct {
	Id             int `PK`
	SpeciesId      int
	Height         int
	Weight         int
	BaseExperience int
	//Order          int
	IsDefault      bool
}
