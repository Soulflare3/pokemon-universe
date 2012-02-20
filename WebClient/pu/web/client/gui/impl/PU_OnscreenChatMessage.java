package pu.web.client.gui.impl;

import java.util.ArrayList;

import pu.web.client.PUWeb;
import pu.web.client.PU_Engine;
import pu.web.client.PU_Font;
import pu.web.client.PU_Game;
import pu.web.client.PU_Tile;
import pu.web.client.resources.fonts.Fonts;

public class PU_OnscreenChatMessage
{
	private String mName;
	private int mX;
	private int mY;
	private ArrayList<PU_OnscreenChatLine> mLines = new ArrayList<PU_OnscreenChatLine>();
	
	public PU_OnscreenChatMessage(String name, int x, int y)
	{
		mName = name;
		mX = x;
		mY = y;
	}
	
	public PU_OnscreenChatMessage(String name, int x, int y, String text)
	{
		mName = name;
		mX = x;
		mY = y;
		addText(text);
	}
	
	public void setName(String name)
	{
		mName = name;
	}
	
	public String getName()
	{
		return mName;
	}
	
	public void setPosition(int x, int y)
	{
		mX = x;
		mY = y;
	}
	
	public int getX()
	{
		return mX;
	}
	
	public int getY()
	{
		return mY;
	}
	
	public int getLineCount()
	{
		return mLines.size();
	}
	
	public ArrayList<PU_OnscreenChatLine> getLines()
	{
		return this.mLines;
	}
	
	public boolean draw(int ticks)
	{
		boolean ret = true;
		
		int[] offset = PUWeb.game().getScreenOffset();
		int offsetX = offset[0];
		int offsetY = offset[1];
		
		PU_Font font = PUWeb.resources().getFont(Fonts.FONT_ARIALBLK_BOLD_14_OUTLINE);
		
		int lineHeight = font.getLineHeight();
		int height = lineHeight + (lineHeight * mLines.size());
		
		boolean center = false;
		
		int drawX = PU_Game.MID_X - (PUWeb.game().getSelf().getX() - mX);
		int drawY = PU_Game.MID_Y - (PUWeb.game().getSelf().getY() - mY);
		
		drawX = (drawX * PU_Tile.TILE_WIDTH) - PU_Tile.TILE_WIDTH - 22 + offsetX;
		drawY = (drawY * PU_Tile.TILE_HEIGHT) - PU_Tile.TILE_HEIGHT + offsetY;
		
		if(drawY-height < 0)
		{
			drawY = 0;
		}
		else if(drawY > PU_Engine.SCREEN_HEIGHT)
		{
			drawY = PU_Engine.SCREEN_HEIGHT - height;
		}
		else
		{
			drawY -= height;
			drawY += lineHeight;
		}
		
		String header = mName + " says:";
		
		int widest = font.getStringWidth(header);
		
		for(PU_OnscreenChatLine line : mLines)
		{
			if(line != null)
			{
				int len = font.getStringWidth(line.getText());
				if(len > widest)
				{
					widest = len;
				}
			}
		}
		
		if(drawX - (int)Math.ceil((float)widest/2.0f) < 0)
		{
			drawX = 0;
		}
		else if(drawX + (int)Math.ceil((float)widest/2.0f) > PU_Engine.SCREEN_WIDTH)
		{
			drawX = PU_Engine.SCREEN_WIDTH - widest;
		}
		else
		{
			center = true;
		}
		
		int posHalf = 0;
		if(!center)
		{
			posHalf = drawX + (int)Math.ceil((float)((drawX+widest)-drawX)/2); 
		}
		else
		{
			posHalf = (drawX - ((int)Math.ceil((float)widest / 2.0))) + (int)(Math.ceil((float)((drawX+widest)-(drawX-((int)Math.ceil((float)widest/2.0))))/2));
		}
		
		int nameHalf = (int)(Math.floor((float)font.getStringWidth(header) / 2.0));
        int centerPos = posHalf - nameHalf;

        font.drawTextInBatch(header, centerPos, drawY);
        
        for(int i = 0; i < mLines.size(); i++)
        {
        	PU_OnscreenChatLine line = mLines.get(i);
        	if(line != null)
        	{
        		nameHalf = (int)Math.floor((float)font.getStringWidth(line.getText())/2.0);
        		centerPos = posHalf - nameHalf;
        		
        		font.drawTextInBatch(line.getText(), centerPos, drawY+((i+1)*lineHeight));
        	}
        }
        
        updateLines(ticks);
        if(mLines.size() <= 0)
		{
        	ret = false;
		}
		
		return ret;
	}
	
	public void addLine(String line)
	{
		if(mLines.size() >= 4)
		{
			mLines.remove(0);
		}
		mLines.add(new PU_OnscreenChatLine(line));
	}
	
	public void addText(String textToAdd)
	{
		PU_Font font = PUWeb.resources().getFont(Fonts.FONT_ARIALBLK_BOLD_14_OUTLINE);
		int curSize = 2;
		String curText = "";
		int maxWidth = 160;
		int textWidth = font.getStringWidth(textToAdd) + 2;
		
		if(textWidth > maxWidth)
		{
			String text = textToAdd;
			int curPos = 0;

			while(curPos < text.length())
			{
				String word = nextWord(text, curPos);
				int wordSize = font.getStringWidth(word);
				if(curSize+wordSize < maxWidth)
				{
					curText += word;
					curSize += wordSize;
					curPos += word.length();
				}
				else
				{
					if(!curText.equals(""))
					{
						addLine(curText);
						
						curText = "";
						curSize = 2;
					}
					else
					{
						for(int i = 0; i < word.length(); i++)
						{
							int charWidth = font.getStringWidth("" + word.charAt(i));
							if(curSize+charWidth > maxWidth)
							{
								curText += "-";
								
								addLine(curText);
								
								curText = "";
								curSize = 2;
								
								curPos += i;
								
								break;
							}
							curText += word.charAt(i);
							curSize += charWidth;
						}
					}
				}
			}
			if(!curText.equals(""))
			{
				addLine(curText);
			}
		}
		else
		{
			addLine(textToAdd);
		}
	}
	
	public String nextWord(String text, int start) 
	{
		for(int i = start; i < text.length(); i++)
		{
			if(text.charAt(i) == ' ')
			{
				return text.substring(start, i + 1);
			}
		}
		return text.substring(start);
	}	
	
	public void updateLines(int ticks)
	{
		for(int i = 0; i < mLines.size();)
		{
			PU_OnscreenChatLine line = mLines.get(i);
			if(line != null)
			{
				line.setTicks(line.getTicks()-ticks);
				if(line.getTicks() <= 0)
				{
					mLines.remove(i);
				}
				else
				{
					i++;
				}
			}
			else
			{
				mLines.remove(i);
			}
		}
	}
}
