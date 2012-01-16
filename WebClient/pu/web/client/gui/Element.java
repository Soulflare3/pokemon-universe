package pu.web.client.gui;

import pu.web.client.PU_Rect;

public class Element
{
	private PU_Rect mRect = null;
	private boolean mVisible = true;
	private Element mParent = null;
	private boolean mFocus = false;
	private boolean mFocusable = false;
	
	public Element(int x, int y, int width, int height) 
	{
		mRect = new PU_Rect(x, y, width, height);
		mVisible = true;
	}
	
	public void draw(PU_Rect drawArea)
	{
		
	}
	
	public void mouseDown(int x, int y)
	{
		
	}
	
	public void mouseUp(int x, int y)
	{
		
	}
	
	public void mouseMove(int x, int y)
	{
		
	}
	
	public void mouseScroll(int direction)
	{
		
	}
	
	public void keyDown(int button)
	{
		
	}
	
	public void keyUp(int button)
	{
		
	}
	
	public void textInput(int charCode)
	{

	}
	
	public void setVisible(boolean visible)
	{
		this.mVisible = visible;
	}
	
	public boolean isVisible() 
	{
		return mVisible;
	}
	
	public PU_Rect getRect()
	{
		return mRect;
	}
	
	public Element getParent()
	{
		return mParent;
	}
	
	public void setParent(Element element)
	{
		mParent = element;
	}
	
	public boolean hasFocus()
	{
		return mFocus;
	}
	
	public void setFocus(boolean focus)
	{
		mFocus = focus;
	}
	
	public void setFocusable(boolean focusable)
	{
		mFocusable = focusable;
	}
	
	public boolean isFocusable()
	{
		return mFocusable;
	}
	
	public Window getWindow()
	{
		if(mParent != null)
		{
			if(mParent instanceof Window)
			{
				return (Window)mParent;
			}
			
			return mParent.getWindow();
		}
		return null;
	}
	
	public boolean inDrawArea(PU_Rect drawArea)
	{
		if(mRect == null || drawArea == null)
			return false;
		
		if(drawArea.contains(mRect.x, mRect.y))
			return true;
		
		if(drawArea.contains(mRect.x+mRect.width, mRect.y))
			return true;
		
		if(drawArea.contains(mRect.x, mRect.y+mRect.height))
			return true;
		
		if(drawArea.contains(mRect.x+mRect.width, mRect.y+mRect.height))
			return true;
		
		return false;
	}
	
}
