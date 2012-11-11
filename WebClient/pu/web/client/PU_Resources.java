package pu.web.client;

import java.util.HashMap;

import pu.web.client.resources.fonts.Fonts;
import pu.web.client.resources.gui.GuiImages;
import pu.web.client.resources.pokemon.Pokemon;
import pu.web.client.resources.tiles.Tiles;
import pu.web.shared.ImageLoadEvent;

import com.google.gwt.dom.client.ImageElement;
import com.google.gwt.resources.client.ImageResource;
import com.google.gwt.xml.client.Document;
import com.google.gwt.xml.client.Element;
import com.google.gwt.xml.client.NodeList;
import com.google.gwt.xml.client.XMLParser;
import com.googlecode.gwtgl.binding.WebGLTexture;

public class PU_Resources
{
	private PU_Font[] mFonts = new PU_Font[Fonts.FONT_COUNT];
	private PU_Image[] mGuiImages = null;
	private HashMap<Integer, PU_Image> mTiles = new HashMap<Integer, PU_Image>();
	private HashMap<Long, PU_Image> mCreatureImages = new HashMap<Long, PU_Image>();
	private HashMap<Integer, PU_Image> mPokeImage_Front = new HashMap<Integer, PU_Image>();
	private HashMap<Integer, PU_Image> mPokeImage_Back = new HashMap<Integer, PU_Image>();
	private HashMap<Integer, PU_Image> mPokeImage_Icon = new HashMap<Integer, PU_Image>();
	
	private int mFontCount = 0;
	private int mFontCountLoaded = 0;
	
	private int mGuiImageCount = 0;
	private int mGuiImageCountLoaded = 0;
	
	private int mSpriteCount = 0;
	private int mSpriteCountLoaded = 0;
	
	private int mPokemonCount = 0;
	private int mPokemonCountLoaded = 0;
	
	private WebGLTexture mSpriteTexture = null;
	private WebGLTexture mPokemonTexture = null;
	
	private boolean mResourcesLoaded = false;
	
	public PU_Resources()
	{
		mFontCount = Fonts.FONT_COUNT;
		mGuiImageCount = GuiImages.getImages().length;
	}
	
	public void checkComplete()
	{
		if(!mResourcesLoaded)
		{
			boolean complete = true;
			
			if(mFontCount <= 0 || mFontCount != mFontCountLoaded)
			{
				complete = false;
			}
			
			if(mGuiImageCount <= 0 || mGuiImageCount != mGuiImageCountLoaded)
			{
				complete = false;
			}
			
			if(mSpriteCount <= 0 || mSpriteCount != mSpriteCountLoaded)
			{
				complete = false;
			}
			
			if(mPokemonCount <= 0 || mPokemonCount != mPokemonCountLoaded)
			{
				complete = false;
			}
			
			if(complete)
			{
				mResourcesLoaded = true;
				PUWeb.resourcesLoaded();
			}	
		}
	}
	
	public float getLoadProgress()
	{
		float progress = 0.0f;
		
		// Font progress (10%)
		progress += ((float)mFontCountLoaded/(float)mFontCount)*0.1f;
		
		// GUI progress (30%)
		progress += ((float)mFontCountLoaded/(float)mFontCount)*0.3f;
		
		// Sprites progress (30%)
		if(mSpriteCount > 0)
		{
			progress += ((float)mSpriteCountLoaded/(float)mSpriteCount)*0.3f;	
		}
		
		// Pokemon progress (30%)
		if(mPokemonCount > 0)
		{
			progress += ((float)mPokemonCountLoaded/(float)mPokemonCount)*0.3f;	
		}
		
		// Check if loading is complete
		checkComplete();
		
		return progress;
	}
	
	public int getFontLoadProgress()
	{
		return (int)((float)((float)mFontCountLoaded/(float)mFontCount)*100.0);
	}
	
