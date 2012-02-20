package pu.web.client.resources.gui;

import com.google.gwt.resources.client.ImageResource;

public class GuiImages
{
	private static ImageResource[] mImages;
	
	public static final int IMG_GUI_INTROBG = 0;
	public static final int IMG_GUI_WORLD_BOTTOMBASE = 1;
	
	static {
		mImages = new ImageResource[] {
				GuiImageBundle.INSTANCE.getLoginBg(),
				GuiImageBundle.INSTANCE.getChatpanel()
		};
	}
	
	public static ImageResource[] getImages()
	{
		return mImages;
	}
}
