package pu.web.client;

import pu.web.client.resources.fonts.Fonts;

public class PU_Game
{
	public static final int GAMESTATE_UNLOADING = 0;
	public static final int GAMESTATE_LOADING = 1;
	public static final int GAMESTATE_LOGIN = 2;
	public static final int GAMESTATE_WORLD = 3;
	public static final int GAMESTATE_BATTLE_INIT = 4;
	public static final int GAMESTATE_BATTLE = 5;
	
	private int mState = GAMESTATE_LOADING;
	
	public PU_Game()
	{
		
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
				}
			}
			break;
			
			case GAMESTATE_LOGIN:
			{
				
			}
			break;
		}
	}
}
