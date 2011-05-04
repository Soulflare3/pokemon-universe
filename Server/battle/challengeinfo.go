package main

const (
	ChallengeDesc_Sent uint8 = iota
	ChallengeDesc_Accepted
	ChallengeDesc_Cancelled
	ChallengeDesc_Busy
	ChallengeDesc_Refused
	ChallengeDesc_InvalidTeam
	ChallengeDesc_InvalidGen

	ChallengeDescLast
)
	
type ChallengeInfo struct {
	description	uint8
	opponent 	uint32
	clauses		uint32
	mode		uint8
}

func NewChallengeInfo() *ChallengeInfo {
	return &ChallengeInfo{ }
}