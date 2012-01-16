package pu.web.client.gui.impl;

import pu.web.client.PUWeb;
import pu.web.client.PU_Image;
import pu.web.client.PU_Rect;
import pu.web.client.gui.Panel;
import pu.web.client.resources.gui.GuiImages;

public class PU_WorldPanel extends Panel
{
	private PU_Image mImageBase = null;
	
	public PU_WorldPanel(int x, int y, int width, int height)
	{
		super(x, y, width, height);
		
		mImageBase = PUWeb.resources().getGuiImage(GuiImages.IMG_GUI_WORLD_BOTTOMBASE);
	}
	
	@Override
	public void draw(PU_Rect drawArea)
	{
		// Draw the base (fading white background, chat box, chat input box)
		mImageBase.draw(0, 568);
		
		super.draw(drawArea);
	}
}
