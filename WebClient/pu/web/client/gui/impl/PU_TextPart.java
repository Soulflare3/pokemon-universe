package pu.web.client.gui.impl;

import pu.web.client.PU_Color;

public class PU_TextPart
{
	private String mText;
	private PU_Color mColor;
	
	public PU_TextPart(String text, PU_Color color)
	{
		mText = text;
		mColor = color;
	}
	
	public void setText(String text)
	{
		mText = text;
	}
	
	public String getText()
	{
		return mText;
	}
	
	public void setColor(PU_Color color)
	{
		mColor = color;
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
