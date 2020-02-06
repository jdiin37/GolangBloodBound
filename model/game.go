package model

type Game struct {
	ID           int
	PlayerCount  int
	RoundCount   int
	CurrentRound int
	AllPlayerStatus []PlayerStatus
}

type Round struct{
	ID             int
	PlayerID         int
	StabPlayerID     int
	BlockingPlayerID int
}

type PlayerStatus struct{
	PlayerID int
	FirstBloodType string
	SeconedBloodType string	
	ShowRank string
	Role Role
}