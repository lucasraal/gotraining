package main

import "fmt"

type team struct {
	team_id		int
	team_name	string
	power_def	int
	power_mid	int
	power_atk	int
	owner_id	string
}

func main() {
	t1 := team{
		team_id:	01,
		team_name:	"Flamengo",
		power_def:	33,
		power_mid:	72,
		power_atk:	85,
		owner_id:	"0xHKL0G",
	}
	fmt.Println(t1)
	fmt.Println(t1.team_name, t1.power_atk)
}

