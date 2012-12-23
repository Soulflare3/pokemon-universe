package models

const (
	MapchangeAccount_IdmapchangeAccount string = "mapchange_account.idmapchange_account"
	MapchangeAccount_Username           string = "mapchange_account.username"
	MapchangeAccount_Password           string = "mapchange_account.password"
)

type MapchangeAccount struct {
	IdmapchangeAccount int `PK`
	Username           string
	Password           string
}
