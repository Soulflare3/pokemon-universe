package pu.web.client.resources.gui;

import com.google.gwt.resources.client.ImageResource;

public class GuiImages
{
	private static ImageResource[] mImages;
	
	public static final int IMG_GUI_INTROBG = 0;
	public static final int IMG_GUI_WORLD_BOTTOMBASE = 1;
	public static final int IMG_GUI_WORLD_POKEMONBAR = 2;
	public static final int IMG_GUI_WORLD_POKEMONSLOT = 3;
	public static final int IMG_GUI_WORLD_HPBAR_EXP = 4;
	public static final int IMG_GUI_WORLD_HPBAR_GREEN = 5;
	public static final int IMG_GUI_WORLD_HPBAR_YELLOW = 6;
	public static final int IMG_GUI_WORLD_HPBAR_RED = 7;
	
	static {
		mImages = new ImageResource[] {
				GuiImageBundle.INSTANCE.getLoginBg(),
				GuiImageBundle.INSTANCE.getChatpanel(),
				GuiImageBundle.INSTANCE.getPokemonBar(),
				GuiImageBundle.INSTANCE.getPokemonSlot(),
				GuiImageBundle.INSTANCE.getHpBarExp(),
				GuiImageBundle.INSTANCE.getHpBarGreen(),
				GuiImageBundle.INSTANCE.getHpBarYellow(),
				GuiImageBundle.INSTANCE.getHpBarRed()
		};
	}
	
	public static ImageResource[] getImages()
	{
		return mImages;
	}
}
