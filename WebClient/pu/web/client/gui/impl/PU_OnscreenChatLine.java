package pu.web.client.gui.impl;

public class PU_OnscreenChatLine
{
	public static final int ONSCREENCHAT_TICKS = 3000;
	
	private String mText;
	private int mTicks;
	
	public PU_OnscreenChatLine(String text)
	{
		mText = text;
		mTicks = ONSCREENCHAT_TICKS;
	}
	
	public void setText(String text)
	{
		mText = text;
	}
	
	public String getText()
	{
		return mText;
	}
	
	public void setTicks(int ticks)
	{
		mTicks = ticks;
	}
	
	public int getTicks()
	{
		return mTicks;
	}
}
