package pu.web.client;

import java.util.ArrayList;
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
	private ArrayList<PU_PlayerName> mPlayerNames = new ArrayList<PU_PlayerName>();
	
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
		
		mPlayerNames.clear();
		
		PUWeb.engine().beginSpriteBatch();
		PU_Tile currentXTile = null;
		PU_Tile currentYTile = null;
		for(int x = 0; x < NUMTILES_X; x++)
		{
			if(currentXTile == null)
			{
				int mapx = (mSelf.getX()-MID_X) + x;
				int mapy = (mSelf.getY()-MID_Y);
				currentXTile = PUWeb.map().getTile(mapx, mapy);
				currentYTile = currentXTile;
			}
			
			for(int y = 0; y < NUMTILES_Y; y++)
			{
				if(currentYTile == null)
				{
					int mapx = (mSelf.getX()-MID_X) + x;
					int mapy = (mSelf.getY()-MID_Y) + y;
					currentYTile = PUWeb.map().getTile(mapx, mapy);
				}
				
				if(currentYTile != null)
				{
					for(int i = 0; i < 2; i++)
					{
						currentYTile.drawLayer(i, x, y);
					}
					if(currentYTile.getLayer(2) != null)
					{
						layer2tiles[layer2tilesCount] = currentYTile;
						layer2tilesCount++;
					}
					
					currentYTile = currentYTile.getSouthNeighbour();
				}
			}
			if(currentXTile != null)
			{
				currentXTile = currentXTile.getEastNeighbour();
				currentYTile = currentXTile;
			}
		}
		
		Collection<PU_Creature> creatures = PUWeb.map().getCreatures();
		for(PU_Creature creature : creatures)
		{
			int screenx = MID_X - (mSelf.getX() - creature.getX());
			int screeny = MID_Y - (mSelf.getY() - creature.getY());
			drawCreature(creature, screenx, screeny);
			
			if(creature.isWalking())
			{
				walkers[walkerCount] = creature;
				walkerCount++;
			}
		}
		
		for(int i = 0; i < layer2tilesCount; i++)
		{
			PU_Tile tile = layer2tiles[i];
			
			int screenx = MID_X - (mSelf.getX() - tile.getX());
			int screeny = MID_Y - (mSelf.getY() - tile.getY());
			
			tile.drawLayer(2, screenx, screeny);
		}
		PUWeb.engine().endSpriteBatch();
		
		PU_Font nameFont = PUWeb.resources().getFont(Fonts.FONT_PURITAN_BOLD_14);
		nameFont.setColor(255, 242, 0);
		for(PU_PlayerName name : mPlayerNames)
		{
			nameFont.drawBorderedText(name.name, name.x, name.y);
		}
		
		
		for(int i = 0; i < walkerCount; i++)
		{
			if(walkers[i] instanceof PU_Player)
				((PU_Player)walkers[i]).updateWalk();
		}
	}
	
	public void drawCreature(PU_Creature creature, int x, int y)
	{
		int[] offset = getScreenOffset();
		int offsetX = offset[0];
		int offsetY = offset[1];
		
		int drawX = 0;
		int drawY = 0;
		
		if(creature.isWalking()) 
		{
            switch(creature.getDirection()) 
            {
            case PU_Player.DIR_NORTH:
                    drawX = (x * PU_Tile.TILE_WIDTH) - PU_Tile.TILE_WIDTH - 22 + offsetX;
                    drawY = ((y * PU_Tile.TILE_HEIGHT) + (PU_Tile.TILE_HEIGHT - creature.getOffset())) - PU_Tile.TILE_HEIGHT + offsetY;
                    break;

            case PU_Player.DIR_EAST:
                    drawX = ((x * PU_Tile.TILE_WIDTH) - (PU_Tile.TILE_WIDTH - creature.getOffset())) - PU_Tile.TILE_WIDTH - 22 + offsetX;
                    drawY = (y * PU_Tile.TILE_HEIGHT) - PU_Tile.TILE_HEIGHT + offsetY;
                    break;

            case PU_Player.DIR_SOUTH:
                    drawX = (x * PU_Tile.TILE_WIDTH) - PU_Tile.TILE_WIDTH - 22 + offsetX;
                    drawY = ((y * PU_Tile.TILE_HEIGHT) - (PU_Tile.TILE_HEIGHT - creature.getOffset())) - PU_Tile.TILE_HEIGHT + offsetY;
                    break;

            case PU_Player.DIR_WEST:
                    drawX = ((x * PU_Tile.TILE_WIDTH) + (PU_Tile.TILE_WIDTH - creature.getOffset())) - PU_Tile.TILE_WIDTH - 22 + offsetX;
                    drawY = (y * PU_Tile.TILE_HEIGHT) - PU_Tile.TILE_HEIGHT + offsetY;
                    break;
            }
		} 
		else 
		{
            drawX = (x * PU_Tile.TILE_WIDTH) - PU_Tile.TILE_WIDTH - 22 + offsetX;
            drawY = (y * PU_Tile.TILE_HEIGHT) - PU_Tile.TILE_HEIGHT + offsetY;
		}
		creature.draw(drawX, drawY);
		
		if(creature instanceof PU_Player)
		{
			String name = ((PU_Player) creature).getName();
			int posHalf = (drawX - 48) + (((drawX + 96) - (drawX - 48)) / 2);
	        int nameHalf = PUWeb.resources().getFont(Fonts.FONT_PURITAN_BOLD_14).getStringWidth(name) / 2;
            int centerPos = posHalf - nameHalf;
            mPlayerNames.add(new PU_PlayerName(name, centerPos, drawY - 14));
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
						mLastDirKey = keycode;
						mSelf.walk(PU_Creature.DIR_WEST);
						
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
						mLastDirKey = keycode;
						mSelf.walk(PU_Creature.DIR_NORTH);
						
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
						mLastDirKey = keycode;
						mSelf.walk(PU_Creature.DIR_EAST);
						
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
						mLastDirKey = keycode;
						mSelf.walk(PU_Creature.DIR_SOUTH);
						
					}
					break;
				}
			}
		}
	}
}
