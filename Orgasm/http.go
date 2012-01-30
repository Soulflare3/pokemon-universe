package main

/* 
import (
	"crypto/sha1"
	"fmt"
	"hash"
	"net/http"

	"mysql"
	"strings"
)

var loadedImages map[string]*TileCollection = make(map[string]*TileCollection)

func addHandler(w http.ResponseWriter, r *http.Request) {
	g_dblock.Lock()
	defer g_dblock.Unlock()

	r.ParseForm()
	username := r.FormValue("username")
	password := r.FormValue("password")
	query := fmt.Sprintf("INSERT INTO mapchange_account (username, password) VALUES ('%s', '%s')", username, SHA1(password))
	if err := g_db.Query(query); err != nil {
		fmt.Printf("Database query error: %s\n", err.Error())
		return
	}
	http.Redirect(w, r, "/users/", http.StatusFound)
}

func addFormHandler(_w http.ResponseWriter, _r *http.Request) {
	fmt.Fprintf(_w, adduserpage)
}

func removeUser(_username string, _w http.ResponseWriter, _r *http.Request) {
	g_dblock.Lock()
	defer g_dblock.Unlock()

	query := fmt.Sprintf("DELETE FROM mapchange_account WHERE username = '%s'", _username)
	if err := g_db.Query(query); err != nil {
		fmt.Printf("Database query error: %s\n", err.Error())
		return
	}
	http.Redirect(_w, _r, "/users/", http.StatusFound)
}

func SHA1(_plain string) string {
	var h hash.Hash = sha1.New()
	h.Write([]byte(_plain))

	return fmt.Sprintf("%x", h.Sum(nil))
}

func SourceHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, r.URL.Path[1:])
}

func imageData(_w http.ResponseWriter, _r *http.Request) {
	path := strings.Trim(_r.URL.Path, "/")
	parts := strings.Split(path, "/")
	image := parts[len(parts)-1]

	http.ServeFile(_w, _r, "images/"+image+".png")
}

func showPreview(_id string, _w http.ResponseWriter) {
	if _, contains := loadedImages[_id]; !contains {
		if !createImage(_id) {
			fmt.Fprintf(_w, "Preview not found")
			return
		}
	}

	collection := loadedImages[_id]
	createDatabaseImage(_id)
	previewPage := strings.Replace(page, "PAARD1", "/image/"+_id, -1)
	previewPage = strings.Replace(previewPage, "PAARD2", "/image/"+_id+"_db", -1)
	previewPage = strings.Replace(previewPage, "PAARD3", "/image/"+_id+"_bl", -1)
	previewPage = strings.Replace(previewPage, "USERNAME", collection.username, -1)
	previewPage = strings.Replace(previewPage, "COMMITLOG", collection.description, -1)
	previewPage = strings.Replace(previewPage, "CHANGEID", _id, -1)
	fmt.Fprintf(_w, previewPage)
}

func getDbTiles(_id string) *TileCollection {
	collection := loadedImages[_id]
	if collection != nil {
		dbcollection := &TileCollection{}
		dbcollection.x = collection.x
		dbcollection.y = collection.y
		dbcollection.width = collection.width
		dbcollection.height = collection.height
		for _, t := range collection.Tiles {
			tile, found := g_map.GetTile(t.Position.Hash())
			if found == true {
				dbtile := NewTile(tile.Position)

				for i, layer := range tile.layers {
					if layer != nil {
						dbtile.AddLayer(i, layer.id)
					}
				}
				dbcollection.AddTile(dbtile)
			}
		}
		return dbcollection
	}
	return nil
}

func createDatabaseImage(_id string) bool {
	var tiles *TileCollection
	if tiles = getDbTiles(_id); tiles == nil {
		return false
	}

	bg := NewImage(tiles.width*TILE_WIDTH, tiles.height*TILE_HEIGHT)
	for _, tile := range tiles.Tiles {
		drawX := (tile.x - tiles.x) * TILE_WIDTH
		drawY := (tile.y - tiles.y) * TILE_HEIGHT

		for _, layer := range tile.layers {
			if layer != nil {
				if tileimage, ok := ImagesMap[layer.id]; ok {
					tileimage.DrawOn(bg, drawX, drawY)
				}
			}
		}
	}
	url := fmt.Sprintf("images/%s_db.png", _id)
	bg.WriteToFile(url)
	loadedImages[_id+"_db"] = tiles
	http.HandleFunc("/image/"+_id+"_db", imageData)
	return true
}

func getTiles(_id string) *TileCollection {
	g_dblock.Lock()
	defer g_dblock.Unlock()

	var query string = "SELECT mapchange.idmapchange, start_x, start_y, width, height, username, description, mapchange_tile.idmapchange_tile, x, y, z, movement, `index`, sprite FROM mapchange "
	query += "INNER JOIN mapchange_tile ON mapchange_tile.idmapchange = mapchange.idmapchange "
	query += "INNER JOIN mapchange_layer ON mapchange_tile.idmapchange_tile = mapchange_layer.idmapchange_tile "
	query += fmt.Sprintf("WHERE mapchange.idmapchange = %s", _id)
	var err error
	if err = g_db.Query(query); err != nil {
		fmt.Printf("Query error: %s", err.Error())
		return nil
	}

	var result *mysql.Result
	result, err = g_db.UseResult()
	if err != nil {
		fmt.Printf("Query error: %s", err.Error())
		return nil
	}
	defer result.Free()

	collectionInited := false
	collection := &TileCollection{}

	currentTileID := 0
	var tile *Tile
	for {
		row := result.FetchMap()
		if row == nil {
			if tile != nil {
				collection.AddTile(tile)
			}
			break
		}

		if !collectionInited {
			collection.username = row["username"].(string)
			if row["description"] != nil {
				collection.description = string(row["description"].([]uint8))
			} else {
				collection.description = ""
			}
			collection.x = row["start_x"].(int)
			collection.y = row["start_y"].(int)
			collection.width = row["width"].(int)
			collection.height = row["height"].(int)
			collectionInited = true
		}

		tileid := row["idmapchange_tile"].(int)
		if tileid != currentTileID {
			if tile != nil {
				collection.AddTile(tile)
			}
			tileX := row["x"].(int)
			tileY := row["y"].(int)
			tileZ := row["z"].(int)
			tile = NewTile(tileX, tileY, tileZ)
			tile.blocking = row["movement"].(int)

			currentTileID = row["idmapchange_tile"].(int)
		}

		tile.AddLayer(row["index"].(int), row["sprite"].(int))
	}
	if len(collection.Tiles) == 0 {
		return nil
	}
	return collection
}

func createImage(_id string) bool {
	var tiles *TileCollection
	if tiles = getTiles(_id); tiles == nil {
		return false
	}
	blockbg := NewImage(tiles.width*TILE_WIDTH, tiles.height*TILE_HEIGHT)
	bg := NewImage(tiles.width*TILE_WIDTH, tiles.height*TILE_HEIGHT)
	for _, tile := range tiles.Tiles {
		if tile != nil {
			drawX := (tile.x - tiles.x) * TILE_WIDTH
			drawY := (tile.y - tiles.y) * TILE_HEIGHT

			blockimage := NewImageFile(fmt.Sprintf("blocking/%d.png", tile.blocking))
			if blockimage != nil {
				blockimage.DrawOn(blockbg, drawX, drawY)
			}
			for _, layer := range tile.layers {
				if layer != nil {
					tileimage := ImagesMap[layer.id]
					if tileimage != nil {
						tileimage.DrawOn(bg, drawX, drawY)
					}
				}
			}
		}
	}
	blockbg.WriteToFile(fmt.Sprintf("images/%s_bl.png", _id))
	http.HandleFunc("/image/"+_id+"_bl", imageData)

	url := fmt.Sprintf("images/%s.png", _id)
	bg.WriteToFile(url)
	loadedImages[_id] = tiles
	http.HandleFunc("/image/"+_id, imageData)
	return true
}

func userHandler(_w http.ResponseWriter, _r *http.Request) {
	g_dblock.Lock()
	defer g_dblock.Unlock()

	var query string = "SELECT * FROM mapchange_account"
	var err error
	if err = g_db.Query(query); err != nil {
		fmt.Printf("Query error: %s", err.Error())
		return
	}

	var result *mysql.Result
	result, err = g_db.UseResult()
	if err != nil {
		fmt.Printf("Query error: %s", err.Error())
		return
	}
	defer result.Free()
	fmt.Fprintf(_w, "<table border=\"1\">")
	fmt.Fprintf(_w, "<tr><td>Username</td><td> </td></tr>")
	for {
		row := result.FetchMap()
		if row == nil {
			break
		}
		if row["username"] == nil {
			continue
		}

		username := row["username"].(string)
		fmt.Fprintf(_w, "<tr>")
		fmt.Fprintf(_w, "<td>%s</td>", username)
		fmt.Fprintf(_w, "<td><a href=\"/removeuser/%s\">Remove</a></td>", username)
		fmt.Fprintf(_w, "</tr>")
	}
	fmt.Fprintf(_w, "</table>")
	fmt.Fprintf(_w, "<a href='/add/'>Add new</a>")
}

func listHandler(_w http.ResponseWriter, _r *http.Request) {
	g_dblock.Lock()
	defer g_dblock.Unlock()

	var query string = "SELECT * FROM mapchange WHERE status = 0"
	var err error
	if err = g_db.Query(query); err != nil {
		fmt.Printf("Query error: %s", err.Error())
		return
	}

	var result *mysql.Result
	result, err = g_db.UseResult()
	if err != nil {
		fmt.Printf("Query error: %s", err.Error())
		return
	}
	defer result.Free()
	fmt.Fprintf(_w, "<table border=\"1\">")
	fmt.Fprintf(_w, "<tr><td>Username</td><td>Date</td><td>Description</td><td> </td></tr>")
	for {
		row := result.FetchMap()
		if row == nil {
			break
		}
		fmt.Fprintf(_w, "<tr>")
		fmt.Fprintf(_w, "<td>%s</td>", row["username"].(string))
		fmt.Fprintf(_w, "<td>%s</td>", string(row["submit_date"].([]uint8)))
		if row["description"] != nil {
			fmt.Fprintf(_w, "<td>%s</td>", string(row["description"].([]uint8)))
		} else {
			fmt.Fprintf(_w, "<td></td>")
		}

		fmt.Fprintf(_w, "<td><a href=\"/preview/%d\">View</a></td>", row["idmapchange"].(int))
		fmt.Fprintf(_w, "</tr>")
	}
	fmt.Fprintf(_w, "</table>")
}

func handleIndex(_w http.ResponseWriter, _r *http.Request) {
	url := strings.Trim(_r.URL.RawPath, "/")
	values := strings.Split(url, "/")

	if len(values) == 2 {
		switch values[0] {
		case "preview":
			showPreview(values[1], _w)

		case "removeuser":
			removeUser(values[1], _w, _r)
		}
	} else if values[0] == "changehandler" {
		fmt.Println("Change Handler")
		submitType := _r.FormValue("handle")
		if submitType == "Accept" {
			fmt.Println("FORM accept")
			acceptPreview(_r.FormValue("change_id"), _r.FormValue("reason"), _w, _r)
		} else {
			fmt.Println("FORM declined")
			declinePreview(_r.FormValue("change_id"), _r.FormValue("reason"), _w, _r)
		}
	}
}

func acceptPreview(_id string, _message string, _w http.ResponseWriter, _r *http.Request) {
	g_dblock.Lock()
	defer g_dblock.Unlock()

	collection := loadedImages[_id]
	if collection != nil {
		for _, tile := range collection.Tiles {
			dbtile, found := g_map.GetTileFromCoordinates(tile.x, tile.y, tile.z)
			if found == true {
				query := fmt.Sprintf("DELETE FROM tile WHERE x = %d AND y = %d AND z = %d", tile.x, tile.y, tile.z)
				if err := g_db.Query(query); err != nil {
					fmt.Printf("Database query error: %s\n", err.Error())
					return
				}
				query = fmt.Sprintf("DELETE FROM tile_layer WHERE idtile = %d", tile.id)
				if err := g_db.Query(query); err != nil {
					fmt.Printf("Database query error: %s\n", err.Error())
					return
				}

				g_map.RemoveTile(dbtile)
			}

			query := fmt.Sprintf("INSERT INTO tile (x,y,z,movement,idmap,idlocation) VALUES (%d, %d, %d, %d, 1, 0)", tile.x, tile.y, tile.z, tile.blocking)
			if err := g_db.Query(query); err != nil {
				fmt.Printf("Database query error: %s\n", err.Error())
				return
			}

			g_map.Add(tile)

			tileid := g_db.LastInsertId

			for i, layer := range tile.layers {
				if layer != nil {
					query = fmt.Sprintf("INSERT INTO tile_layer (idtile, layer, sprite) VALUES (%d, %d, %d)", tileid, i, layer.id)
					if err := g_db.Query(query); err != nil {
						fmt.Printf("Database query error: %s\n", err.Error())
						return
					}
				}
			}
		}

		sendPrivateMessage(_id, _message, true)
		query := fmt.Sprintf("UPDATE mapchange SET status=1 WHERE idmapchange = %s", _id)
		if err := g_db.Query(query); err != nil {
			fmt.Printf("Database query error: %s\n", err.Error())
			return
		}
	}
	http.Redirect(_w, _r, "/list/", http.StatusFound)
}

func declinePreview(_id string, _message string, _w http.ResponseWriter, _r *http.Request) {
	g_dblock.Lock()
	defer g_dblock.Unlock()

	sendPrivateMessage(_id, _message, false)

	query := fmt.Sprintf("DELETE FROM mapchange WHERE idmapchange = %s", _id)
	if err := g_db.Query(query); err != nil {
		fmt.Printf("Database query error: %s\n", err.Error())
		return
	}
	http.Redirect(_w, _r, "/list/", http.StatusFound)
}

// Hackzorz: Inject PM in website database
// je weet
func sendPrivateMessage(_id string, _extra string, _accepted bool) {
	query := fmt.Sprintf("SELECT username FROM mapchange WHERE idmapchange = %s", _id)
	if err := g_db.Query(query); err != nil {
		fmt.Printf("Database query error: %s\n", err.Error())
		return
	}

	result, err := g_db.UseResult()
	if err != nil {
		fmt.Printf("Query error: %s", err.Error())
		return
	}
	defer result.Free()
	row := result.FetchMap()
	if row != nil {
		username := row["username"].(string)
		InjectMessage(username, _extra, _accepted)
	}
}

const page = `
<!doctype html>
<head>
        <meta charset="UTF-8">
        <meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1">
        
        <title>PU MapEditor Acceptor</title>

        <link rel="stylesheet" href="/inc/css/style.css?v=2">
</head>
<body>
<h1>PU MapEditor Acceptor</h1>
<p><b>USERNAME</b></p>
<p>COMMITLOG</p>
<div id="slider"></div>
<form id="blocker">
        <label for="blocktoggle">Block layer</label><input type="checkbox" id="blocktoggle" />
</form>
<div id="switch">
        <img src="PAARD3" id="blockimage" />
        <img src="PAARD1" id="topimage" />
        <img src="PAARD2" id="image" />
        <div style="clear: both"></div>
</div>

<div id="controls">
        <form name="input" action="/changehandler" method="post">
                <input type="radio" name="handle" value="Accept" /> Accept</br>
                <input type="radio" name="handle" value="Decline" /> Decline</br>
                <input type="hidden" value="CHANGEID" name="change_id" />
                Extra message (optional)</br><textarea name="reason" rows="4" cols="30"></textarea></br>
                <input type="submit" value="Submit" name="submit" />
        </form>
</div>

<script src="/inc/js/jquery-1.5.1.min.js"></script>
<script src="/inc/js/plugins.js"></script>
<script src="/inc/js/jquery-ui-1.8.11.custom.min.js"></script>
<script src="/inc/js/script.js"></script>
<link type="text/css" href="/inc/css/ui-lightness/jquery-ui-1.8.11.custom.css" rel="Stylesheet" />
</body>
</html>
`

const adduserpage = `
<form action="/adduser/" method="post">
<div>Username: <div><input name="username"/></div>
<div>Password: <div><input type="password" name="password"/></div>
<div><input type="submit" value="Add"></div>
</form>
`*/
