package pu.web.client.resources.fonts;

import com.google.gwt.core.client.GWT;
import com.google.gwt.resources.client.ClientBundle;
import com.google.gwt.resources.client.ImageResource;
import com.google.gwt.resources.client.TextResource;

public interface Fonts extends ClientBundle
{
	public static int FONT_COUNT = 2;
	public static Fonts INSTANCE = GWT.create(Fonts.class);

	public static int FONT_ARIALBLK_BOLD_14 = 0;
	@Source(value = { "arialblk_bold_14.fnt" })
	TextResource puritanBold14Info();
	
	@Source(value = { "arialblk_bold_14.png" })
	ImageResource puritanBold14Bitmap();
	
	public static int FONT_ARIALBLK_BOLD_14_OUTLINE = 1;
	@Source(value = { "arialblk_bold_14_outline.fnt" })
	TextResource arialBlk14OutlineInfo();
	
	@Source(value = { "arialblk_bold_14_outline.png" })
	ImageResource arialBlk14OutlineBitmap();
}