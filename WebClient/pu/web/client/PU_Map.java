package pu.web.client;

import java.util.Collection;
import java.util.HashMap;

public class PU_Map
{
	private HashMap<Long, PU_Tile> mTileMap = new HashMap<Long, PU_Tile>();
	private HashMap<Long, PU_Creature> mCreatureMap = new HashMap<Long, PU_Creature>();
	
	public PU_Map()
	{
		
	}
	
	public PU_Tile addTile(int x, int y)
	{
		long posIndex = PU_Tile.TILE_INDEX(x,y,0);
		PU_Tile tile = null;
		if(!mTileMap.containsKey(posIndex))
		{
			tile = new PU_Tile(x, y);
			mTileMap.put(posIndex, tile);
		}
		return tile;
	}
	
	public void removeTile(PU_Tile tile)
	{
		long posIndex = PU_Tile.TILE_INDEX(tile.getX(), tile.getY(), 0);
		mTileMap.remove(posIndex);
	}

	public PU_Tile getTile(int x, int y)
	{
		long posIndex = PU_Tile.TILE_INDEX(x,y,0);
		return mTileMap.get(posIndex);
	}
	
	public PU_Tile getTile(long tileIndex)
	{
		return mTileMap.get(tileIndex);
	}
	
	public PU_Creature getCreatureById(long id)
	{
		return mCreatureMap.get(id);
	}
	
	public void addCreature(PU_Creature creature)
	{
		if(!mCreatureMap.containsKey(creature.getId()))
		{
			mCreatureMap.put(creature.getId(), creature);
		}
	}
	
	public int getCreatureCount()
	{
		return mCreatureMap.size();
	}
	
	public Collection<PU_Creature> getCreatures()
	{
		return mCreatureMap.values();
	}
}
