package pu.web.client;

public class PU_Protocol
{
	private PU_Connection mConn;
	
	public PU_Protocol(PU_Connection conn)
	{
		mConn = conn;
	}
	
	public void parsePacket(PU_Packet packet)
	{
		byte header = packet.readUint8();
		switch(header)
		{
			case PU_Packet.HEADER_LOGIN:
				receiveLoginStatus(packet);
				break;
		}
	}
	
	// Send
	public void sendLogin(String username, String password, int version)
	{
		PU_Packet packet = new PU_Packet();
		packet.addUInt8(PU_Packet.HEADER_LOGIN);
		packet.addString(username);
		packet.addString(password);
		packet.addUint16(version);
		mConn.sendPacket(packet);		
	}
	
	public void sendRequestLogin()
	{
		PU_Packet packet = new PU_Packet();
		packet.addUInt8(PU_Packet.HEADER_LOGIN);
		mConn.sendPacket(packet);
	}
	
	// Receive
	public void receiveLoginStatus(PU_Packet packet)
	{
		int loginStatus = packet.readUint8();
		PU_Login.setLoginStatus(loginStatus);
	}
}
