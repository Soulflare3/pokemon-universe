package pu.web.client;

import com.googlecode.gwtgl.binding.WebGLTexture;

public class PU_Image
{
	private int mWidth;
	private int mHeight;
	WebGLTexture mTexture;
	private int mBlendMode = PU_Engine.BLENDMODE_BLEND;
	private PU_Color mColorMod = new PU_Color();
	private PU_Rect mTextureCoords = null;
	private int mTextureWidth = 0;
	private int mTextureHeight = 0;
	private int mOffsetX = 0;
	private int mOffsetY = 0;
	
	public PU_Image(int width, int height, WebGLTexture texture)
	{
		mWidth = width;
		mHeight = height;
		
		mTexture = texture;
	}
	
	public int getWidth()
	{
		return mWidth;
	}
	
	public int getHeight()
	{
		return mHeight;
	}
	
	public WebGLTexture getTexture()
	{
		return mTexture;
	}
	
	public void setBlendMode(int blendMode)
	{
		mBlendMode = blendMode;
	}
	
	public int getBlendMode()
	{
		return mBlendMode;
	}
	
	public PU_Color getColor()
	{
		return mColorMod;
	}
	
	public void setAlphaMod(int alpha)
	{
		mColorMod.a = alpha;
	}
	
	public void setColorMod(int red, int green, int blue)
	{
		mColorMod.r = red;
		mColorMod.g = green;
		mColorMod.b = blue;
	}
	
	public PU_Rect getTextureCoords()
	{
		return mTextureCoords;
	}
	
	public void setTextureCoords(PU_Rect coords, int textureWidth, int textureHeight)
	{
		mTextureWidth = textureWidth;
		mTextureHeight = textureHeight;
		mTextureCoords = coords;
	}
	
	public int getTextureWidth()
	{
		return mTextureWidth;
	}
	
	public int getTextureHeight()
	{
		return mTextureHeight;
	}
	
	public int getOffsetX()
	{
		return mOffsetX;
	}
	
	public void setOffsetX(int offset)
	{
		mOffsetX = offset;
	}
	
	public int getOffsetY()
	{
		return mOffsetY;
	}
	
	public void setOffsetY(int offset)
	{
		mOffsetY = offset;
	}
	
	public void draw(int x, int y)
	{
		PU_Rect src = new PU_Rect(0, 0, mWidth, mHeight);
		PU_Rect dst = new PU_Rect(x, y, mWidth, mHeight);
		
		PUWeb.engine().renderTexture(this, src, dst);
	}
	
	public void drawRect(PU_Rect rect)
	{
		PU_Rect src = new PU_Rect(0, 0, mWidth, mHeight);
		
		PUWeb.engine().renderTexture(this, src, rect);
	}
	
	public void drawClip(int x, int y, PU_Rect clip)
	{
		PU_Rect dst = new PU_Rect(x, y, mWidth, mHeight);
		
		PUWeb.engine().renderTexture(this, clip, dst);
	}
	
	public void drawRectClip(PU_Rect rect, PU_Rect clip)
	{
		PUWeb.engine().renderTexture(this, clip, rect);
	}
	
	public void drawInRect(int x, int y, PU_Rect inRect)
	{
		PU_Rect imgRect = new PU_Rect(x, y, mWidth, mHeight);
		inRect = inRect.intersection(imgRect);
		PU_Rect dstRect = new PU_Rect(inRect.x, inRect.y, inRect.width, inRect.height);
		inRect.x -= x;
		inRect.y -= y;
		
		PUWeb.engine().renderTexture(this, inRect, dstRect);
	}
	
	public void drawRectInRect(PU_Rect rect, PU_Rect inRect)
	{
		inRect = inRect.intersection(rect);
		PU_Rect dstRect = new PU_Rect(inRect.x, inRect.y, inRect.width, inRect.height);
		inRect.x -= rect.x;
		inRect.y -= rect.y;
		
		PUWeb.engine().renderTexture(this, inRect, dstRect);
	}
	
	public void drawRectClipInRect(PU_Rect rect, PU_Rect clip, PU_Rect inRect)
	{
		inRect = inRect.intersection(new PU_Rect(rect.x, rect.y, clip.width, clip.height));
		PU_Rect dstRect = new PU_Rect(inRect.x, inRect.y, inRect.width, inRect.height);
		inRect.x -= rect.x;
		inRect.y -= rect.y;
		inRect.x += clip.x;
		inRect.y += clip.y;
		
		PUWeb.engine().renderTexture(this, inRect, dstRect);
	}
}
