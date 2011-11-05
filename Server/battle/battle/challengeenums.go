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

const (
	CHALLENGEDESC_SENT int = iota
	CHALLENGEDESC_ACCEPTED
	CHALLENGEDESC_CANCELLED
	CHALLENGEDESC_BUSY
	CHALLENGEDESC_REFUSED
	CHALLENGEDESC_INVALIDTEAM
	CHALLENGEDESC_INVALIDGEN
	
	CHALLENGEDESC_CHALLENGEDESCLAST
)

const (
	BATTLERESULT_FORFEIT int = iota
	BATTLERESULT_WIN
	BATTLERESULT_TIE
	BATTLERESULT_CLOSE
)

const (
	MODE_SINGLES int = iota
	MODE_DOUBLES
	MODE_TRIPLES
	MODE_ROTATION
)

type Clauses interface {
	Mask() int
	String() string
	BattleText() string
}

type SleepClause struct {
}
func (SleepClause) Mask() int {
	return 1
}
func (SleepClause) String() string {
	return "Sleep Clause"
}
func (SleepClause) BattleText() string {
	return "Sleep Clause prevented the sleep inducing effect of the move from working."
}

type FreezeClause struct {
}
func (FreezeClause) Mask() int {
	return 2
}
func (FreezeClause) String() string {
	return "Freeze Clause"
}
func (FreezeClause) BattleText() string {
	return "Freeze Clause prevented the freezing effect of the move from working."
}

type DisallowSpectator struct {
}
func (DisallowSpectator) Mask() int {
	return 4
}
func (DisallowSpectator) String() string {
	return "Disallow Spectators"
}
func (DisallowSpectator) BattleText() string {
	return ""
}

type ItemClause struct {
}
func (ItemClause) Mask() int {
	return 8
}
func (ItemClause) String() string {
	return "Item Clause"
}
func (ItemClause) BattleText() string {
	return ""
}

type ChallengeCup struct {
}
func (ChallengeCup) Mask() int {
	return 16
}
func (ChallengeCup) String() string {
	return "Challenge Cup"
}
func (ChallengeCup) BattleText() string {
	return ""
}

type NoTimeout struct {
}
func (NoTimeout) Mask() int {
	return 32
}
func (NoTimeout) String() string {
	return "No Timeout"
}
func (NoTimeout) BattleText() string {
	return "The battle ended by timeout."
}

type SpeciesClause struct {
}
func (SpeciesClause) Mask() int {
	return 64
}
func (SpeciesClause) String() string {
	return "Species Clause"
}
func (SpeciesClause) BattleText() string {
	return ""
}

type RearrangeTeams struct {
}
func (RearrangeTeams) Mask() int {
	return 128
}
func (RearrangeTeams) String() string {
	return "Wifi Battle"
}
func (RearrangeTeams) BattleText() string {
	return ""
}

type SelfKO struct {
}
func (SelfKO) Mask() int {
	return 256
}
func (SelfKO) String() string {
	return "Self-KO Clause"
}
func (SelfKO) BattleText() string {
	return "The Self-KO Clause acted as a tiebreaker."
}