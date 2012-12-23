package models

const (
	LocationSection_IdlocationSection string = "location_section.idlocation_section"
	LocationSection_Idlocation        string = "location_section.idlocation"
	LocationSection_Name              string = "location_section.name"
)

type LocationSection struct {
	IdlocationSection int `PK`
	Idlocation        int
	Name              string
}
