package pu.web.client;

public class PU_Attack
{
	private long mId;
	private String mName;
	private String mFlavor;
	private int mType;
	private String mTypeName;
	private int mPp;
	private int mMaxPp;
	private int mPower;
	private int mAccuracy;
	private int mTargetId;

	public PU_Attack()
	{
		
	}
	
	public long getId()
	{
		return mId;
	}

	public void setId(long mId)
	{
		this.mId = mId;
	}

	public String getName()
	{
		return mName;
	}

	public void setName(String mName)
	{
		this.mName = mName;
	}

	public String getFlavor()
	{
		return mFlavor;
	}

	public void setFlavor(String mFlavor)
	{
		this.mFlavor = mFlavor;
	}

	public int getType()
	{
		return mType;
	}

	public void setType(int mType)
	{
		this.mType = mType;
	}

	public String getTypeName()
	{
		return mTypeName;
	}

	public void setTypeName(String mTypeName)
	{
		this.mTypeName = mTypeName;
	}

	public int getPp()
	{
		return mPp;
	}

	public void setPp(int mPp)
	{
		this.mPp = mPp;
	}

	public int getMaxPp()
	{
		return mMaxPp;
	}

	public void setMaxPp(int mMaxPp)
	{
		this.mMaxPp = mMaxPp;
	}

	public int getPower()
	{
		return mPower;
	}

	public void setPower(int mPower)
	{
		this.mPower = mPower;
	}

	public int getAccuracy()
	{
		return mAccuracy;
	}

	public void setAccuracy(int mAccuracy)
	{
		this.mAccuracy = mAccuracy;
	}

	public int getTargetId()
	{
		return mTargetId;
	}

	public void setTargetId(int mTargetId)
	{
		this.mTargetId = mTargetId;
	}
}
