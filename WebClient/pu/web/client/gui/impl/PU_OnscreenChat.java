package pu.web.client.gui.impl;

import java.util.ArrayList;

import pu.web.client.PUWeb;
import pu.web.client.PU_Creature;
import pu.web.client.PU_Font;
import pu.web.client.resources.fonts.Fonts;

public class PU_OnscreenChat
{
	private long mLastTicks;
	private ArrayList<PU_OnscreenChatMessage> mMessages = new ArrayList<PU_OnscreenChatMessage>();
	private boolean mVisible = false;
	
	public PU_OnscreenChat()
	{
		mLastTicks = System.currentTimeMillis();
	}
	
	public void setVisible(boolean visible)
	{
		mVisible = visible;
	}
	
	public void draw()
	{
		if(!mVisible)
			return;
		
		long ticks = System.currentTimeMillis() - mLastTicks;
		mLastTicks = System.currentTimeMillis();
		
		if(mMessages.size() > 0)
		{
			PU_Font font = PUWeb.resources().getFont(Fonts.FONT_ARIALBLK_BOLD_14_OUTLINE);
			
			int drawCount = 0;
			for(PU_OnscreenChatMessage message : mMessages)
			{
				drawCount += message.getName().length() + 6;
				for(PU_OnscreenChatLine line :  message.getLines())
				{
					drawCount += line.getText().length();
				}
			}
			
			PUWeb.engine().beginTextureBatch(font.getImage().getTexture(), font.getImage().getWidth(), drawCount, 0, 162, 232, 255);
			for(int i = 0; i < mMessages.size();)
			{
				PU_OnscreenChatMessage message = mMessages.get(i);
				if(message != null)
				{
					if(!message.draw((int)ticks))
					{
						mMessages.remove(i);
					}
					else
					{
						i++;
					}
				}
			}
			
			PUWeb.engine().endTextureBatch();
		}
	}
	
	public void add(String name, String text)
	{
		PU_Creature sender = PUWeb.map().getCreatureByName(name);
		if(sender == null)
			return;
		
		for(PU_OnscreenChatMessage message : mMessages)
		{
			if(message != null && message.getName().equals(name) && message.getX() == sender.getX() && message.getY() == sender.getY())
			{
				message.addText(text);
				return;
			}
		}
		
		PU_OnscreenChatMessage message = new PU_OnscreenChatMessage(name,sender.getX(), sender.getY());
		message.addText(text);
		mMessages.add(message);
	}
}
