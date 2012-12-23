package models

const (
	MoveFlavorText_IdMove         string = "move_flavor_text.id_move"
	MoveFlavorText_VersionGroupId string = "move_flavor_text.version_group_id"
	MoveFlavorText_LanguageId     string = "move_flavor_text.language_id"
	MoveFlavorText_FlavorText     string = "move_flavor_text.flavor_text"
)

type MoveFlavorText struct {
	IdMove         int `PK`
	VersionGroupId int `PK`
	LanguageId     int `PK`
	FlavorText     string
}
