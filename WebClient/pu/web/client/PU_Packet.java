package pu.web.client;

public class PU_Packet
{
	private int msgSize;
	private int readPos;

	private byte[] buffer;

	final int PACKET_MAXSIZE = 16384;

	public PU_Packet()
	{
		buffer = new byte[PACKET_MAXSIZE];
		reset();
	}
	
	public PU_Packet(String data)
	{
		try
		{
			buffer = data.getBytes("UTF-8");
		}
		catch(Exception e)
		{
			
		}
		reset();
	}

	public void reset()
	{
		msgSize = 0;
		readPos = 2;
	}

	public boolean canAdd(int size)
	{
		if (size + readPos < PACKET_MAXSIZE - 16)
		{
			return true;
		} else
		{
			return false;
		}
	}

	public String buildMessage()
	{
		String message = "";
		buffer[0] = (byte) (msgSize);
		buffer[1] = (byte) (msgSize >> 8);
		try
		{
			message = new String(buffer, "UTF-8");
		}
		catch(Exception e)
		{
			
		}
		return message;
	}

	public void addUInt8(byte value)
	{
		if (!canAdd(1))
		{
			return;
		}

		buffer[readPos++] = value;
		msgSize++;
	}

	public void addUint8(int value)
	{
		if (!canAdd(1))
		{
			return;
		}

		buffer[readPos++] = (byte) value;
		msgSize++;
	}

	public void addUint16(int value)
	{
		if (!canAdd(2))
		{
			return;
		}

		buffer[readPos++] = (byte) (value);
		buffer[readPos++] = (byte) (value >> 8);
		msgSize += 2;
	}

	public void addUint32(long value)
	{
		if (!canAdd(4))
		{
			return;
		}

		buffer[readPos++] = (byte) (value);
		buffer[readPos++] = (byte) (value >> 8);
		buffer[readPos++] = (byte) (value >> 16);
		buffer[readPos++] = (byte) (value >> 24);
		msgSize += 4;
	}

	public void addString(String string)
	{
		addUint16(string.length());
		for (int i = 0; i < string.length(); i++)
		{
			buffer[readPos++] = (byte) string.charAt(i);
		}
		msgSize += string.length();
	}

	public byte readUint8()
	{
		return buffer[readPos++];
	}

	public int readUint16()
	{
		int v = ((buffer[readPos] & 0xFF)) | ((buffer[readPos + 1] & 0xFF) << 8);

		readPos += 2;
		return v;
	}

	public long readUint32()
	{
		long v = ((buffer[readPos] & 0xFF) | ((buffer[readPos + 1] & 0xFF) << 8) | ((buffer[readPos + 2] & 0xFF) << 16) | ((buffer[readPos + 3] & 0xFF) << 24));

		readPos += 4;
		return v;
	}

	public String readString()
	{
		String string = "";
		int stringlength = readUint16();
		for (int i = 0; i < stringlength; i++)
		{
			string = string + (char) buffer[readPos++];
		}
		return string;
	}
}