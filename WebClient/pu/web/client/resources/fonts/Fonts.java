package pu.web.client.resources.fonts;

import com.google.gwt.core.client.GWT;
import com.google.gwt.resources.client.ClientBundle;
import com.google.gwt.resources.client.ImageResource;
import com.google.gwt.resources.client.TextResource;

public interface Fonts extends ClientBundle
{
	public static Fonts INSTANCE = GWT.create(Fonts.class);

	@Source(value = { "puritan_bold_14.fnt" })
	TextResource puritanBold14Info();
	
	@Source(value = { "puritan_bold_14.png" })
	ImageResource puritanBold14Bitmap();
}