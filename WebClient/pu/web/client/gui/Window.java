package pu.web.client.gui;

import java.util.ArrayList;

public class Window extends Container
{
	private String mTitle = "";
	ArrayList<Window> mPopups = new ArrayList<Window>();
	
	public Window(int x, int y, int width, int height)
	{
		super(x, y, width, height);
	}
	
	public String getTitle()
	{
		return mTitle;
	}
	
	public void addChild(Element element)
	{
		element.setParent(this);
		getChildren().add(element);
	}
	
	public void keyDown(int button)
	{
		if(button == 0/*TODO: TAB*/)
		{
			nextFocus();
		} 
		else
		{
			super.keyDown(button);
		}
	}
	
	public void nextFocus()
	{
		ArrayList<Element> elements = new ArrayList<Element>();
		fillFocusList(elements);
		
		if(elements.size() > 0)
		{
			if(elements.size() > 1)
			{
				int currentFocusIndex = 0;
				for(int i = 0; i < elements.size(); i++)
				{
					Element element = elements.get(i);
					if(element.hasFocus())
					{
						currentFocusIndex = i;
					}
					
					element.setFocus(false);
				}
				
				currentFocusIndex++;
				if(currentFocusIndex > elements.size()-1)
				{
					currentFocusIndex = 0;
				}
				
				elements.get(currentFocusIndex).setFocus(true);
			}
			else
			{
				elements.get(0).setFocus(true);
			}
		}
	}
	
	public void focusElement(Element element)
	{
		ArrayList<Element> elements = new ArrayList<Element>();
		fillFocusList(elements);
		
		for(Element curElement : elements)
		{
			curElement.setFocus(false);
		}
		
		element.setFocus(true);
	}
}
