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
}
