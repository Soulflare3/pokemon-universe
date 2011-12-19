package pu.web.client;

public class PU_Rect
{
	public int x, y, width, height;

	public PU_Rect()
	{
		x = 0;
		y = 0;
		width = 0;
		height = 0;
	}

	public PU_Rect(int _x, int _y, int _width, int _height)
	{
		x = _x;
		y = _y;
		width = _width;
		height = _height;
	}
	
	public PU_Rect(PU_Rect rect)
	{
		x = rect.x;
		y = rect.y;
		width = rect.width;
		height = rect.height;
	}

	public boolean equals(PU_Rect r)
	{
		if (r.x == x && r.y == y && r.width == width && r.height == height)
			return true;
		else
			return false;
	}

	public boolean isAll(int i)
	{
		if (x == i && y == i && width == i && height == i)
			return true;
		else
			return false;
	}

	public boolean contains(int _x, int _y)
	{
		return _x >= x && _x <= x + width && _y >= y && _y <= y + height;
	}

	public boolean contains(PU_Rect rect)
	{
		return contains(rect.x, rect.y) && contains(rect.x + rect.width, rect.y) && contains(rect.x, rect.y + rect.height) && contains(rect.x + rect.width, rect.y + rect.height);
	}

	public boolean intersects(PU_Rect rect)
	{
		return !(x > rect.x + rect.width || rect.x > x + width || y > rect.y + rect.height || rect.y > y + height);
	}

	public PU_Rect intersection(PU_Rect rect)
	{
		int tempX = 0, tempY = 0, tempW = 0, tempH = 0;

		if (intersects(rect))
		{
			tempX = x > rect.x ? x : rect.x;
			tempY = y > rect.y ? y : rect.y;
			tempW = x + width < rect.x + rect.width ? x + width : rect.x + rect.width;
			tempH = y + height < rect.y + rect.height ? y + height : rect.y + rect.height;

			tempW -= tempX;
			tempH -= tempY;
		}

		return new PU_Rect(tempX, tempY, tempW, tempH);
	}
}