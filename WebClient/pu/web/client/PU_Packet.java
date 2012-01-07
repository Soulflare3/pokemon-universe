package pu.web.client;

import com.googlecode.gwtgl.array.ArrayBuffer;
import com.googlecode.gwtgl.array.Uint8Array;


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

	private Uint8Array buffer;

	final int PACKET_MAXSIZE = 16384;

	public PU_Packet()
	{
		buffer = Uint8Array.create(PACKET_MAXSIZE);
		reset();
	}
	
	public PU_Packet(ArrayBuffer data)
	{
		try
		{
			buffer = Uint8Array.create(data);
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
	
	public ArrayBuffer getBuffer()
	{
		Uint8Array data = Uint8Array.create(msgSize+2);
		for(int i = 0; i < msgSize+2; i++)
			data.set(i, buffer.get(i));

		return data.getBuffer();
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
	
	public void setHeader()
	{
		buffer.set(0, (byte) (msgSize));
		buffer.set(1, (byte) (msgSize >> 8));
	}

	public void addUInt8(byte value)
	{
		if (!canAdd(1))
		{
			return;
		}

		buffer.set(readPos++, value);
		msgSize++;
	}

	public void addUint8(int value)
	{
		if (!canAdd(1))
		{
			return;
		}

		buffer.set(readPos++, (byte) value);
		msgSize++;
	}

	public void addUint16(int value)
	{
		if (!canAdd(2))
		{
			return;
		}
		
		buffer.set(readPos++, (byte) (value));
		buffer.set(readPos++, (byte) (value >> 8));
		msgSize += 2;
	}

	public void addUint32(long value)
	{
		if (!canAdd(4))
		{
			return;
		}
		
		buffer.set(readPos++, (byte) (value));
		buffer.set(readPos++, (byte) (value >> 8));
		buffer.set(readPos++, (byte) (value >> 16));
		buffer.set(readPos++, (byte) (value >> 24));
		msgSize += 4;
	}

	public void addString(String string)
	{
		addUint16(string.length());
		for (int i = 0; i < string.length(); i++)
		{
			buffer.set(readPos++, (byte) (byte) string.charAt(i));
		}
		msgSize += string.length();
	}

	public byte readUint8()
	{
		return ((byte)buffer.get(readPos++));
	}

	public int readUint16()
	{
		int v = ((((byte)buffer.get(readPos)) & 0xFF)) | ((((byte)buffer.get(readPos + 1)) & 0xFF) << 8);

		readPos += 2;
		return v;
	}

	public long readUint32()
	{
		long v = ((((byte)buffer.get(readPos)) & 0xFF) | ((((byte)buffer.get(readPos + 1)) & 0xFF) << 8) | ((((byte)buffer.get(readPos + 2)) & 0xFF) << 16) | ((((byte)buffer.get(readPos + 3)) & 0xFF) << 24));

		readPos += 4;
		return v;
	}
	
	public long readUint64()
	{
		long v = ((((byte)buffer.get(readPos)) & 0xFF) | 
				((((byte)buffer.get(readPos + 1)) & 0xFF) << 8) | 
				((((byte)buffer.get(readPos + 2)) & 0xFF) << 16) | 
				((((byte)buffer.get(readPos + 3)) & 0xFF) << 24) |
				((((byte)buffer.get(readPos + 4)) & 0xFF) << 32) |
				((((byte)buffer.get(readPos + 5)) & 0xFF) << 40) |
				((((byte)buffer.get(readPos + 6)) & 0xFF) << 48) |
				((((byte)buffer.get(readPos + 7)) & 0xFF) << 56));

		readPos += 8;
		return v;
	}

	public String readString()
	{
		String string = "";
		int stringlength = readUint16();
		for (int i = 0; i < stringlength; i++)
		{
			string = string + (char) ((byte)buffer.get(readPos++));
		}
		return string;
	}
}