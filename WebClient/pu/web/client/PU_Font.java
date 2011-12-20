package pu.web.client;

import com.google.gwt.xml.client.Document;
import com.google.gwt.xml.client.Element;
import com.google.gwt.xml.client.NodeList;
import com.google.gwt.xml.client.XMLParser;
import com.googlecode.gwtgl.binding.WebGLTexture;

public class PU_Font
{
	private PU_Image mImage;
	private int mLineHeight =  0;
	private PU_FontCharacter[] mCharacters = new PU_FontCharacter[256];
	private PU_Color mColor = new PU_Color(255, 255, 255, 255);
	
	public PU_Font(WebGLTexture texture, String info)
	{
		mImage = new PU_Image(256, 256, texture);
		
		Document infoDom = XMLParser.parse(info);
		Element common = (Element)infoDom.getElementsByTagName("common").item(0);
		mLineHeight = Integer.parseInt(common.getAttribute("lineHeight"));
		
		NodeList chars = infoDom.getElementsByTagName("char");
		for(int i = 0; i < chars.getLength(); i++)
		{
			Element character = (Element) chars.item(i);
			
			PU_FontCharacter fontCharacter = new PU_FontCharacter();
			int id = Integer.parseInt(character.getAttribute("id"));
			fontCharacter.x = Integer.parseInt(character.getAttribute("x"));
			fontCharacter.y = Integer.parseInt(character.getAttribute("y"));
			fontCharacter.width = Integer.parseInt(character.getAttribute("width"));
			fontCharacter.height = Integer.parseInt(character.getAttribute("height"));
			fontCharacter.xOffset = Integer.parseInt(character.getAttribute("xoffset"));
			fontCharacter.yOffset = Integer.parseInt(character.getAttribute("yoffset"));
			fontCharacter.xAdvance = Integer.parseInt(character.getAttribute("xadvance"));
			mCharacters[id] = fontCharacter;
		}
	}
	
	public void setColor(int red, int green, int blue)
	{
		mColor.r = red;
		mColor.g = green;
		mColor.b = blue;
	}
	
	public PU_FontCharacter getCharacter(int id)
	{
		return mCharacters[id];
	}
	
	public int getLineHeight()
	{
		return mLineHeight;
	}
	
	public PU_Image getImage()
	{
		return mImage;
	}
	
	public int getStringWidth(String text)
	{
		int width = 0;
		for(int i = 0; i < text.length(); i++)
		{
			int id = text.charAt(i);
			PU_FontCharacter character = mCharacters[id];
			if(character != null)
			{
				width += character.xAdvance;
			}
		}
		return width;
	}
	
	public void drawText(String text, int x, int y)
	{
		int drawX = x;
		int drawY = y;
		mImage.setColorMod(mColor.r, mColor.g, mColor.b);
		for(int i = 0; i < text.length(); i++)
		{
			int id = text.charAt(i);
			PU_FontCharacter character = mCharacters[id];
			if(character != null)
			{
				PU_Rect srcRect = new PU_Rect(character.x, character.y, character.width, character.height);
				PU_Rect dstRect = new PU_Rect(drawX+character.xOffset, drawY+character.yOffset, character.width, character.height);
				mImage.drawRectClip(dstRect, srcRect);
				
				drawX += character.xAdvance;
			}
		}
	}
	
	public void drawTextInRect(String text, int x, int y, PU_Rect rect)
	{
		int drawX = x;
		int drawY = y;
		mImage.setColorMod(mColor.r, mColor.g, mColor.b);
		for(int i = 0; i < text.length(); i++)
		{
			int id = text.charAt(i);
			PU_FontCharacter character = mCharacters[id];
			if(character != null)
			{
				PU_Rect srcRect = new PU_Rect(character.x, character.y, character.width, character.height);
				PU_Rect dstRect = new PU_Rect(drawX+character.xOffset, drawY+character.yOffset, character.width, character.height);
				mImage.drawRectClipInRect(dstRect, srcRect, rect);
				
				drawX += character.xAdvance;
			}
		}
	}
	
	public void drawBorderedText(String text, int x, int y)
	{
		int red = mColor.r;
		int green = mColor.g;
		int blue = mColor.b;
		setColor(0, 0, 0);
		drawText(text, x-1, y-1);
		drawText(text, x+1, y-1);
		drawText(text, x-1, y+1);
		drawText(text, x+1, y+1);
		setColor(red, green, blue);
		drawText(text, x, y);
	}
}
