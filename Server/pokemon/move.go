package main

type Move struct {
	MoveId				int
	Identifier			string
	TypeId				int
	Power				int
	Accuracy			int
	Priority			int
	TargetId			int
	DamageClassId		int
	EffectId			int
	EffectChance		int
	ContestType			int
	ContestEffect		int
	SuperContestEffect	int
}

func NewMove() *Move {
	return &Move{}
}