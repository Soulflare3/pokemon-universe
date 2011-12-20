package pu.web.client;

import com.googlecode.gwtgl.binding.WebGLTexture;

public class PU_Image
{
	private int mWidth;
	private int mHeight;
	WebGLTexture mTexture;
	private int mBlendMode = PU_Engine.BLENDMODE_BLEND;
	private PU_Color mColorMod = new PU_Color();
	
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
		//PUWeb.log("rect 1: " + rect.x + " " + rect.y + " " + rect.width + " " + rect.height);
		//PUWeb.log("rect 2: " + clip.x + " " + clip.y + " " + clip.width + " " + clip.height);
		//PUWeb.log("rect 3: " + inRect.x + " " + inRect.y + " " + inRect.width + " " + inRect.height);
		inRect = inRect.intersection(new PU_Rect(rect.x, rect.y, clip.width, clip.height));
		//PUWeb.log("rect 4: " + inRect.x + " " + inRect.y + " " + inRect.width + " " + inRect.height);
		PU_Rect dstRect = new PU_Rect(inRect.x, inRect.y, inRect.width, inRect.height);
		inRect.x -= rect.x;
		inRect.y -= rect.y;
		inRect.x += clip.x;
		inRect.y += clip.y;
		
		PUWeb.engine().renderTexture(this, inRect, dstRect);
	}
}
