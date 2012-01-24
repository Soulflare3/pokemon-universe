package pu.web.client.gui.impl;

import pu.web.client.PU_Color;

public class PU_TextPart
{
	private String mText;
	private PU_Color mColor;
	
	public PU_TextPart(String text, int red, int green, int blue)
	{
		mText = text;
		mColor = new PU_Color(red, green, blue, 255);
	}
	
	public void setText(String text)
	{
		mText = text;
	}
	
	public String getText()
	{
		return mText;
	}
	
	public void setColor(int red, int green, int blue)
	{
		mColor = new PU_Color(red, green, blue, 255);
	}
	
	public PU_Color getColor()
	{
		return mColor;
	}
}
