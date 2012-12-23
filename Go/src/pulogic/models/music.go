package models

const (
	Music_Idmusic  string = "music.idmusic"
	Music_Title    string = "music.title"
	Music_Filename string = "music.filename"
)

type Music struct {
	Idmusic  int `PK`
	Title    string
	Filename string
}
