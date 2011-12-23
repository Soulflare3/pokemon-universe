package pu.web.client;

public class PU_Connection
{
	public static final int STATE_DISCONNECTED = 0;
	public static final int STATE_CONNECTING = 1;
	public static final int STATE_CONNECTED = 2;
	
	private String mServer;
	private int mState = STATE_DISCONNECTED; 
	
	private PU_Protocol mProtocol;

	public PU_Connection(String server)
	{
		mServer = server;
		mProtocol = new PU_Protocol(this);
	}
	
	public void connect()
	{
		mState = STATE_CONNECTING; 
		nativeConnect(mServer);
	}
	
	public int getState()
	{
		return mState;
	}
	
	public PU_Protocol getProtocol()
	{
		return mProtocol;
	}

	private native boolean nativeConnect(String server) /*-{
		var connection = this;
		var websocket = null;
		
		if ($wnd.WebSocket) {
			websocket = $wnd.WebSocket;
		} else if ($wnd.MozWebSocket) {
			websocket = $wnd.MozWebSocket;
		}

		if (!websocket) {
			alert("Websocket connections not supported by this browser. Get the latest Chrome or Firefox!");
			return false;
		}

		$wnd.socket = new websocket(server);
		console.log("Websocket tried to connect to " + server + " Readystate: "  + $wnd.socket.readyState);

		$wnd.socket.onopen = function() {
			console.log("Readystate: "  + $wnd.socket.readyState);
			connection.@pu.web.client.PU_Connection::onSocketOpen()();
		};

		$wnd.socket.onmessage = function(response) {
			connection.@pu.web.client.PU_Connection::onSocketReceive(Ljava/lang/String;)(response.data);
		};

		$wnd.socket.onclose = function(m) {
			connection.@pu.web.client.PU_Connection::onSocketClose()();
		};
		
		return true;
	}-*/;

	private final void onSocketOpen()
	{
		mState = STATE_CONNECTED;
	}

	private final void onSocketClose()
	{
		if(mState == STATE_CONNECTING)
		{
			PUWeb.log("Connection could not be established.");
		}
		else
		{
			PUWeb.log("Connection closed.");
		}
		mState = STATE_DISCONNECTED;
	}

	private final void onSocketReceive(String message)
	{
		PU_Packet packet = new PU_Packet(message);
		mProtocol.parsePacket(packet);
	}
	
	public void sendPacket(PU_Packet packet)
	{
		nativeSend(packet.buildMessage());
	}
	
	private native void nativeSend(String message) /*-{
		if ($wnd.socket) {
			if ($wnd.socket.readyState == 1) {
				$wnd.socket.send(message);
			} else {
				console.log("Send error: Socket is not ready to send data.");
			}
		} else {
			console.log("Send error: Socket not created or opened.");
		}
	}-*/;
}
