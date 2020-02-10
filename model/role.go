package model

type Role struct {
	Name        string
	Color       string
	Rank        int
	FirstBlood  string
	SecondBlood string
	RightColor  string
}

type RoleList struct {
	RedRole  []Role
	BlueRole []Role
}

func AllRoleDeck() RoleList {
	roleList := RoleList{}
	redRoleDeck := []Role{
		Role{
			Name:        "R1",
			Color:       "red",
			Rank:        1,
			FirstBlood:  "red",
			SecondBlood: "red",
		},
		Role{
			Name:        "R2",
			Color:       "red",
			Rank:        2,
			FirstBlood:  "?",
			SecondBlood: "?",
		},
		Role{
			Name:        "R3",
			Color:       "red",
			Rank:        3,
			FirstBlood:  "?",
			SecondBlood: "?",
		},
		Role{
			Name:        "R4",
			Color:       "red",
			Rank:        4,
			FirstBlood:  "?",
			SecondBlood: "?",
		},
		Role{
			Name:        "R5",
			Color:       "red",
			Rank:        5,
			FirstBlood:  "red",
			SecondBlood: "red",
		},
		Role{
			Name:        "R6",
			Color:       "red",
			Rank:        6,
			FirstBlood:  "red",
			SecondBlood: "red",
		},
		Role{
			Name:        "R7",
			Color:       "red",
			Rank:        7,
			FirstBlood:  "red",
			SecondBlood: "?",
		},
		Role{
			Name:        "R8",
			Color:       "red",
			Rank:        8,
			FirstBlood:  "red",
			SecondBlood: "?",
		},
		Role{
			Name:        "R9",
			Color:       "red",
			Rank:        9,
			FirstBlood:  "red",
			SecondBlood: "?",
		},
	}

	blueRoleDeck := []Role{
		Role{
			Name:        "B1",
			Color:       "blue",
			Rank:        1,
			FirstBlood:  "blue",
			SecondBlood: "blue",
		},
		Role{
			Name:        "B2",
			Color:       "blue",
			Rank:        2,
			FirstBlood:  "?",
			SecondBlood: "?",
		},
		Role{
			Name:        "B3",
			Color:       "blue",
			Rank:        3,
			FirstBlood:  "?",
			SecondBlood: "?",
		},
		Role{
			Name:        "B4",
			Color:       "blue",
			Rank:        4,
			FirstBlood:  "?",
			SecondBlood: "?",
		},
		Role{
			Name:        "B5",
			Color:       "blue",
			Rank:        5,
			FirstBlood:  "blue",
			SecondBlood: "blue",
		},
		Role{
			Name:        "B6",
			Color:       "blue",
			Rank:        6,
			FirstBlood:  "blue",
			SecondBlood: "blue",
		},
		Role{
			Name:        "B7",
			Color:       "blue",
			Rank:        7,
			FirstBlood:  "blue",
			SecondBlood: "?",
		},
		Role{
			Name:        "B8",
			Color:       "blue",
			Rank:        8,
			FirstBlood:  "blue",
			SecondBlood: "?",
		},
		Role{
			Name:        "B9",
			Color:       "blue",
			Rank:        9,
			FirstBlood:  "blue",
			SecondBlood: "?",
		},
	}

	roleList.RedRole = redRoleDeck
	roleList.BlueRole = blueRoleDeck
	return roleList

}
