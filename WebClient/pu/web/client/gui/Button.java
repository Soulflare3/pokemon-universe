package pu.web.client.gui;

import pu.web.client.PUWeb;
import pu.web.client.PU_Font;
import pu.web.client.PU_Image;
import pu.web.client.PU_Rect;

public class Button extends TextElement
{
	private String mCaption = "";
	protected boolean mMouseDown = false;
	private boolean mHover = false;
	
	private ElementImage mImage = new ElementImage();
	private OnClickListener mClickListener = null;
	
	public Button(int x, int y, int width, int height, String caption)
	{
		super(x, y, width, height);
		mCaption = caption;
		setFont(PUWeb.gui().getDefaultFont());
	}
	
	public String getCaption()
	{
		return mCaption;
	}
	
	public void setCaption(String caption)
	{
		mCaption = caption;
	}
	
	public ElementImage getImage()
	{
		return mImage;
	}
	
	public void setOnClickListener(OnClickListener listener)
	{
		mClickListener = listener;
	}
	
	@Override
	public void draw(PU_Rect drawArea)
	{
		PU_Rect realRect = new PU_Rect(getRect().x + drawArea.x, getRect().y + drawArea.y, getRect().width, getRect().height);
		PU_Rect inRect = drawArea.intersection(realRect);
		
		PU_Image drawable = null;
		if(mMouseDown && mImage.getMouseDownImage() != null)
		{
			drawable = mImage.getMouseDownImage();
		}
		else if(mHover && mImage.getHoverImage() != null)
		{
			drawable = mImage.getHoverImage();
		}
		else
		{
			drawable = mImage.getImage();
		}
		
		if(drawable != null)
		{
			mImage.getImage().drawRectInRect(getRect(), drawArea);
		}
		
		PUWeb.engine().setColor(0, 0, 0, 255);
		PUWeb.engine().renderRect(realRect.x+2, realRect.y+2, realRect.width-4, realRect.height-4);
		
		PU_Font font = getFont();
		if(!mCaption.equals("") && font != null)
		{
			int captionX = getRect().x+((getRect().width/2)-(font.getStringWidth(mCaption)/2));
			int captionY = getRect().y+((getRect().height/2)-(font.getLineHeight()/2));
			font.setColor(getFontColor().r, getFontColor().g, getFontColor().b);
			font.drawTextInRect(mCaption, drawArea.x + captionX, drawArea.y + captionY, inRect);
		}
	}
	
	@Override
	public void mouseMove(int x, int y)
	{
		if(getRect().contains(x, y))
		{
			mHover = true;
		}
		else
		{
			mHover = false;
		}
	}
	
	@Override
	public void mouseDown(int x, int y)
	{
		if(getRect().contains(x, y))
		{
			mMouseDown = true;
		}
	}
	
	@Override
	public void mouseUp(int x, int y)
	{
		if(getRect().contains(x, y))
		{
			if(mMouseDown && this.mClickListener != null)
			{
				mClickListener.onClick(x, y);
			}
		}
		mMouseDown = false;
	}
}
