package pu.web.client.gui;

import java.util.ArrayList;

import pu.web.client.PU_Rect;

public class Container extends Element
{
	private ArrayList<Element> mChildren = new ArrayList<Element>();
	
	public Container(int x, int y, int width, int height)
	{
		super(x, y, width, height);
	}
	
	public void fillFocusList(ArrayList<Element> elements)
	{
		for(Element child : mChildren)
		{
			if(child instanceof Container)
			{
				((Container) child).fillFocusList(elements);
			}
			else
			{
				if(child.isFocusable())
				{
					elements.add(child);
				}
			}		
		}
	}
	
	public ArrayList<Element> getChildren()
	{
		return mChildren;
	}
	
	public void addChild(Element element)
	{
		element.setParent(this);
		mChildren.add(element);
	}
	
	public void removeChild(Element element)
	{
		element.setParent(null);
		mChildren.remove(element);
	}
	
	@Override
	public void draw(PU_Rect drawArea)
	{
		if(isVisible() && inDrawArea(drawArea))
		{
			PU_Rect childDrawArea = new PU_Rect(drawArea);
			childDrawArea.x += getRect().x;
			childDrawArea.y += getRect().y;
			childDrawArea.width = getRect().width;
			childDrawArea.height = getRect().height;
			
			for(Element child : mChildren)
			{
				if(child.isVisible() && child.inDrawArea(drawArea))
				{
					child.draw(childDrawArea);
				}
			}
		}
	}
	
	@Override
	public void mouseDown(int x, int y)
	{
		for(Element child : mChildren)
		{
			if(child.isVisible())
			{
				child.mouseDown(x - getRect().x, y - getRect().y);
			}
		}
	}
	
	@Override
	public void mouseUp(int x, int y)
	{
		for(Element child : mChildren)
		{
			if(child.isVisible())
			{
				child.mouseUp(x - getRect().x, y - getRect().y);
			}
		}
	}
	
	@Override
	public void mouseMove(int x, int y)
	{
		for(Element child : mChildren)
		{
			if(child.isVisible())
			{
				child.mouseMove(x - getRect().x, y - getRect().y);
			}
		}
	}
	
	@Override
	public void mouseScroll(int direction)
	{
		for(Element child : mChildren)
		{
			if(child.isVisible())
			{
				child.mouseScroll(direction);
			}
		}
	}
	
	@Override
	public void keyDown(int button)
	{
		for(Element child : mChildren)
		{
			if(child.isVisible())
			{
				child.keyDown(button);
			}
		}
	}
	
	@Override
	public void keyUp(int button)
	{
		for(Element child : mChildren)
		{
			if(child.isVisible())
			{
				child.keyUp(button);
			}
		}
	}
	
	@Override
	public void textInput(int charCode)
	{
		for(Element child : mChildren)
		{
			if(child.isVisible())
			{
				child.textInput(charCode);
			}
		}
	}
}
