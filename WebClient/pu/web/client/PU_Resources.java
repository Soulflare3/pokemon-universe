package pu.web.client;

import pu.web.client.resources.fonts.Fonts;
import pu.web.shared.ImageLoadEvent;

import com.google.gwt.dom.client.ImageElement;
import com.google.gwt.resources.client.ImageResource;
import com.googlecode.gwtgl.binding.WebGLTexture;

public class PU_Resources
{
	private PU_Font[] mFonts = new PU_Font[Fonts.FONT_COUNT];
	
	public PU_Resources()
	{
	}
	
	public void loadFonts()
	{
		loadFont(Fonts.FONT_PURITAN_BOLD_14, (ImageResource) Fonts.INSTANCE.puritanBold14Bitmap(), Fonts.INSTANCE.puritanBold14Info().getText());
	}
	
	public void loadFont(final int fontId, ImageResource imageResource, final String fontInfo)
	{
		final WebGLTexture texture = PUWeb.engine().createEmptyTexture();
		final ImageElement image = PUWeb.engine().getImageElement(imageResource);
		loadImage(new ImageLoadEvent() 
		{
			@Override
			public void loaded()
			{
				PUWeb.engine().fillTexture(texture, image);
				setFont(fontId, new PU_Font(texture, fontInfo));
				
				PUWeb.log("font loaded");
			}
		}, image);
	}
	
	public PU_Font getFont(int fontId)
	{
		return mFonts[fontId];
	}
	
	public native boolean imageLoaded(ImageElement image) /*-{
		return image.complete;
	}-*/;
	
	public native void loadImage(ImageLoadEvent callback, ImageElement image) /*-{
		var events = this;
		
		if(image.complete)
		{
			callback.@pu.web.shared.ImageLoadEvent::loaded()();
		}
		else
		{
			image.addEventListener("load", function(e) {
				callback.@pu.web.shared.ImageLoadEvent::loaded()();
			}, false);
		}
	}-*/;
	
	public void setFont(int fontId, PU_Font font)
	{
		mFonts[fontId] = font;
	}
}
