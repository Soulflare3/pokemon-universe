package pu.web.client.gui;

import pu.web.client.PUWeb;
import pu.web.client.PU_Font;
import pu.web.client.PU_Rect;

public class Label extends TextElement
{
	private String mText = "";

	public Label(int x, int y, String text)
	{
		super(x, y, 0, 0);
		
		mText = text;
		setFont(PUWeb.gui().getDefaultFont());
	}
	
	public String getText()
	{
		return mText;
	}
	
	public void setText(String text)
	{
		mText = text;
		updateSize();
	}
	
	@Override
	public void setFont(PU_Font font)
	{
		super.setFont(font);
		updateSize();
	}

	public void updateSize()
	{
		int width = 0;
		int height = 0;
		PU_Font font = getFont();
		if(font != null)
		{
			width = font.getStringWidth(mText);
			height = font.getLineHeight()+2;
		}
		
		getRect().width = width;
		getRect().height = height;
	}
	
	@Override
	public void draw(PU_Rect drawArea)
	{
		PU_Rect realRect = new PU_Rect(getRect().x + drawArea.x, getRect().y + drawArea.y, getRect().width, getRect().height);
		PU_Rect inRect = drawArea.intersection(realRect);
		
		PU_Font font = getFont();
		if(font != null)
		{
			font.setColor(getFontColor().r, getFontColor().g, getFontColor().b);
			font.drawTextInRect(mText, drawArea.x + getRect().x, drawArea.y + getRect().y, inRect);
		}	
	}
}
