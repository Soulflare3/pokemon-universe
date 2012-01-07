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
		//
		switch(header)
		{
			case PU_Packet.HEADER_LOGIN:
				receiveLoginStatus(packet);
				break;
				
			case PU_Packet.HEADER_IDENTITY:
				receiveIdentity(packet);
				break;
				
			case PU_Packet.HEADER_TILES:
				receiveTiles(packet);
				break;
				
			case PU_Packet.HEADER_WALK:
				receiveCreatureMove(packet);
				break;
				
			case PU_Packet.HEADER_TURN:
				receiveCreatureTurn(packet);
				break;
				
			default:
				PUWeb.log("Received packet with unknown header: " + header);
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
	
	public void sendWalk(int direction, boolean requestTiles)
	{
		PU_Packet packet = new PU_Packet();
		packet.addUInt8(PU_Packet.HEADER_WALK);
		packet.addUint16(direction);
		if(requestTiles)
		{
			packet.addUint16(1);
		}
		else
		{
			packet.addUint16(0);
		}
		mConn.sendPacket(packet);
	}
	
	public void sendTurn(int direction)
	{
		PU_Packet packet = new PU_Packet();
		packet.addUInt8(PU_Packet.HEADER_TURN);
		packet.addUint16(direction);
		mConn.sendPacket(packet);
	}
	
	// Receive
	public void receiveLoginStatus(PU_Packet packet)
	{
		int loginStatus = packet.readUint8();
		PU_Login.setLoginStatus(loginStatus);
	}
	
	public void receiveIdentity(PU_Packet packet)
	{
		PU_Player player = new PU_Player(packet.readUint64());
		player.setName(packet.readString());
		int x = packet.readUint16();
		int y = packet.readUint16();
		player.setPosition(x, y);
		player.setDirection(packet.readUint16());
		player.setMoney(packet.readUint32());
		
		PUWeb.map().addCreature(player);
		PUWeb.game().setSelf(player);
	}
	
	public void receiveTiles(PU_Packet packet)
	{
		int tileCount = packet.readUint16();
		if(tileCount > 0)
		{
			for(int i = 0; i < tileCount; i++)
			{
				receiveTile(packet);
			}
		}
	}
	
	public void receiveTile(PU_Packet packet)
	{
		boolean tileExists = true;
		int[] layers = new int[]{-1, -1, -1};
		
		int x = packet.readUint16();
		int y = packet.readUint16();
		int movement = packet.readUint16();
		
		PU_Tile tile = PUWeb.map().getTile(x, y);
		if(tile == null)
			tileExists = false;
		
		int numLayers = packet.readUint16();
		for(int i = 0; i < numLayers; i++)
		{
			int layer = packet.readUint16();
			int id = (int)packet.readUint32();
			
			layers[layer] = id;
		}
		
		if(!tileExists)
		{
			tile = PUWeb.map().addTile(x, y);
			tile.setMovement(movement);
			
			for(int i = 0; i < 3; i++)
			{
				if(layers[i] != -1)
				{
					tile.addLayer(i, layers[i]);
				}
			}
		}
		else
		{
			long signature = (long)movement;
			int shift = 16;
			for(int i = 0; i < 3; i++)
			{
				if(layers[i] != -1)
				{
					signature |= ((long)layers[i] << shift);
				}
				shift += 16;
			}
			
			if(tile.getSignature() != signature)
			{
				tile.setMovement(movement);
				for(int i = 0; i < 3; i++)
				{
					tile.removeLayer(i);
					if(layers[i] != -1)
					{
						tile.addLayer(i, layers[i]);
					}
				}
			}
		}
		
		packet.readUint16(); // town id
		packet.readString(); // town name
	}
	
	public void receiveCreatureMove(PU_Packet packet)
	{
		PU_Creature creature = PUWeb.map().getCreatureById(packet.readUint32());
		int fromX = packet.readUint16();
		int fromY = packet.readUint16();
		int toX = packet.readUint16();
		int toY = packet.readUint16();
		PU_Tile toTile = PUWeb.map().getTile(toX, toY);
		PU_Tile fromTile = PUWeb.map().getTile(fromX, fromY);
		if(creature != null)
		{
			if(Math.abs(fromX-toX) > 1 || Math.abs(fromY-toY) > 1)
			{
				creature.setPosition(toX, toY);
			}
			else
			{
				if(creature instanceof PU_Player)
				{
					((PU_Player)creature).receiveWalk(fromTile, toTile);
				}
			}
		}
	}
	
	public void receiveCreatureTurn(PU_Packet packet)
	{
		PU_Creature creature = PUWeb.map().getCreatureById(packet.readUint32());
		int direction = packet.readUint16();
		if(creature != null && creature != PUWeb.game().getSelf())
		{
			creature.setDirection(direction);
		}
	}
}
