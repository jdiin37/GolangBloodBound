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
			ID: 1,
			NickName:"matt",
		}, model.Player{
			ID: 2,
			NickName:"wing",
		},model.Player{
			ID: 3,
			NickName:"allen",
		},
		model.Player{
			ID: 4,
			NickName:"kevin",
		},
		model.Player{
			ID: 5,
			NickName:"corey",
		},
		model.Player{
			ID: 6,
			NickName:"linster",
		},
	}

	if game := gameAction.CreateGame(players);game != nil{	
		fmt.Printf("遊戲開始\n")
		fmt.Printf("%+v\n",*game.Game)
	}
	//gameAction.RandomCreateRole(6)

}
