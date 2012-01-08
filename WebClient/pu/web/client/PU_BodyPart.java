package pu.web.client;

public class PU_BodyPart
{
	public int id = 0;
	public int red = 0;
	public int green = 0;
	public int blue = 0;
	
	public PU_BodyPart(int id)
	{
		this.id = id;
	}
	
	public void setColor(int red, int green, int blue)
	{
		this.red = red;
		this.green = green;
		this.blue = blue;
	}
}
