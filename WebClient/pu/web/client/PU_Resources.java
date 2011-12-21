package pu.web.client;

import pu.web.client.resources.fonts.Fonts;
import pu.web.client.resources.gui.GuiImageBundle;
import pu.web.shared.ImageLoadEvent;

import com.google.gwt.dom.client.ImageElement;
import com.google.gwt.resources.client.ImageResource;
import com.google.gwt.resources.client.ResourcePrototype;
import com.googlecode.gwtgl.binding.WebGLTexture;

public class PU_Resources
{
	private PU_Font[] mFonts = new PU_Font[Fonts.FONT_COUNT];
	private PU_Image[] mGuiImages = null;
	
	private int mFontCount = 0;
	private int mFontCountLoaded = 0;
	
	private int mGuiImageCount = 0;
	private int mGuiImageCountLoaded = 0;
	
	public PU_Resources()
	{
		mFontCount = Fonts.FONT_COUNT;
		mGuiImageCount = GuiImageBundle.INSTANCE.getResources().length;
	}
	
	public void checkComplete()
	{
		boolean complete = true;
		
		if(mFontCount != mFontCountLoaded)
			complete = false;
		
		if(mGuiImageCount != mGuiImageCountLoaded)
			complete = false;
		
		if(complete)
			PUWeb.resourcesLoaded();
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
			
			image.addEventListener("error", function(e) {
				callback.@pu.web.shared.ImageLoadEvent::error()();
			}, false);
		}
	}-*/;
	
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
				
				mFontCountLoaded++;
				checkComplete();
			}

			@Override
			public void error()
			{
				mFontCountLoaded++;
				checkComplete();
				
			}
		}, image);
	}
	
	public PU_Font getFont(int fontId)
	{
		return mFonts[fontId];
	}
	
	public void setFont(int fontId, PU_Font font)
	{
		mFonts[fontId] = font;
	}
	
	public void loadGuiImages()
	{
		final ResourcePrototype[] resources = GuiImageBundle.INSTANCE.getResources();
		mGuiImages = new PU_Image[resources.length];
		for(ResourcePrototype resource : resources)
		{
			String name = resource.getName();
			final int id = Integer.parseInt(name.replace("res_", ""));
			
			final WebGLTexture texture = PUWeb.engine().createEmptyTexture();
			final ImageElement image = PUWeb.engine().getImageElement((ImageResource)resource);
			loadImage(new ImageLoadEvent() 
			{
				@Override
				public void loaded()
				{
					PUWeb.engine().fillTexture(texture, image);
					if(id >= 0 && id < resources.length)
					{
						mGuiImages[id] = new PU_Image(image.getWidth(), image.getHeight(), texture);
					}
					
					mGuiImageCountLoaded++;
					checkComplete();
				}

				@Override
				public void error()
				{
					mGuiImageCountLoaded++;
					checkComplete();
				}
			}, image);
		}
	}
	
	public PU_Image getGuiImage(int id)
	{
		if(mGuiImages != null && id >= 0 && id < mGuiImages.length)
		{
			return mGuiImages[id];
		}
		return null;
	}
}
