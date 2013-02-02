package models

const (
	Map_Idmap string = "map.idmap"
	Map_Name  string = "map.name"
)

type Map struct {
	Idmap int `PK`
	Name  string
}