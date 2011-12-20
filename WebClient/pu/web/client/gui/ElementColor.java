package pu.web.client.gui;

import pu.web.client.PU_Color;

public class ElementColor
{
	private PU_Color mColor = new PU_Color(255, 255, 255, 255);
	
	public ElementColor()
	{
		
	}
	
	public PU_Color getColor()
	{
		return mColor;
	}
	
	public void setColor(int red, int green, int blue)
	{
		mColor.r = red;
		mColor.g = green;
		mColor.b = blue;
	}
}