package pu.web.client;

public class PU_Packet
{
	// Headers
    public static final byte HEADER_PING   = (byte)0x00;
    public static final byte HEADER_LOGIN  = (byte)0x01;
    public static final byte HEADER_LOGOUT = (byte)0x02;
    
    public static final byte HEADER_CHAT = (byte)0x10;

    public static final byte HEADER_IDENTITY = (byte)0xAA;

    public static final byte HEADER_WALK       = (byte)0xB1;
    public static final byte HEADER_CANCELWALK = (byte)0xB2;
    public static final byte HEADER_WARP       = (byte)0xB3;
    public static final byte HEADER_TURN       = (byte)0xB4;

    public static final byte HEADER_TILES          = (byte)0xC1;
    public static final byte HEADER_ADDCREATURE    = (byte)0xC2;
    public static final byte HEADER_REMOVECREATURE = (byte)0xC3;
    
    public static final byte HEADER_REFRESHCOMPLETE = (byte)0x03;
    public static final byte HEADER_REFRESHWORLD    = (byte)0xC4;
    
    public static final byte HEADER_REFRESHPOKEMON    = (byte)0xD1;
    //
	
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
	
	public PU_Packet(byte[] data)
	{
		try
		{
			buffer = data;
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
	
	public long readUint64()
	{
		long v = ((buffer[readPos] & 0xFF) | 
				((buffer[readPos + 1] & 0xFF) << 8) | 
				((buffer[readPos + 2] & 0xFF) << 16) | 
				((buffer[readPos + 3] & 0xFF) << 24) |
				((buffer[readPos + 4] & 0xFF) << 32) |
				((buffer[readPos + 5] & 0xFF) << 40) |
				((buffer[readPos + 6] & 0xFF) << 48) |
				((buffer[readPos + 7] & 0xFF) << 56));

		readPos += 8;
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