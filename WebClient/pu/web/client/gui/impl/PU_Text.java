package pu.web.client.gui.impl;

import java.util.ArrayList;

import pu.web.client.PU_Font;

public class PU_Text
{
	private PU_Font mFont = null;
	private ArrayList<PU_TextPart> mParts = new ArrayList<PU_TextPart>();
	private boolean mDirty = false;
	private int mWidth = 0;
	
	public PU_Text(PU_Font font)
	{
		mFont = font;
	}
	
	public PU_Font getFont()
	{
		return mFont;
	}
	
	public int getSize()
	{
		return mParts.size();
	}
	
	public void add(String text, int red, int green, int blue)
	{
		mParts.add(new PU_TextPart(text, red, green, blue));
		mDirty = true;
	}
	
	public void addToLast(String text)
	{
		PU_TextPart part = mParts.get(mParts.size()-1);
		if(part != null)
		{
			part.setText(part.getText() + text);
		}
		mDirty = true;
	}
	
	public PU_TextPart getPart(int index)
	{
		return mParts.get(index);
	}
	
	public String getAll()
	{
		String str = "";
		for(PU_TextPart part : mParts)
		{
			str += part.getText();
		}
		return str;
	}
	
	public int getWidth()
	{
		if(mDirty)
		{
			mWidth = mFont.getStringWidth(getAll());
			mDirty = false;
		}
		return mWidth;
	}
}
