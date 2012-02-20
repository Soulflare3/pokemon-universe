package pu.web.client;

public class PU_Creature
{
	public static final int DIR_SOUTH = 1;
	public static final int DIR_WEST = 2;
	public static final int DIR_NORTH = 3;
	public static final int DIR_EAST = 4;
	
	protected long mId = 0;
	protected String mName = "";
	
	protected int mX = 0;
	protected int mY = 0;
	
	protected boolean mWalking = false;
	protected boolean mWalkEnded = false;
	protected int mPreWalkX = 0;
	protected int mPreWalkY = 0;
	protected int mOffset = 0;
	protected float mWalkProgress = 0.0f;
	protected int mSpeed = 300;
	
	protected int mDirection = DIR_SOUTH;
	
	protected int mFrame = 0;
	protected int mFrames = 3;
	
	protected boolean mAnimationRunning = false;
	protected int mAnimationInterval = 150;
	protected long mAnimationLastTicks = System.currentTimeMillis();
	
	public PU_Creature()
	{
		
	}
	
	public void draw(int x, int y)
	{
		
	}
	
	public long getId()
	{
		return mId;
	}
	
	public void setName(String name)
	{
		mName = name;
	}
	
	public String getName()
	{
		return mName;
	}
	
	public boolean isWalking()
	{
		return mWalking;
	}
	
	public int getOffset()
	{
		return mOffset;
	}
	
	public void setDirection(int direction)
	{
		mDirection = direction;
	}
	
	public int getDirection()
	{
		return mDirection;
	}
	
	public int getX()
	{
		if(mWalking)
		{
			return mPreWalkX;
		}
		return mX;
	}
	
	public int getY()
	{
		if(mWalking)
		{
			return mPreWalkY;
		}
		return mY;
	}
	
	public void setPosition(int x, int y)
	{
		mX = x;
		mY = y;
	}
	
	public void startAnimation()
	{
		mAnimationRunning = true;
	}
	
	public void stopAnimation()
	{
		mAnimationRunning = false;
		mFrame = 0;
	}
	
	public void updateAnimation()
	{
		if(mAnimationRunning)
		{
			long passedTicks = System.currentTimeMillis() - mAnimationLastTicks;
			if(passedTicks >= mAnimationInterval)
			{
				mFrame++;
				if(mFrame > mFrames)
				{
					mFrame = 0;
				}
				
				mAnimationLastTicks = System.currentTimeMillis();
			}
		}
	}
}
