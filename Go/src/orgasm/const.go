package main

const (
	HEADER_LOGIN 		   byte = 0x00
	HEADER_REQUEST_MAP_PIECE 	= 0x01
	HEADER_TILE_CHANGE			= 0x02
	HEADER_REQUEST_MAP_LIST		= 0x03
	HEADER_ADD_MAP				= 0x04
	HEADER_DELETE_MAP			= 0x05
	HEADER_UPDATE_TILEEVENT		= 0x06
	HEADER_ADD_NPC				= 0x07
	HEADER_EDIT_NPC_OUTFIT		= 0x08
	HEADER_EDIT_NPC_POSITION	= 0x09
	HEADER_DELETE_NPC			= 0x0A
	HEADER_GET_NPC_DATA			= 0x0B
	HEADER_GET_NPC_EVENTS		= 0x0C
)

const (
	QUERY_SELECT_ACCOUNT string = "SELECT * FROM mapchange_account WHERE username = '%s'"
	
	//TODO The idlocation and the idtielis a constant 1 yet
	QUERY_INSERT_TILE string = "INSERT INTO tile (idtile, x, y, z, movement, idlocation, idtile_event) VALUES (%d, %d, %d, %d, %d, 1, %s);\n"
	QUERY_UPDATE_TILE string = "UPDATE tile SET movement='%d', idtile_event=%s WHERE idtile='%d';\n"
	QUERY_DELETE_TILE string = "DELETE FROM tile WHERE idtile='%d';\n"
	QUERY_LOAD_TILES string = 	"SELECT t.`x`, t.`y`, t.`z`, t.`idlocation`, t.`movement`, t.`idtile_event`," +
								" tl.`sprite`, tl.`layer`," +
								" t.`idtile`, tl.`idtilelayer`" +
								" FROM tile `t`" +
								" INNER JOIN tile_layer `tl` ON t.`idtile` = tl.`idtile`" 
								
	QUERY_LOAD_TILE_ID string = "SELECT t.`id` FROM tile `t` WHERE t.`x` = %d AND t.`y` = %d AND t.`z` = %d"
	
	QUERY_INSERT_EVENT string = "INSERT INTO tile_events (eventtype, param1, param2, param3, param4, param5, param6, param7, param8) " + 
									"VALUES (%d, %s, %s, %s, %s, %s, %s, %s, %s)"
	QUERY_UPDATE_EVENT string = "UPDATE tile_events SET eventtype=%d, param1='%s', param1='%s', param1='%s', param1='%s', param1='%s', " + 
									"param1='%s', param1='%s', param1='%s' WHERE idtile_event=%d"							
	QUERY_DELETE_EVENT string = "DELETE FROM tile_events WHERE idtile_event = %d"

	QUERY_INSERT_TILELAYER string = "INSERT INTO tile_layer (idtilelayer, tileid, layer, sprite) VALUES (%d, %d, %d, %d);\n"
	QUERY_UPDATE_TILELAYER string = "UPDATE tile_layer SET sprite='%d' WHERE idtilelayer='%d';\n"
	QUERY_DELETE_TILELAYER string = "DELETE FROM tile_layer WHERE idtilelayer='%d';\n"
	
	QUERY_INSERT_MAP string = "INSERT INTO map (name) VALUES ('%s')"
	QUERY_DELETE_MAP string = "DELETE map, tile, tile_layer FROM map " +  
								"LEFT JOIN tile ON map.idmap = tile.z "+ 
								"LEFT JOIN tile_layer ON tile.idtile = tile_layer.idtile " +
								"WHERE map.idmap= '%d'"
								
	QUERY_SELECT_NPCS string = "SELECT npc.idnpc, npc.name, npc_outfit.head, npc_outfit.nek, npc_outfit.upper, npc_outfit.lower, " + 
								 "npc_outfit.feet, npc.position, npc_events.event, npc_events.initId " + 
								 "FROM npc " + 
								 "INNER JOIN npc_outfit ON npc.idnpc = npc_outfit.idnpc " + 
								 "INNER JOIN npc_events ON npc.idnpc = npc_events.idnpc "+ 
								 "ORDER BY npc.idnpc"
	QUERY_INSERT_NPC string = "INSERT INTO npc (name) VALUES ('%s')"
	QUERY_UPDATE_NPC string = "UPDATE npc SET name='%s', position=%d WHERE idnpc=%d" 
	QUERY_DELETE_NPC string = "DELETE FROM npc WHERE idnpc=%d"
		
	QUERY_INSERT_NPC_OUTFIT string = "INSERT INTO npc_outfit (idnpc,head,nek,upper,lower,feet) VALUES (%d,%d,%d,%d,%d,%d)"
	QUERY_UPDATE_NPC_OUTFIT string = "UPDATE npc_outfit SET head=%d, nek=%d, upper=%d, lower=%d, feet=%d WHERE idnpc = %d"
	QUERY_DELETE_NPC_OUTFIT string = "DELETE FROM npc_outfit WHERE idnpc = %d"
	
	QUERY_INSERT_NPC_EVENT string = "INSERT INTO npc_events (idnpc) VALUES (%d)"
	QUERY_UPDATE_NPC_EVENT string = "UPDATE npc_events SET event='%s', initId='%d' WHERE idnpc = %d"
	QUERY_DELETE_NPC_EVENT string = "DELETE FROM npc_events WHERE idnpc = %d"
	
	QUERY_SELECT_NPC_POKEMON string = "SELECT idnpc_pokemon, idpokemon, iv_hp, iv_attack, iv_attack_spec, " + 
										"iv_defence, iv_defence_spec, iv_speed, gender, held_item " + 
										"FROM npc_pokemon WHERE idnpc=%d"
	QUERY_DELETE_NPC_POKEMON string = "DELETE FROM npc_pokemon WHERE idnpc_pokemon = %d"
	QUERY_INSERT_NPC_POKEMON string = ""
	QUERY_UPDATE_NPC_POKEMON string = ""
)

const (
	TILEBLOCK_BLOCK       int = 1
	TILEBLOCK_WALK            = 2
	TILEBLOCK_SURF            = 3
	TILEBLOCK_TOP             = 4
	TILEBLOCK_BOTTOM          = 5
	TILEBLOCK_RIGHT           = 6
	TILEBLOCK_LEFT            = 7
	TILEBLOCK_TOPRIGHT        = 8
	TILEBLOCK_BOTTOMRIGHT     = 9
	TILEBLOCK_BOTTOMLEFT      = 10
	TILEBLOCK_TOPLEFT         = 11
)

const (
	TILEEVENT_NONE	int = 0
	TILEEVENT_WARP		= 1
)