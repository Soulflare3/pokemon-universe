package pu.web.client.gui.impl;

import java.util.ArrayList;

import pu.web.client.gui.Element;

public class PU_Chatbox extends Element
{
	public static final int CHATBOX_BUFFERSIZE = 100;
	
	private ArrayList<PU_Text> mLines = new ArrayList<PU_Text>();
	
	public PU_Chatbox(int x, int y, int width, int height)
	{
		super(x, y, width, height);
	}
	
	public void addLine(PU_Text line)
	{
		if(mLines.size() > CHATBOX_BUFFERSIZE)
		{
			mLines.remove(0);
		}
		mLines.add(line);
		
		//updateScrollbar
	}
	
	public void addText(PU_Text textToAdd)
	{
		int curSize = 0;
		String curText = "";
		PU_Text newText = new PU_Text(textToAdd.getFont());
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
					int wordSize = textToAdd.getFont().getStringWidth(word);
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
							newText.add(curText, textToAdd.getPart(part).getColor());
							addLine(newText);
							newText = new PU_Text(textToAdd.getFont());
							
							curText = "";
							curSize = 0;
						}
						else
						{
							for(int i = 0; i < word.length(); i++)
							{
								int charWidth = textToAdd.getFont().getStringWidth("" + word.charAt(i));
								if(curSize+charWidth > maxWidth)
								{
									curText += "-";
									
									newText.add(curText, textToAdd.getPart(part).getColor());
									addLine(newText);
									newText = new PU_Text(textToAdd.getFont());
									
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
					newText.add(curText, textToAdd.getPart(part).getColor());
					curText = "";
					if(part+1 >= text.length()) 
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
}
