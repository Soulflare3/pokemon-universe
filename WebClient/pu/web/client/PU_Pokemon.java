package pu.web.client;

public class PU_Pokemon
{
	private long mId;
	private int mSpeciesId;
	private int mLevel;
	private int mHp;
	private int mHpmax;
	private int mSex;
	private int mNature;
	
	private int mExpPerc;
	private long mExpCurrent;
	private long mExpTnl;
	
	private String mName;
	private String mFlavor;
	
	private int mType1;
	private int mType2;
	
	private int[] mStats = new int[6];
	private PU_Attack[] mAttacks = new PU_Attack[4];
	
	public PU_Pokemon()
	{
		
	}
	
	public void setId(long id)
	{
		this.mId = id;
	}
	
	public long getId()
	{
		return this.mId;
	}
	
	public void setSpeciesId(int speciesId)
	{
		this.mSpeciesId = speciesId;
	}
	
	public int getSpeciesId()
	{
		return this.mSpeciesId;
	}
	
	public void setLevel(int level)
	{
		this.mLevel = level;
	}
	
	public int getLevel()
	{
		return this.mLevel;
	}
	
	public void setHp(int hp)
	{
		this.mHp = hp;
	}
	
	public int getHp()
	{
		return this.mHp;
	}
	
	public void setHpmax(int hpmax)
	{
		this.mHpmax = hpmax;
	}
	
	public int getHpmax()
	{
		return this.mHpmax;
	}
	
	public void setSex(int sex)
	{
		this.mSex = sex;
	}
	
	public int getSex()
	{
		return this.mSex;
	}
	
	public void setNature(int nature)
	{
		this.mNature = nature;
	}
	
	public int getNature()
	{
		return this.mNature;
	}
	
	public void setExpPerc(int expPerc)
	{
		this.mExpPerc = expPerc;
	}
	
	public int getExpPerc()
	{
		return this.mExpPerc;
	}
	
	public void setExpCurrent(long expCurrent)
	{
		this.mExpCurrent = expCurrent;
	}
	
	public long getExpCurrent()
	{
		return this.mExpCurrent;
	}
	
	public void setExpTnl(long expTnl)
	{
		this.mExpTnl = expTnl;
	}
	
	public long getExpTnl()
	{
		return this.mExpTnl;
	}
	
	public void setName(String name)
	{
		this.mName = name;
	}
	
	public String getName()
	{
		return this.mName;
	}
	
	public void setFlavor(String flavor)
	{
		this.mFlavor = flavor;
	}
	
	public String getFlavor()
	{
		return this.mFlavor;
	}
	
	public void setType1(int type)
	{
		this.mType1 = type;
	}
	
	public int getType1()
	{
		return this.mType1;
	}
	
	public void setType2(int type)
	{
		this.mType2 = type;
	}
	
	public int getType2()
	{
		return this.mType2;
	}
	
	public void setStat(int stat, int value)
	{
		this.mStats[stat] = value;
	}
	
	public int getStat(int stat)
	{
		return this.mStats[stat];
	}
	
	public void setAttack(int slot, PU_Attack attack)
	{
		this.mAttacks[slot] = attack;
	}
	
	public PU_Attack getAttack(int slot)
	{
		return this.mAttacks[slot];
	}
	
	public static String getTypeById(int type)
	{
		switch(type)
		{
		case 100:
            return "ground";

	    case 101:
	            return "water";
	
	    case 102:
	            return "ghost";
	
	    case 103:
	            return "bug";
	
	    case 104:
	            return "fighting";
	
	    case 105:
	            return "psychic";
	
	    case 106:
	            return "grass";
	
	    case 107:
	            return "dark";
	
	    case 108:
	            return "normal";
	
	    case 109:
	            return "poison";
	
	    case 110:
	            return "electric";
	
	    case 111:
	            return "unknown";
	
	    case 112:
	            return "steel";
	
	    case 113:
	            return "rock";
	
	    case 114:
	            return "dragon";
	
	    case 115:
	            return "flying";
	
	    case 116:
	            return "fire";
	
	    case 117:
	            return "ice";
		}
		return "unknown";
	}
}
