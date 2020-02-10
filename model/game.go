package model

type Game struct {
	ID              int
	PlayerCount     int
	RoundCount      int            //所有回合數
	CurrentRound    int            //目前回合數
	AllPlayerStatus []PlayerStatus //所有玩家當前狀態
	AllRound        []Round        //所有回合LOG
	RedLeader       Role           //紅陣營老大
	BlueLeader      Role           //藍陣營老大
}

type Round struct {
	ID               int
	PlayerID         int //玩家ID
	StabPlayerID     int //刺殺玩家ID
	BlockingPlayerID int //擋刀玩家ID
}

type PlayerStatus struct {
	PlayerID         int    //玩家ID
	FirstBloodType   string //流第一滴血類型
	SeconedBloodType string //流第二滴血類型
	ShowRank         string //show軍階
	BeKill           bool   //是否死亡
	Role             Role   // 玩家所操縱的角色
	ActionCount      int    //行動回合次數
}
