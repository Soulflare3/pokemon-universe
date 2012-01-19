package pu.web.client.gui;

import pu.web.client.PU_Rect;

public class ScrollButton extends Button
{
	private PU_Rect mBoundaries = null;
	private boolean mScrolling = false;
	private int mStartX;
	private int mStartY;
	private int mMouseDownX;
	private int mMouseDownY;
	private OnScrollChangeListener mScrollChangeListener = null;
	
	public ScrollButton(int x, int y, int width, int height, PU_Rect boundaries)
	{
		super(x, y, width, height, "");
		mBoundaries = boundaries;
	}
	
	public PU_Rect getBoundaries()
	{
		return mBoundaries;
	}
	
	public void setBoundaries(PU_Rect boundaries)
	{
		mBoundaries = boundaries;
	}
	
	public void setOnScrollChangeListener(OnScrollChangeListener listener)
	{
		mScrollChangeListener = listener;
	}
	
	private void updateScrollChangeListener()
	{
		if(mScrollChangeListener != null)
		{
			mScrollChangeListener.onScrollChange(getScrolledX(), getScrolledY());
		}
	}
	
	@Override
	public void mouseDown(int x, int y)
	{
		mStartX = getRect().x;
		mStartY = getRect().y;
		
		mMouseDownX = x;
		mMouseDownY = y;
		
		super.mouseDown(x, y);
	}
	
	@Override
	public void mouseUp(int x, int y)
	{
		if(!mScrolling && !mMouseDown && mBoundaries.contains(x, y))
		{
			int newX = (mBoundaries.x + (x - mBoundaries.x)) - (getRect().width / 2);
			if(newX >= mBoundaries.x && newX+getRect().width <= mBoundaries.x+mBoundaries.width)
			{
				getRect().x = newX;
			}
			
			int newY = (mBoundaries.y + (y - mBoundaries.y)) - (getRect().height / 2);
			if(newY >= mBoundaries.y && newY+getRect().height <= mBoundaries.y+mBoundaries.height) 
			{
				getRect().y = newY;
			}
			
			updateScrollChangeListener();
		}
		mScrolling = false;
		super.mouseUp(x, y);
	}
	
	@Override
	public void mouseMove(int x, int y)
	{
		if(mMouseDown) 
		{
			mScrolling = true;

			int newX = mStartX + (x - mMouseDownX);
			if(newX >= mBoundaries.x && newX+getRect().width <= mBoundaries.x+mBoundaries.width)
			{
				getRect().x = newX;
			} 
			else if(newX < mBoundaries.x) 
			{
				getRect().x = mBoundaries.x;
			} 
			else if(newX > mBoundaries.x+mBoundaries.width) 
			{
				getRect().x = mBoundaries.width - getRect().width;
			}

			int newY = mStartY + (y - mMouseDownY);
			if(newY >= mBoundaries.y && newY+getRect().height <= mBoundaries.y+mBoundaries.height) 
			{
				getRect().y = newY;
			} 
			else if(newY < mBoundaries.y) 
			{
				getRect().y = mBoundaries.y;
			} 
			else if(newY > mBoundaries.height) 
			{
				getRect().y = mBoundaries.height;
			}

			updateScrollChangeListener();
		}
	}
	
	public int getScrolledX()
	{
		return getRect().x - mBoundaries.x;
	}
	
	public int getScrolledY()
	{
		return getRect().y - mBoundaries.y;
	}
}
