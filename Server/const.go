/*Pokemon Universe MMORPG
Copyright (C) 2010 the Pokemon Universe Authors

This program is free software; you can redistribute it and/or
modify it under the terms of the GNU General Public License
as published by the Free Software Foundation; either version 2
of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program; if not, write to the Free Software
Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA  02110-1301, USA.*/
package main

import "time"

type ReturnValue int

const (
	RET_NOERROR ReturnValue = iota
	RET_NOTPOSSIBLE
	RET_PLAYERISTELEPORTED
)

const NANOSECONDS_TO_MILLISECONDS = 0.000001

func PUSYS_TIME() int64 {
	timeNano := float64(time.Nanoseconds())
	return int64(timeNano * NANOSECONDS_TO_MILLISECONDS)
}

func GetTypeValueById(_id int) (toReturn string) {
	toReturn = "Unknown"
	switch _id {
		case TYPE_NORMAL:
			toReturn = "Normal"
		case TYPE_FIGHTING:
			toReturn = "Fighting"
			case TYPE_FLYING:
			toReturn = "Flying"
		case TYPE_POISON:
			toReturn = "Poison"
		case TYPE_GROUND:
			toReturn = "Ground"
		case TYPE_ROCK:
			toReturn = "Rock"
		case TYPE_BUG:
			toReturn = "Bug"
		case TYPE_GHOST:
			toReturn = "Ghost"
		case TYPE_STEEL:
			toReturn = "Steel"
		case TYPE_FIRE:
			toReturn = "Fire"
		case TYPE_WATER:
			toReturn = "Water"
		case TYPE_GRASS:
			toReturn = "Grass"
		case TYPE_ELECTRIC:
			toReturn = "Electric"
		case TYPE_PSYCHIC:
			toReturn = "Psychic"
		case TYPE_ICE:
			toReturn = "Ice"
		case TYPE_DRAGON:
			toReturn = "Dragon"
		case TYPE_DARK:
			toReturn = "Dark"
	}
	
	return
}

func GetStatById(_id int) (toReturn string) {
	toReturn = ""
	switch _id {
		case 1:
			toReturn = STAT_ATTACK
		case 2:
			toReturn = STAT_DEFENSE
		case 3:
			toReturn = STAT_SPATTACK
		case 4:
			toReturn = STAT_SPDEFENSE
		case 5:
			toReturn = STAT_SPEED
		case 6:
			toReturn = STAT_ACCURACY
		case 7:
			toReturn = STAT_EVASION
	}
	return toReturn
}

func GetStatusById(_id int) (toReturn string) {
	toReturn = ""
	switch _id {
		case STATUS_FINE:
			toReturn = "fine"
		case STATUS_PARALYSED:
			toReturn = "paralysed"
		case STATUS_ASLEEP:
			toReturn = "asleep"
		case STATUS_FROZEN:
			toReturn = "frozen"			
		case STATUS_BURNT:
			toReturn = "burnt"
		case STATUS_POISONED:
			toReturn = "poisoned"
		case STATUS_CONFUSED:
			toReturn = "confused"
		case STATUS_ATTRACTED:
			toReturn = "attracted"
		case STATUS_WRAPPED:
			toReturn = "wrapped"
		case STATUS_NIGHTMARED:
			toReturn = "nightmared"
		case STATUS_TORMENTED:
			toReturn = "tormented"
		case STATUS_DISABLED:
			toReturn = "disabled"
		case STATUS_DROWSY:
			toReturn = "drowsy"
		case STATUS_HEALBLOCKED:
			toReturn = "healblocked"
		case STATUS_SLEUTHED:
			toReturn = "sleuthed"
		case STATUS_SEEDED:
			toReturn = "seeded"
		case STATUS_EMBARGOED:
			toReturn = "embargoed"
		case STATUS_REQUIEMED:
			toReturn = "requimed"
		case STATUS_ROOTED:
			toReturn = "rooted"
		case STATUS_KEOD:
			toReturn = "keod"
	}
	return
}