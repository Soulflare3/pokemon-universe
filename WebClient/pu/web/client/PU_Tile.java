package pu.web.client;

public class PU_Tile
{
	public static final int TILE_WIDTH  = 48;
	public static final int TILE_HEIGHT = 48;
	
	public static final int TILE_BLOCKING       = 1;
	public static final int TILE_WALK           = 2;
	public static final int TILE_SURF           = 3;
	public static final int TILE_BLOCKTOP       = 4;
	public static final int TILE_BLOCKBOTTOM    = 5;
	public static final int TILE_BLOCKRIGHT     = 6;
	public static final int TILE_BLOCKLEFT      = 7;
	public static final int TILE_BLOCKCORNER_TR = 8;
	public static final int TILE_BLOCKCORNER_BR = 9;
	public static final int TILE_BLOCKCORNER_BL = 10;
	public static final int TILE_BLOCKCORNER_TL = 11;
	
	private int mPositionX = 0;
	private int mPositionY = 0;
	private int mMovement = TILE_WALK;
	private PU_Layer mLayers[] = new PU_Layer[4];
	
	private PU_Tile mNorthNeighbour = null;
	private PU_Tile mEastNeighbour = null;
	private PU_Tile mSouthNeighbour = null;
	private PU_Tile mWestNeighbour = null;
	
	public PU_Tile(int x, int y)
	{
		mPositionX = x;
		mPositionY = y;
	}
	
	public void drawLayer(int layer, int x, int y)
	{
		if(mLayers[layer] != null)
		{	
			int drawX = (x * TILE_WIDTH) - TILE_WIDTH - 22 + PUWeb.game().getScreenOffsetX();
			int drawY = (y * TILE_HEIGHT) - TILE_HEIGHT + PUWeb.game().getScreenOffsetY();
			
			mLayers[layer].draw(drawX, drawY);
		}
	}
	
	public int getX()
	{
		return mPositionX;
	}
	
	public int getY()
	{
		return mPositionY;
	}
	
	public void setMovement(int movement)
	{
		mMovement = movement;
	}
	
	public int getMovement()
	{
		return mMovement;
	}
	
	public void addLayer(int layer, int id)
	{
		if(mLayers[layer] == null)
		{
			mLayers[layer] = new PU_Layer(id);
		}
		else
		{
			mLayers[layer].setId(id);
		}
	}
	
	public void removeLayer(int layer)
	{
		mLayers[layer] = null;
	}
	
	public PU_Layer getLayer(int layer)
	{
		return mLayers[layer];
	}
	
	public long getSignature()
	{
		long signature = (long)mMovement;
		int shift = 16;
		for(int i = 0; i < 3; i++)
		{
			if(mLayers[i] != null)
			{
				signature |= ((long)mLayers[i].getId() << shift);
			}
			shift += 16;
		}
		return signature;
	}
	
	public static long TILE_INDEX(int _x, int _y, int _z) 
	{
		long x64 = 0;
		if(_x < 0) {
			x64 = (((long)1) << 34)  | (~(((long)_x)-1) << 18);
		} else {
			x64 = (((long)_x) << 24);
		}
		
		long y64 = 0;
		if(_y < 0) {
			y64 = (((long)1) << 17)  | (~(((long)_y)-1) << 1);
		} else {
			y64 = (((long)_y) << 1);
		}

		long z64 = (long)_z;
		long index = (long)(x64 | y64 | z64);
		
		return index;
	}
	
	public PU_Tile getNorthNeighbour()
	{
		return mNorthNeighbour;
	}
	
	public void setNorthNeighbour(PU_Tile tile)
	{
		mNorthNeighbour = tile;
	}
	
	public PU_Tile getEastNeighbour()
	{
		return mEastNeighbour;
	}
	
	public void setEastNeighbour(PU_Tile tile)
	{
		mEastNeighbour = tile;
	}
	
	public PU_Tile getSouthNeighbour()
	{
		return mSouthNeighbour;
	}
	
	public void setSouthNeighbour(PU_Tile tile)
	{
		mSouthNeighbour = tile;
	}
	
	public PU_Tile getWestNeighbour()
	{
		return mWestNeighbour;
	}
	
	public void setWestNeighbour(PU_Tile tile)
	{
		mWestNeighbour = tile;
	}
}
