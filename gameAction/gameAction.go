package gameAction

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"../model"
)

type action interface {
	fillPlayer()
	roundStart()
	roundEnd()
}

type TheGame struct {
	Game *model.Game
}

func CheckGameIsOver(theGame *TheGame) string {
	for _, playerStatus := range theGame.Game.AllPlayerStatus {
		if playerStatus.BeKill {
			if playerStatus.Role == theGame.Game.RedLeader { //殺到老大 自己陣營獲勝
				return "blue"
			} else if playerStatus.Role == theGame.Game.BlueLeader { //殺到老大 自己陣營獲勝
				return "red"
			} else if playerStatus.Role.Color == "red" { //沒殺到老大 對方陣營獲勝
				return "red"
			} else if playerStatus.Role.Color == "blue" { //沒殺到老大 對方陣營獲勝
				return "blue"
			}
		}
	}
	return ""
}

func randomCreateRole(playerCnt int) []model.Role {
	rand.Seed(time.Now().UnixNano())
	allRoleDeck := model.AllRoleDeck()
	redDeck := allRoleDeck.RedRole
	blueDeck := allRoleDeck.BlueRole

	allRole := []model.Role{}
	randomAllRole := []model.Role{}

	switch playerCnt {
	case 6:
		//red
		for len(allRole) < 3 {
			redDeckIndex := rand.Intn(len(redDeck))
			allRole = append(allRole, redDeck[redDeckIndex])
			redDeck = append(redDeck[:redDeckIndex], redDeck[(redDeckIndex+1):]...)
		}
		//blue
		for len(allRole) < 6 {
			redDeckIndex := rand.Intn(len(blueDeck))
			allRole = append(allRole, blueDeck[redDeckIndex])
			blueDeck = append(blueDeck[:redDeckIndex], blueDeck[(redDeckIndex+1):]...)
		}
	case 8:
		//red
		for len(allRole) < 4 {
			redDeckIndex := rand.Intn(len(redDeck))
			allRole = append(allRole, redDeck[redDeckIndex])
			redDeck = append(redDeck[:redDeckIndex], redDeck[(redDeckIndex+1):]...)
		}
		//blue
		for len(allRole) < 8 {
			redDeckIndex := rand.Intn(len(blueDeck))
			allRole = append(allRole, blueDeck[redDeckIndex])
			blueDeck = append(blueDeck[:redDeckIndex], blueDeck[(redDeckIndex+1):]...)
		}
	default:
	}
	fmt.Println("beforeRandom", allRole)

	//start 隨機抽卡
	for len(allRole) > 0 {
		redDeckIndex := rand.Intn(len(allRole))
		randomAllRole = append(randomAllRole, allRole[redDeckIndex])
		allRole = append(allRole[:redDeckIndex], allRole[(redDeckIndex+1):]...)
	}
	//set rightRole Color
	for i := 0; i < len(randomAllRole); i++ {
		if i+1 == len(randomAllRole) {
			randomAllRole[i].RightColor = randomAllRole[0].Color
		} else {
			randomAllRole[i].RightColor = randomAllRole[i+1].Color
		}
	}
	fmt.Println("afterRandom", randomAllRole)

	return randomAllRole
}

//CreateGame 建立一場遊戲
func CreateGame(allPlayers []model.Player) *TheGame {
	if len(allPlayers) != 6 && len(allPlayers) != 8 {
		fmt.Println("人數不足必須為6或8人", len(allPlayers))
		return nil
	}

	game := &model.Game{}
	game.PlayerCount = len(allPlayers)
	game.RedLeader = model.Role{
		Rank: 9,
	}
	game.BlueLeader = model.Role{
		Rank: 9,
	}

	randomAllRole := randomCreateRole(game.PlayerCount)
	for _, role := range randomAllRole {
		if role.Color == "red" {
			if game.RedLeader.Rank > role.Rank { //rank最小的才是老大
				game.RedLeader = role
			}
		} else if role.Color == "blue" {
			if game.BlueLeader.Rank > role.Rank { //rank最小的才是老大
				game.BlueLeader = role
			}
		}
	}

	allPlayerStatus := []model.PlayerStatus{}
	for i, player := range allPlayers {
		playerStatus := model.PlayerStatus{
			PlayerID:         player.ID,
			FirstBloodType:   "",
			SeconedBloodType: "",
			ShowRank:         "",
			Role:             randomAllRole[i],
		}
		allPlayerStatus = append(allPlayerStatus, playerStatus)
	}

	game.AllPlayerStatus = allPlayerStatus

	return &TheGame{
		Game: game,
	}
}

func RoundStart(theGame *TheGame, playerIndex int) int {
	theGame.Game.RoundCount += 1
	theGame.Game.AllPlayerStatus[playerIndex].ActionCount += 1
	stabPlayerIndex := rand.Intn(theGame.Game.PlayerCount)

	Stab(theGame, stabPlayerIndex)
	round := model.Round{
		ID:               0,
		PlayerID:         theGame.Game.AllPlayerStatus[playerIndex].PlayerID,
		StabPlayerID:     theGame.Game.AllPlayerStatus[stabPlayerIndex].PlayerID,
		BlockingPlayerID: 0,
	}

	theGame.Game.AllRound = append(theGame.Game.AllRound, round)

	fmt.Printf("Round %v  %+v\n", theGame.Game.RoundCount, theGame.Game.AllRound)

	return stabPlayerIndex
}

func Stab(theGame *TheGame, beStabPlayerIndex int) {
	if theGame.Game.AllPlayerStatus[beStabPlayerIndex].FirstBloodType == "" {
		theGame.Game.AllPlayerStatus[beStabPlayerIndex].FirstBloodType = theGame.Game.AllPlayerStatus[beStabPlayerIndex].Role.FirstBlood
	} else if theGame.Game.AllPlayerStatus[beStabPlayerIndex].SeconedBloodType == "" {
		theGame.Game.AllPlayerStatus[beStabPlayerIndex].SeconedBloodType = theGame.Game.AllPlayerStatus[beStabPlayerIndex].Role.SecondBlood
	} else if theGame.Game.AllPlayerStatus[beStabPlayerIndex].ShowRank == "" {
		theGame.Game.AllPlayerStatus[beStabPlayerIndex].ShowRank = strconv.Itoa(theGame.Game.AllPlayerStatus[beStabPlayerIndex].Role.Rank)
	} else {
		theGame.Game.AllPlayerStatus[beStabPlayerIndex].BeKill = true
	}

}
