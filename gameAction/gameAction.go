package gameAction

import (
	"../model"
	"fmt"
	"math/rand"
	"time"
)

type action interface {
	fillPlayer()
	roundStart()
	roundEnd()
}

type TheGame struct {
	Game *model.Game
}

func checkGameIsOver() {

}

func RandomCreateRole(playerCnt int) []model.Role {
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

func CreateGame(allPlayers []model.Player) *TheGame {
	if len(allPlayers) != 6 && len(allPlayers) != 8 {
		fmt.Println("人數不足必須為6或8人", len(allPlayers))
		return nil
	}

	game := &model.Game{}
	game.PlayerCount = len(allPlayers)

	randomAllRole := RandomCreateRole(game.PlayerCount)

	allPlayerStatus := []model.PlayerStatus{}
	for i, player := range allPlayers {
		playerStatus := model.PlayerStatus{
			PlayerID:         player.ID,
			FirstBloodType:   "",
			SeconedBloodType: "",
			ShowRank:         "",
			Role:randomAllRole[i],
		}
		allPlayerStatus = append(allPlayerStatus, playerStatus)
	}

	game.AllPlayerStatus = allPlayerStatus

	return &TheGame{
		Game: game,
	}
}

func (theGame *TheGame) fillPlayer(playerCount int) {
	theGame.Game.PlayerCount = playerCount

}
