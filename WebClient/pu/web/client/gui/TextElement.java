package pu.web.client.gui;

import pu.web.client.PU_Color;
import pu.web.client.PU_Font;

public class TextElement extends Element
{
	private PU_Font mFont = null;
	private PU_Color mColor = new PU_Color(255, 255, 255, 255);
	private boolean mBold = false;
	private boolean mItalic = false;
	private boolean mUnderlined = false;
	
	public TextElement(int x, int y, int width, int height)
	{
		super(x, y, width, height);
	}
	
	public void setFont(PU_Font font)
	{
		mFont = font;
	}
	
	public PU_Font getFont()
	{
		return mFont;
	}
	
	public void setFontColor(int red, int green, int blue)
	{
		mColor.r = red;
		mColor.g = green;
		mColor.b = blue;
	}
	
	public PU_Color getFontColor()
	{
		return mColor;
	}
	
	public void setFontStyle(boolean bold, boolean italic, boolean underlined)
	{
		mBold = bold;
		mItalic = italic;
		mUnderlined = underlined;
	}
	
	public boolean isBold()
	{
		return mBold;
	}
	
	public boolean isItalic()
	{
		return mItalic;
	}
	
	public boolean isUnderlined()
	{
		return mUnderlined;
	}
}
