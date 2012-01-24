package pu.web.client.gui.impl;

import pu.web.client.PUWeb;
import pu.web.client.PU_Image;
import pu.web.client.PU_Rect;
import pu.web.client.gui.OnKeyDownListener;
import pu.web.client.gui.Panel;
import pu.web.client.gui.Scrollbar;
import pu.web.client.gui.TextField;
import pu.web.client.resources.fonts.Fonts;
import pu.web.client.resources.gui.GuiImages;

public class PU_WorldPanel extends Panel
{
	private PU_Image mImageBase = null;
	
	// Controls
	private PU_Chatbox mChatbox = null;
	private Scrollbar mChatboxScrollbar = null;
	private TextField mChatInput = null;
	
	public PU_WorldPanel(int x, int y, int width, int height)
	{
		super(x, y, width, height);
		
		mImageBase = PUWeb.resources().getGuiImage(GuiImages.IMG_GUI_WORLD_BOTTOMBASE);
		
		mChatbox = new PU_Chatbox(13, 571, 350, 110);
		addChild(mChatbox);
		
		mChatboxScrollbar = new Scrollbar(368, 573, 20, 105, Scrollbar.SCROLLBAR_VERTICAL);
		mChatbox.setScrollbar(mChatboxScrollbar);
		addChild(mChatboxScrollbar);
		
		mChatInput = new TextField(10, 694, 349, 18);
		mChatInput.setFontColor(0, 0, 0);
		mChatInput.setOnKeyDownListener(mChatInputKeydownListener);
		PUWeb.gui().getRoot().focusElement(mChatInput);
		addChild(mChatInput);
	}
	
	private OnKeyDownListener mChatInputKeydownListener = new OnKeyDownListener()
	{	
		@Override
		public void OnKeyDown(int button)
		{
			if(button == 13)
			{				
				String message = mChatInput.getText();
				
				PU_Text test = new PU_Text(PUWeb.resources().getFont(Fonts.FONT_ARIALBLK_BOLD_14));
				test.add("Urmel: ", 255, 100 ,100);
				test.add(message, 100, 100, 255);
				mChatbox.addText(test);
				
				mChatInput.setText("");
			}
		}
	};
	
	@Override
	public void draw(PU_Rect drawArea)
	{
		// Draw the base (fading white background, chat box, chat input box)
		mImageBase.draw(0, 568);
		
		super.draw(drawArea);
	}
}
