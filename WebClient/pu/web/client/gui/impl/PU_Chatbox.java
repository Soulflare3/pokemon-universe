package pu.web.client.gui.impl;

import java.util.ArrayList;

import pu.web.client.PUWeb;
import pu.web.client.PU_Font;
import pu.web.client.PU_Rect;
import pu.web.client.gui.Element;
import pu.web.client.gui.Scrollbar;

public class PU_Chatbox extends Element
{
	public static final int CHATBOX_BUFFERSIZE = 100;
	
	private ArrayList<PU_Text> mLines = new ArrayList<PU_Text>();
	private Scrollbar mScrollbar = null;
	private PU_Font mFont = null;
	
	public PU_Chatbox(int x, int y, int width, int height)
	{
		super(x, y, width, height);
		mFont = PUWeb.gui().getDefaultFont();
	}
	
	public PU_Font getFont()
	{
		return mFont;
	}
	
	public void setFont(PU_Font font)
	{
		mFont = font;
	}
	
	public void setScrollbar(Scrollbar scrollbar)
	{
		mScrollbar = scrollbar;
	}
	
	public int getLineCount()
	{
		return mLines.size();
	}
	
	public void addLine(PU_Text line)
	{
		if(mLines.size() > CHATBOX_BUFFERSIZE)
		{
			mLines.remove(0);
		}
		mLines.add(line);
		
		updateScrollbar();
	}
	
	public void addText(PU_Text textToAdd)
	{
		int curSize = 0;
		String curText = "";
		PU_Text newText = new PU_Text(mFont);
		int maxWidth = getRect().width - 6;
		
		if(textToAdd.getWidth() > maxWidth)
		{
			for(int part = 0; part < textToAdd.getSize(); part++)
			{
				String text = textToAdd.getPart(part).getText();
				int curPos = 0;
				while(curPos < text.length())
				{
					String word = nextWord(text, curPos);
					int wordSize = mFont.getStringWidth(word);
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
							PU_TextPart curPart = textToAdd.getPart(part);
							newText.add(curText, curPart.getColor().r, curPart.getColor().g, curPart.getColor().b);
							addLine(newText);
							newText = new PU_Text(mFont);
							
							curText = "";
							curSize = 0;
						}
						else
						{
							for(int i = 0; i < word.length(); i++)
							{
								int charWidth = mFont.getStringWidth("" + word.charAt(i));
								if(curSize+charWidth > maxWidth)
								{
									curText += "-";
									
									PU_TextPart curPart = textToAdd.getPart(part);
									newText.add(curText, curPart.getColor().r, curPart.getColor().g, curPart.getColor().b);
									addLine(newText);
									newText = new PU_Text(mFont);
									
									curText = "";
									curSize = 0;
									
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
					PU_TextPart curPart = textToAdd.getPart(part);
					newText.add(curText, curPart.getColor().r, curPart.getColor().g, curPart.getColor().b);
					curText = "";
					if(part+1 >= textToAdd.getSize()) 
					{
						addLine(newText);
					}
				}
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
	
	public void updateScrollbar()
	{
		if(mScrollbar != null)
		{
			int fontHeight = mFont.getLineHeight();
			int boxHeight = getRect().height - 6;
			int visibleLines = (int)((float)boxHeight / (float)fontHeight);
			
			int max = mLines.size() - visibleLines;
			if(max <= 0)
			{
				max = 0;
			}
			
			if(mScrollbar.getValue() == mScrollbar.getMaxValue())
			{
				mScrollbar.setMaxValue(max);
				mScrollbar.setValue(max);
			}
			else
			{
				mScrollbar.setMaxValue(max);
			}
		}
	}
	
	@Override
	public void draw(PU_Rect drawArea)
	{
		if(!isVisible())
			return;
		
		int fontHeight = mFont.getLineHeight();
		int boxHeight = getRect().height - 6;
		int visibleLines = (int)((float)boxHeight / (float)fontHeight);
		int scrollInc = 0;
		if(mScrollbar != null)
		{
			scrollInc = mScrollbar.getValue();
		}
		
		PU_Rect realRect = new PU_Rect(getRect().x + drawArea.x, getRect().y + drawArea.y, getRect().width, getRect().height);
		PU_Rect inRect = drawArea.intersection(realRect);
		
		int drawX = realRect.x + 3;
		int drawY = realRect.y + 3;
		
		for(int line = scrollInc; line < (visibleLines + scrollInc); line++)
		{
			if(line < mLines.size())
			{
				PU_Text text = mLines.get(line);
				if(text != null)
				{
					int numParts = text.getSize();
					for(int part = 0; part < numParts; part++)
					{
						PU_TextPart curPart = text.getPart(part);
						
						mFont.setColor(curPart.getColor().r, curPart.getColor().g, curPart.getColor().b);
						mFont.drawTextInRect(curPart.getText(), drawX, drawY, inRect);
						
						drawX += mFont.getStringWidth(curPart.getText());
					}
					drawX = realRect.x + 3;
					drawY += fontHeight;
				}
			}
		}
	}
}
