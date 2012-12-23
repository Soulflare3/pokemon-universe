package models

const (
	Mapchange_Idmapchange string = "mapchange.idmapchange"
	Mapchange_StartX      string = "mapchange.start_x"
	Mapchange_StartY      string = "mapchange.start_y"
	Mapchange_Width       string = "mapchange.width"
	Mapchange_Height      string = "mapchange.height"
	Mapchange_Username    string = "mapchange.username"
	Mapchange_Description string = "mapchange.description"
	Mapchange_SubmitDate  string = "mapchange.submit_date"
	Mapchange_Status      string = "mapchange.status"
)

type Mapchange struct {
	Idmapchange int `PK`
	StartX      int
	StartY      int
	Width       int
	Height      int
	Username    string
	Description string
	SubmitDate  string
	Status      int
}
