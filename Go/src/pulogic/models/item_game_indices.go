package models

const (
	ItemGameIndices_ItemId       string = "item_game_indices.item_id"
	ItemGameIndices_GenerationId string = "item_game_indices.generation_id"
	ItemGameIndices_GameIndex    string = "item_game_indices.game_index"
)

type ItemGameIndices struct {
	ItemId       int `PK`
	GenerationId int `PK`
	GameIndex    int
}
