package models

const (
	Stats_Id            string = "stats.id"
	Stats_DamageClassId string = "stats.damage_class_id"
	Stats_Identifier    string = "stats.identifier"
	Stats_IsBattleOnly  string = "stats.is_battle_only"
)

type Stats struct {
	Id            int `PK`
	DamageClassId int
	Identifier    string
	IsBattleOnly  bool
}
