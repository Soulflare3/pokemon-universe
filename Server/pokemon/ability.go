package main

type Ability struct {
	AbilityId		int
	Identifier		string
}

func NewAbility() *Ability {
	return &Ability{}
}