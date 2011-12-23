package pu.web.client.gui;

import pu.web.client.PUWeb;
import pu.web.client.PU_Font;
import pu.web.client.PU_Rect;

public class TextField extends TextElement
{
	public static final int CARET_TIME = 500;
	
	private ElementColor mBackgroundColor = null;
	private ElementColor mBorderColor = null;
	private ElementImage mImage = new ElementImage();
	
	private OnKeyDownListener mKeyDownListener = null;
	
	private String mText = "";
	private boolean mReadOnly = false;
	private boolean mPassword = false;
	private boolean mCaret = false;
	private long mCaretLast = 0;
	
	public TextField(int x, int y, int width, int height)
	{
		super(x, y, width, height);
		
		setFocusable(true);
		setFont(PUWeb.gui().getDefaultFont());
		
		mCaretLast = System.currentTimeMillis();
	}
	
	public String getText()
	{
		return mText;
	}
	
	public void setText(String text)
	{
		mText = text;
	}
	
	public boolean isReadOnly()
	{
		return mReadOnly;
	}
	
	public void setReadOnly(boolean readonly)
	{
		mReadOnly = readonly;
	}
	
	public boolean isPassword()
	{
		return mPassword;
	}
	
	public void setPassword(boolean password)
	{
		mPassword = password;
	}
	
	public void setBackgroundColor(int red, int green, int blue)
	{
		if(mBackgroundColor == null)
			mBackgroundColor = new ElementColor();
		
		mBackgroundColor.setColor(red, green, blue);
	}
	
	public void setBorderColor(int red, int green, int blue)
	{
		if(mBorderColor == null)
			mBorderColor = new ElementColor();
		
		mBorderColor.setColor(red, green, blue);
	}
	
	public void setOnKeyDownListener(OnKeyDownListener listener)
	{
		mKeyDownListener = listener;
	}
	
	@Override
	public void keyDown(int button)
	{
		if(!mReadOnly && hasFocus())
		{
			if(button == 8/*TODO: backspace keycode*/)
			{
				if(mText.length() == 1)
				{
					mText = "";
				}
				else if(mText.length() > 1)
				{
					mText = mText.substring(0, mText.length()-1);
				}
			}
		}
		
		if(hasFocus() && mKeyDownListener != null)
		{
			mKeyDownListener.OnKeyDown(button);
		}
	}
	
	@Override
	public void textInput(int charCode)
	{
		if(!mReadOnly && hasFocus())
		{
			if(charCode != 0 && charCode > 31)
			{
				mText += (char) charCode;
			}
		}
	}
	
	@Override
	public void draw(PU_Rect drawArea)
	{
		PU_Rect realRect = new PU_Rect(getRect().x + drawArea.x, getRect().y + drawArea.y, getRect().width, getRect().height);
		PU_Rect inRect = drawArea.intersection(realRect);
		
		if(mBackgroundColor != null)
		{
			PUWeb.engine().setColor(mBackgroundColor.getColor().r, mBackgroundColor.getColor().g, mBackgroundColor.getColor().b, mBackgroundColor.getColor().a);
			PUWeb.engine().renderFillRect(inRect.x, inRect.y, inRect.width, inRect.height);
		}
		
		if(mBorderColor !=  null)
		{
			PUWeb.engine().setColor(mBorderColor.getColor().r, mBorderColor.getColor().g, mBorderColor.getColor().b, mBorderColor.getColor().a);
			PUWeb.engine().renderRect(inRect.x, inRect.y, inRect.width, inRect.height);
		}
		
		if(mImage.getImage() != null)
		{
			mImage.getImage().drawRectInRect(getRect(), drawArea);
		}
		
		PU_Font font = getFont();
		int caretX = 0;
		int textX = getRect().x + font.getStringWidth(" ");
		int textY = getRect().y + ((getRect().height/2)-(font.getLineHeight()/2));
		if(!mText.equals(""))
		{
			String drawText = "";
			if(mPassword) 
			{
				for(int i = 0; i < mText.length(); i++)
				{
					drawText += "*";
				}
			}
			else
			{
				drawText = mText;
			}
			
			font.setColor(getFontColor().r, getFontColor().g, getFontColor().b);
			font.drawTextInRect(drawText, drawArea.x + textX, drawArea.y + textY, inRect);
			
			caretX = font.getStringWidth(drawText);
		}
		
		//draw caret
		if(mCaret && !mReadOnly && hasFocus()) 
		{
			caretX += textX;
			if(inRect.contains(drawArea.x+caretX, drawArea.y+textY)) 
			{
				PUWeb.engine().setColor(getFontColor().r, getFontColor().g, getFontColor().b, 255);
				PUWeb.engine().renderLine(drawArea.x+caretX, drawArea.y+textY, drawArea.x+caretX, drawArea.y+textY+font.getLineHeight());
			}
		}
		
		if(System.currentTimeMillis()-mCaretLast >= CARET_TIME) 
		{
			mCaret = !mCaret;
			mCaretLast = System.currentTimeMillis();
		}
	}
	
	@Override
	public void mouseUp(int x, int y)
	{
		if(getRect().contains(x, y))
		{
			PUWeb.log("herpderp");
			Window window = getWindow();
			if(window != null)
			{
				window.focusElement(this);
			}
		}
	}
	
}