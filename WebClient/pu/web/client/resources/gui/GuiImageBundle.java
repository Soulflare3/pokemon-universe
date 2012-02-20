package pu.web.client.resources.gui;

import com.google.gwt.core.client.GWT;
import com.google.gwt.resources.client.ClientBundle;
import com.google.gwt.resources.client.ImageResource;

public interface GuiImageBundle extends ClientBundle
{
	public static int IMAGE_COUNT = 2;
	static GuiImageBundle INSTANCE = GWT.create(GuiImageBundle.class);
	
	@Source(value = { "loginBg.png" })
	ImageResource getLoginBg();
	
	@Source(value = { "chatpanel.png" })
	ImageResource getChatpanel();
}