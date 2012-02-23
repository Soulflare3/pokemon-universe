package pu.web.client.gui;

import pu.web.client.PU_Rect;

public class Scrollbar extends Container
{
	public static final int SCROLLBAR_VERTICAL = 0;
	public static final int SCROLLBAR_HORIZONTAL = 1;
	
	private int mValue = 0;
	private int mOrientation = SCROLLBAR_VERTICAL;
	private Panel pnlBackground = null;
	private Button btnLeftTop = null;
	private Button btnRightDown = null;
	private ScrollButton scbScroller = null;
	private int mMinValue = 0;
	private int mMaxValue = 100;
	private OnValueChangedListener mValueChangedListener = null;
	
	public Scrollbar(int x, int y, int width, int height, int orientation)
	{
		super(x, y, width, height);
		mOrientation = orientation;
		
		pnlBackground = new Panel(0, 0, width, height);
		addChild(pnlBackground);
		
		if(mOrientation == SCROLLBAR_VERTICAL)
		{
			btnLeftTop = new Button(0, 0, width, width, "");
			btnRightDown = new Button(0, height-width, width, width, "");
			scbScroller = new ScrollButton(0, width, width, width, new PU_Rect(0, width, width, height-(2*width)));
		}
		else
		{
			btnLeftTop = new Button(0, 0, height, height, "");
			btnRightDown = new Button(width-height, 0, height, height, "");
			scbScroller = new ScrollButton(height, 0, height, height, new PU_Rect(height, 0, height, width-(2*height)));
		}
		
		addChild(btnLeftTop);
		addChild(btnRightDown);
		addChild(scbScroller);
		
		scbScroller.setOnScrollChangeListener(new OnScrollChangeListener()
		{
			@Override
			public void onScrollChange(int x, int y)
			{
				scrollButtonChanged(x, y);				
			}
		});
		
		btnLeftTop.setOnClickListener(new OnClickListener()
		{
			@Override
			public void onClick(int x, int y)
			{
				if(mValue-1 >= mMinValue)
				{
					mValue--;
					updateScrollerPos();
					
					if(mValueChangedListener != null)
					{
						mValueChangedListener.onValueChanged(mValue);
					}
				}
			}
		});
		
		btnRightDown.setOnClickListener(new OnClickListener()
		{
			@Override
			public void onClick(int x, int y)
			{
				if(mValue+1 <= mMaxValue)
				{
					mValue++;
					updateScrollerPos();
					
					if(mValueChangedListener != null)
					{
						mValueChangedListener.onValueChanged(mValue);
					}
				}
			}
		});
	}
	
	public Button getButtonLeftTop()
	{
		return btnLeftTop;
	}
	
	public Button getButtonRightDown()
	{
		return btnRightDown;
	}
	
	public ScrollButton getScroller()
	{
		return scbScroller;
	}
	
	public void setMinValue(int value)
	{
		mMinValue = value;
		updateScrollerPos();
	}
	
	public int getMinValue()
	{
		return mMinValue;
	}
	
	public void setMaxValue(int value)
	{
		mMaxValue = value;
		updateScrollerPos();
	}
	
	public int getMaxValue()
	{
		return mMaxValue;
	}
	
	public void setValue(int value)
	{
		mValue = value;
		updateScrollerPos();
	}
	
	public int getValue()
	{
		return mValue;
	}
	
	public void setOnValueChangedListener(OnValueChangedListener listener)
	{
		mValueChangedListener = listener;
	}
	
	public void scrollButtonChanged(int x, int y)
	{
		mValue = mMinValue + (int)(((float)y / (float)getScrollAreaSize()) * (float)mMaxValue);
		
		if(mValueChangedListener != null)
		{
			mValueChangedListener.onValueChanged(mValue);
		}
	}
	
	public void updateScrollerPos()
	{
		int delta = mMaxValue - mMinValue;
		int size = 0;
		if(delta != 0) 
		{
			size = (int)((float)(mValue - mMinValue) * (float)(getScrollAreaSize()) / (float)delta);
		}

		if(mOrientation == SCROLLBAR_VERTICAL) 
		{
			size += btnLeftTop.getRect().y + btnLeftTop.getRect().height;
			scbScroller.getRect().y = size;
		} 
		else
		{
			size += btnRightDown.getRect().x + btnRightDown.getRect().width;
			scbScroller.getRect().x = size;
		}
	}
	
	private int getScrollAreaSize()
	{
		if(mOrientation == SCROLLBAR_VERTICAL)
		{
			return getRect().height-btnLeftTop.getRect().height-btnRightDown.getRect().height-scbScroller.getRect().height;
		}
		else
		{
			return getRect().width-btnLeftTop.getRect().width-btnRightDown.getRect().width-scbScroller.getRect().width;
		}
	}
}
