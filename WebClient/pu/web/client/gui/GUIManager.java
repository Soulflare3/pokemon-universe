package pu.web.client.gui;

import pu.web.client.PU_Font;

public class GUIManager
{
	Window mRoot = null;
	PU_Font mDefaultFont = null;

	public GUIManager(int x, int y, int width, int height, PU_Font defaultFont)
	{
		mRoot = new Window(x, y, width, height);
		mDefaultFont = defaultFont;
	}
	
	public Window getRoot()
	{
		return mRoot;
	}
	
	public PU_Font getDefaultFont()
	{
		return mDefaultFont;
	}
	
	public void draw() 
	{
		mRoot.draw(mRoot.getRect());
	}
	
	public void mouseDown(int x, int y)
	{
		mRoot.mouseDown(x, y);
	}
	
	public void mouseUp(int x, int y)
	{
		mRoot.mouseUp(x, y);
	}
	
	public void mouseMove(int x, int y)
	{
		mRoot.mouseMove(x, y);
	}
	
	public void mouseScroll(int direction)
	{
		mRoot.mouseScroll(direction);
	}
	
	public void keyDown(int button)
	{
		mRoot.keyDown(button);
	}
	
	public void keyUp(int button)
	{
		mRoot.keyUp(button);
	}
	
	public void textInput(int charCode)
	{
		mRoot.textInput(charCode);
	}
}
