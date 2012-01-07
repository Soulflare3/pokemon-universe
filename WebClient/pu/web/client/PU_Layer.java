package pu.web.client;

public class PU_Layer
{
	private int mId = 0;
	private PU_Image mImage = null;
	
	public PU_Layer(int id)
	{
		setId(id);
	}
	
	public void setId(int id)
	{
		mId = id;
		mImage = PUWeb.resources().getTileImage(id);
	}
	
	public int getId()
	{
		return mId;
	}
	
	public PU_Image getImage()
	{
		return mImage;
	}
	
	public void draw(int x, int y)
	{
		if(mImage != null)
		{
			PUWeb.engine().addToSpriteBatch(mImage, x, y);
		}
	}
}
