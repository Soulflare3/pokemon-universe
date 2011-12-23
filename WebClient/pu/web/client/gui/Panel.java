package pu.web.client.gui;

import pu.web.client.PUWeb;
import pu.web.client.PU_Rect;

public class Panel extends Container
{
	private ElementColor mBackgroundColor = null;
	
	public Panel(int x, int y, int width, int height)
	{
		super(x, y, width, height);
	}
	
	public void setBackgroundColor(int red, int green, int blue)
	{
		if(mBackgroundColor == null)
			mBackgroundColor = new ElementColor();
		
		mBackgroundColor.setColor(red, green, blue);
	}
	
	@Override
	public void draw(PU_Rect drawArea)
	{
		if(mBackgroundColor != null)
		{
			PU_Rect realRect = new PU_Rect(getRect().x + drawArea.x, getRect().y + drawArea.y, getRect().width, getRect().height);
			PU_Rect inRect = drawArea.intersection(realRect);
			
			PUWeb.engine().setColor(mBackgroundColor.getColor().r, mBackgroundColor.getColor().g, mBackgroundColor.getColor().b, mBackgroundColor.getColor().a);
			PUWeb.engine().renderFillRect(inRect.x, inRect.y, inRect.width, inRect.height);
		}
		
		super.draw(drawArea);
	}
}