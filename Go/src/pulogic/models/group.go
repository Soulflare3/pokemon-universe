package models

const (
	Group_Idgroup  string = "group.idgroup"
	Group_Name     string = "group.name"
	Group_Flags    string = "group.flags"
	Group_Priority string = "group.priority"
)

type Group struct {
	Idgroup  int `PK`
	Name     string
	Flags    int
	Priority int
}
