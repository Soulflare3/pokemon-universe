package pu.web.client.gui;

import pu.web.client.PU_Image;

public class ElementImage
{
	private PU_Image mImage = null;
	private PU_Image mHoverImage = null;
	private PU_Image mMouseDownImage = null;
	
	public ElementImage()
	{
		
	}
	
	public PU_Image getImage()
	{
		return mImage;
	}
	
	public void setImage(PU_Image image)
	{
		mImage = image;
	}
	
	public PU_Image getHoverImage()
	{
		return mHoverImage;
	}
	
	public void setHoverImage(PU_Image image)
	{
		mHoverImage = image;
	}
	
	public PU_Image getMouseDownImage()
	{
		return mMouseDownImage;
	}
	
	public void setMouseDownImage(PU_Image image)
	{
		mImage = image;
	}
}