	public int getGuiImageLoadProgress()
	{
		return (int)((float)((float)mGuiImageCountLoaded/(float)mGuiImageCount)*100.0);
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
		loadFont(Fonts.FONT_ARIALBLK_BOLD_14, (ImageResource) Fonts.INSTANCE.puritanBold14Bitmap(), Fonts.INSTANCE.puritanBold14Info().getText());
		loadFont(Fonts.FONT_ARIALBLK_BOLD_14_OUTLINE, (ImageResource) Fonts.INSTANCE.arialBlk14OutlineBitmap(), Fonts.INSTANCE.arialBlk14OutlineInfo().getText());
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
			}

			@Override
			public void error()
			{
				mFontCountLoaded++;
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
		final ImageResource[] resources = GuiImages.getImages();
		mGuiImages = new PU_Image[resources.length];
		for(int i = 0; i < resources.length; i++)
		{
			final int id = i;
			
			final WebGLTexture texture = PUWeb.engine().createEmptyTexture();
			final ImageElement image = PUWeb.engine().getImageElement(resources[i]);
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
				}

				@Override
				public void error()
				{
					mGuiImageCountLoaded++;
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
	
	public void loadSprites()
	{
		final ImageResource imageResource = Tiles.INSTANCE.getTilesBitmap();
		final String imageInfo = Tiles.INSTANCE.getTilesInfo().getText();
	
		mSpriteTexture = PUWeb.engine().createEmptyTexture();
		final ImageElement image = PUWeb.engine().getImageElement(imageResource);
		loadImage(new ImageLoadEvent() 
		{
			@Override
			public void loaded()
			{
				PUWeb.engine().fillTexture(mSpriteTexture, image);
				
				Document infoDom = XMLParser.parse(imageInfo);
				
				NodeList sprites = infoDom.getElementsByTagName("sprite");
				PU_Resources.this.mSpriteCount = sprites.getLength();
				for(int i = 0; i < sprites.getLength(); i++)
				{
					Element element = (Element) sprites.item(i);
					
					String name = element.getAttribute("n");
					
					PU_Rect texCoords = new PU_Rect();
					texCoords.x = Integer.parseInt(element.getAttribute("x"));
					texCoords.y = Integer.parseInt(element.getAttribute("y"));
					texCoords.width = Integer.parseInt(element.getAttribute("w"));
					texCoords.height = Integer.parseInt(element.getAttribute("h"));
					
					int offsetX = 0;
					if(element.hasAttribute("oX"))
						offsetX = Integer.parseInt(element.getAttribute("oX"));
					
					int offsetY = 0;
					if(element.hasAttribute("oY"))
						offsetY = Integer.parseInt(element.getAttribute("oY"));
					
					int width = texCoords.width;
					if(element.hasAttribute("oW"))
						width = Integer.parseInt(element.getAttribute("oW"));
					
					int height = texCoords.height;
					if(element.hasAttribute("oH"))
						height = Integer.parseInt(element.getAttribute("oH"));
					
					PU_Image spriteImage = new PU_Image(width, height, null);
					spriteImage.setTextureCoords(texCoords, image.getWidth(), image.getHeight());
					spriteImage.setOffsetX(offsetX);
					spriteImage.setOffsetY(offsetY);
					if(name.contains("creatures/"))
					{
						// Creature sprite
						parseCreatureSprite(name, spriteImage);
					}
					else
					{
						// Tile sprite
						parseTileSprite(name, spriteImage);
					}
					PU_Resources.this.mSpriteCountLoaded++;
				}
			}

			@Override
			public void error()
			{
				PUWeb.log("Error loading sprites");
			}
		}, image);
	}
	
	public WebGLTexture getSpriteTexture()
	{
		return mSpriteTexture;
	}
	
	public void parseTileSprite(String name, PU_Image image)
	{
		int id = Integer.parseInt(name);
		mTiles.put(id, image);
	}

	public void parseCreatureSprite(String name, PU_Image image)
	{
		String ids = name.replace("creatures/", "");
		String[] parts = ids.split("_");
		long bodypart = Long.parseLong(parts[0]);
		long id = Long.parseLong(parts[1]);
		long dir = Long.parseLong(parts[2]);
		long frame = Long.parseLong(parts[3]);
		
		long key = ((bodypart) | (id << 8) | (dir << 16) | (frame << 24));
		mCreatureImages.put(key, image);
	}
	
	public PU_Image getTileImage(int id)
	{
		return mTiles.get(id);
	}
	
	public PU_Image getCreatureImage(int bodypart, int id, int dir, int frame)
	{
		long key = ((bodypart) | (id << 8) | (dir << 16) | (frame << 24));
		return mCreatureImages.get(key);
	}
	
	public void loadPokemonImages()
	{
		final ImageResource imageResource = Pokemon.INSTANCE.getPokemonBitmap();
		final String imageInfo = Pokemon.INSTANCE.getPokemonInfo().getText();
	
		mPokemonTexture = PUWeb.engine().createEmptyTexture();
		final ImageElement image = PUWeb.engine().getImageElement(imageResource);
		loadImage(new ImageLoadEvent() 
		{
			@Override
			public void loaded()
			{
				PUWeb.engine().fillTexture(mPokemonTexture, image);
				
				Document infoDom = XMLParser.parse(imageInfo);
				
				NodeList sprites = infoDom.getElementsByTagName("sprite");
				PU_Resources.this.mPokemonCount = sprites.getLength();
				for(int i = 0; i < sprites.getLength(); i++)
				{
					Element element = (Element) sprites.item(i);
					
					String name = element.getAttribute("n");
					
					PU_Rect texCoords = new PU_Rect();
					texCoords.x = Integer.parseInt(element.getAttribute("x"));
					texCoords.y = Integer.parseInt(element.getAttribute("y"));
					texCoords.width = Integer.parseInt(element.getAttribute("w"));
					texCoords.height = Integer.parseInt(element.getAttribute("h"));
					
					int offsetX = 0;
					if(element.hasAttribute("oX"))
						offsetX = Integer.parseInt(element.getAttribute("oX"));
					
					int offsetY = 0;
					if(element.hasAttribute("oY"))
						offsetY = Integer.parseInt(element.getAttribute("oY"));
					
					int width = texCoords.width;
					if(element.hasAttribute("oW"))
						width = Integer.parseInt(element.getAttribute("oW"));
					
					int height = texCoords.height;
					if(element.hasAttribute("oH"))
						height = Integer.parseInt(element.getAttribute("oH"));
					
					PU_Image spriteImage = new PU_Image(width, height, null);
					spriteImage.setTextureCoords(texCoords, image.getWidth(), image.getHeight());
					spriteImage.setOffsetX(offsetX);
					spriteImage.setOffsetY(offsetY);
					
					if(name.contains("back/"))
					{
						int id = Integer.parseInt(name.replace("back/", ""));
						mPokeImage_Back.put(id, spriteImage);
					}
					else if(name.contains("front/"))
					{
						int id = Integer.parseInt(name.replace("front/", ""));
						mPokeImage_Front.put(id, spriteImage);
					}
					else if(name.contains("icon/"))
					{
						int id = Integer.parseInt(name.replace("icon/", ""));
						mPokeImage_Icon.put(id, spriteImage);
					}	
					
					PU_Resources.this.mPokemonCountLoaded++;
				}
			}

			@Override
			public void error()
			{
				PUWeb.log("Error loading sprites");
			}
		}, image);
	}
	
	public WebGLTexture getPokemonTexture()
	{
		return mPokemonTexture;
	}
	
	public PU_Image getPokemonFront(int id)
	{
		return mPokeImage_Front.get(id);
	}
	
	public PU_Image getPokemonBack(int id)
	{
		return mPokeImage_Back.get(id);
	}
	
	public PU_Image getPokemonIcon(int id)
	{
		return mPokeImage_Icon.get(id);
	}
}
