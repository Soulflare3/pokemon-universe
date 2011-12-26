package pu.web.client;

import java.util.Collection;

import pu.web.client.resources.fonts.Fonts;

public class PU_Game
{
	public static final int GAMESTATE_UNLOADING = 0;
	public static final int GAMESTATE_LOADING = 1;
	public static final int GAMESTATE_LOGIN = 2;
	public static final int GAMESTATE_WORLD = 3;
	public static final int GAMESTATE_BATTLE_INIT = 4;
	public static final int GAMESTATE_BATTLE = 5;
	
	public static final int NUMTILES_X = 23;
	public static final int NUMTILES_Y = 17;
	
	public static final int MID_X = 11;
	public static final int MID_Y = 8;
	
	private int mState = GAMESTATE_LOADING;
	private int mScreenOffsetX = 0;
	private int mScreenOffsetY = 0;
	private PU_Player mSelf = null;
	private int mLastDirKey = 0;
	
	public PU_Game()
	{
		
	}
	
	public int getState()
	{
		return mState;
	}
	
	public void setState(int state)
	{
		mState = state;
	}
	
	public void draw()
	{
		switch(mState)
		{
			case GAMESTATE_LOADING:
			{
				PU_Font font = PUWeb.resources().getFont(Fonts.FONT_PURITAN_BOLD_14);
				if(font != null)
				{
					font.drawText("Loading, please wait...", 10, 10);
					
					font.drawText("Fonts: " + PUWeb.resources().getFontLoadProgress() + "%", 10, 40);
					font.drawText("GUI: " + PUWeb.resources().getGuiImageLoadProgress() + "%", 10, 70);
					font.drawText("Tiles: " + PUWeb.resources().getTileLoadProgress() + "%", 10, 100);
				}
			}
			break;
			
			case GAMESTATE_LOGIN:
			{
				
			}
			break;
			
			case GAMESTATE_WORLD:
			{
				drawWorld();
			}
			break;
		}
	}
	
	public void drawWorld()
	{
		int[] offset = getScreenOffset();
		mScreenOffsetX = offset[0];
		mScreenOffsetY = offset[1];
		
		PU_Tile layer2tiles[] = new PU_Tile[NUMTILES_X * NUMTILES_Y];
		int layer2tilesCount = 0;
		
		PU_Creature walkers[] = new PU_Creature[PUWeb.map().getCreatureCount()];
		int walkerCount = 0;
		
		for(int x = 0; x < NUMTILES_X; x++)
		{
			for(int y = 0; y < NUMTILES_Y; y++)
			{
				int mapx = (mSelf.getX()-MID_X) + x;
				int mapy = (mSelf.getY()-MID_Y) + y;
				
				PU_Tile tile = PUWeb.map().getTile(mapx, mapy);
				if(tile != null)
				{
					for(int i = 0; i < 2; i++)
					{
						tile.drawLayer(i, x, y);
					}
					if(tile.getLayer(2) != null)
					{
						layer2tiles[layer2tilesCount] = tile;
						layer2tilesCount++;
					}
				}
			}
		}
		
		for(int i = 0; i < layer2tilesCount; i++)
		{
			PU_Tile tile = layer2tiles[i];
			
			int screenx = MID_X - (mSelf.getX() - tile.getX());
			int screeny = MID_Y - (mSelf.getY() - tile.getY());
			
			tile.drawLayer(2, screenx, screeny);
		}
		
		Collection<PU_Creature> creatures = PUWeb.map().getCreatures();
		for(PU_Creature creature : creatures)
		{
			if(creature.isWalking())
			{
				walkers[walkerCount] = creature;
				walkerCount++;
			}
		}
		
		for(int i = 0; i < walkerCount; i++)
		{
			if(walkers[i] instanceof PU_Player)
				((PU_Player)walkers[i]).updateWalk();
		}
	}
	
	public int getScreenOffsetX()
	{
		return mScreenOffsetX;
	}
	
	public int getScreenOffsetY()
	{
		return mScreenOffsetY;
	}
	
	public int[] getScreenOffset()
	{
		int[] coordinates = new int[]{0, 0};
		if(mSelf != null && mSelf.isWalking())
		{
			switch(mSelf.getDirection())
			{
			case PU_Creature.DIR_NORTH:
				coordinates[1] = 0 - (PU_Tile.TILE_HEIGHT - mSelf.getOffset());
				break;
				
			case PU_Creature.DIR_EAST:
				coordinates[0] = PU_Tile.TILE_WIDTH - mSelf.getOffset();
				break;
				
			case PU_Creature.DIR_SOUTH:
				coordinates[1] = PU_Tile.TILE_HEIGHT - mSelf.getOffset();
				break;
				
			case PU_Creature.DIR_WEST:
				coordinates[0] = 0 - (PU_Tile.TILE_WIDTH - mSelf.getOffset());
				break;
			}
		}
		return coordinates;
	}
	
	public void setSelf(PU_Player player)
	{
		mSelf = player;
	}
	
	public PU_Player getSelf()
	{
		return mSelf;
	}
	
	public int getLastDirKey()
	{
		return mLastDirKey;
	}
	
	public void keyDown(int keycode)
	{
		if(mState == GAMESTATE_WORLD)
		{
			if(mSelf != null)
			{
				boolean ctrlDown = PUWeb.events().isKeyDown(17);
				switch(keycode)
				{
				case PU_Events.KEY_LEFT:
					if(ctrlDown)
					{
						if(!mSelf.isWalking())
						{
							mSelf.turn(PU_Creature.DIR_WEST, true);
						}
					}
					else
					{
						mSelf.walk(PU_Creature.DIR_WEST);
						mLastDirKey = keycode;
					}
					break;
					
				case PU_Events.KEY_UP:
					if(ctrlDown)
					{
						if(!mSelf.isWalking())
						{
							mSelf.turn(PU_Creature.DIR_NORTH, true);
						}
					}
					else
					{
						mSelf.walk(PU_Creature.DIR_NORTH);
						mLastDirKey = keycode;
					}
					break;
					
				case PU_Events.KEY_RIGHT:
					if(ctrlDown)
					{
						if(!mSelf.isWalking())
						{
							mSelf.turn(PU_Creature.DIR_EAST, true);
						}
					}
					else
					{
						mSelf.walk(PU_Creature.DIR_EAST);
						mLastDirKey = keycode;
					}
					break;
					
				case PU_Events.KEY_DOWN:
					if(ctrlDown)
					{
						if(!mSelf.isWalking())
						{
							mSelf.turn(PU_Creature.DIR_SOUTH, true);
						}
					}
					else
					{
						mSelf.walk(PU_Creature.DIR_SOUTH);
						mLastDirKey = keycode;
					}
					break;
				}
			}
		}
	}
}
