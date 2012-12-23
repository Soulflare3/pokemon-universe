package models

const (
	Teleport_Idteleport string = "teleport.idteleport"
	Teleport_X          string = "teleport.x"
	Teleport_Y          string = "teleport.y"
	Teleport_Z          string = "teleport.z"
)

type Teleport struct {
	Idteleport int `PK`
	X          int
	Y          int
	Z          int
}
