package models

const (
	Moves_Id                   string = "moves.id"
	Moves_Identifier           string = "moves.identifier"
	Moves_GenerationId         string = "moves.generation_id"
	Moves_TypeId               string = "moves.type_id"
	Moves_Power                string = "moves.power"
	Moves_Pp                   string = "moves.pp"
	Moves_Accuracy             string = "moves.accuracy"
	Moves_Priority             string = "moves.priority"
	Moves_TargetId             string = "moves.target_id"
	Moves_DamageClassId        string = "moves.damage_class_id"
	Moves_EffectId             string = "moves.effect_id"
	Moves_EffectChance         string = "moves.effect_chance"
	Moves_ContestTypeId        string = "moves.contest_type_id"
	Moves_ContestEffectId      string = "moves.contest_effect_id"
	Moves_SuperContestEffectId string = "moves.super_contest_effect_id"
)

type Moves struct {
	Id                   int `PK`
	Identifier           string
	GenerationId         int
	TypeId               int
	Power                int
	Pp                   int
	Accuracy             int
	Priority             int
	TargetId             int
	DamageClassId        int
	EffectId             int
	EffectChance         int
	ContestTypeId        int
	ContestEffectId      int
	SuperContestEffectId int
}

type MovesJoinMoveFlavorText struct {
	Id                   int `PK`
	Identifier           string
	GenerationId         int
	TypeId               int
	Power                int
	Pp                   int
	Accuracy             int
	Priority             int
	TargetId             int
	DamageClassId        int
	EffectId             int
	EffectChance         int
	ContestTypeId        int
	ContestEffectId      int
	SuperContestEffectId int
	
	// move_flavor
	IdMove         		int
	VersionGroupId 		int
	LanguageId     		int
	FlavorText     		string	
}