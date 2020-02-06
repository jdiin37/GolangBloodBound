package model

type Role struct {
	Name        string
	Color       string
	RolePoint   int
	FirstBlood  string
	SecondBlood string
	RightColor string
}

type RoleList struct {
	RedRole  []Role
	BlueRole []Role
}

func AllRoleDeck() RoleList {
	roleList := RoleList{}
	redRoleDeck := []Role{
		Role{
			Name:"R1",
			Color:"red",
			RolePoint:1,
			FirstBlood:"red",
			SecondBlood:"red",			
		},
		Role{
			Name:"R2",
			Color:"red",
			RolePoint:2,
			FirstBlood:"?",
			SecondBlood:"?",			
		},
		Role{
			Name:"R3",
			Color:"red",
			RolePoint:3,
			FirstBlood:"?",
			SecondBlood:"?",			
		},
		Role{
			Name:"R4",
			Color:"red",
			RolePoint:4,
			FirstBlood:"?",
			SecondBlood:"?",			
		},
		Role{
			Name:"R5",
			Color:"red",
			RolePoint:5,
			FirstBlood:"red",
			SecondBlood:"red",			
		},
		Role{
			Name:"R6",
			Color:"red",
			RolePoint:6,
			FirstBlood:"red",
			SecondBlood:"red",			
		},
		Role{
			Name:"R7",
			Color:"red",
			RolePoint:7,
			FirstBlood:"red",
			SecondBlood:"?",			
		},
		Role{
			Name:"R8",
			Color:"red",
			RolePoint:8,
			FirstBlood:"red",
			SecondBlood:"?",			
		},
		Role{
			Name:"R9",
			Color:"red",
			RolePoint:9,
			FirstBlood:"red",
			SecondBlood:"?",			
		},
	}

	blueRoleDeck := []Role{
		Role{
			Name:"B1",
			Color:"blue",
			RolePoint:1,
			FirstBlood:"blue",
			SecondBlood:"blue",			
		},
		Role{
			Name:"B2",
			Color:"blue",
			RolePoint:2,
			FirstBlood:"?",
			SecondBlood:"?",			
		},
		Role{
			Name:"B3",
			Color:"blue",
			RolePoint:3,
			FirstBlood:"?",
			SecondBlood:"?",			
		},
		Role{
			Name:"B4",
			Color:"blue",
			RolePoint:4,
			FirstBlood:"?",
			SecondBlood:"?",			
		},
		Role{
			Name:"B5",
			Color:"blue",
			RolePoint:5,
			FirstBlood:"blue",
			SecondBlood:"blue",			
		},
		Role{
			Name:"B6",
			Color:"blue",
			RolePoint:6,
			FirstBlood:"blue",
			SecondBlood:"blue",			
		},
		Role{
			Name:"B7",
			Color:"blue",
			RolePoint:7,
			FirstBlood:"blue",
			SecondBlood:"?",			
		},
		Role{
			Name:"B8",
			Color:"blue",
			RolePoint:8,
			FirstBlood:"blue",
			SecondBlood:"?",			
		},
		Role{
			Name:"B9",
			Color:"blue",
			RolePoint:9,
			FirstBlood:"blue",
			SecondBlood:"?",			
		},
	}

	roleList.RedRole = redRoleDeck
	roleList.BlueRole = blueRoleDeck
	return roleList

}