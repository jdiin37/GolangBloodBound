package main

import (
	"./gameAction"
	"./model"
	"fmt"
)

func main() {
	// fmt.Println("Hello, 世界")

	players := []model.Player{
		model.Player{
			ID:       1,
			NickName: "matt",
		}, model.Player{
			ID:       2,
			NickName: "wing",
		}, model.Player{
			ID:       3,
			NickName: "allen",
		},
		model.Player{
			ID:       4,
			NickName: "kevin",
		},
		model.Player{
			ID:       5,
			NickName: "corey",
		},
		model.Player{
			ID:       6,
			NickName: "linster",
		},
	}

	if game := gameAction.CreateGame(players); game != nil {
		fmt.Printf("遊戲開始\n")
		fmt.Printf("%+v\n", game.Game)

		currentPlayerIndex := 0
		//first Round
		for gameAction.CheckGameIsOver(game) == "" && game.Game.RoundCount < 20 {
			currentPlayerIndex = gameAction.RoundStart(game, currentPlayerIndex)
		}

		//fmt.Printf("所有玩家狀態 %+v\n", game.Game.AllPlayerStatus)

		for i,playerStatus := range game.Game.AllPlayerStatus{
			fmt.Printf("玩家 %v 狀態 \n %+v \n",i, playerStatus)
		}
		fmt.Printf("獲勝陣營 %+v\n", gameAction.CheckGameIsOver(game))

	}
	//gameAction.RandomCreateRole(6)

}
