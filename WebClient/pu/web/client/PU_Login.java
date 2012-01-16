package pu.web.client;

import pu.web.client.gui.Label;
import pu.web.client.gui.OnKeyDownListener;
import pu.web.client.gui.Panel;
import pu.web.client.gui.TextField;
import pu.web.client.resources.gui.GuiImages;

public class PU_Login extends Panel
{
	public static final int LOGINPHASE_IDLE = 0;
	public static final int LOGINPHASE_CONNECT = 1;
	public static final int LOGINPHASE_CONNECTING = 2;
	public static final int LOGINPHASE_VERIFYDATA = 3;
	public static final int LOGINPHASE_LOADGAMEWORLD = 4;
	
	public final static int LOGINSTATUS_IDLE            = 0;
	public final static int LOGINSTATUS_WRONGACCOUNT    = 1;
	public final static int LOGINSTATUS_SERVERERROR     = 2;
	public final static int LOGINSTATUS_DATABASEERROR   = 3;
	public final static int LOGINSTATUS_ALREADYLOGGEDIN = 4;
	public final static int LOGINSTATUS_READY           = 5;
	public final static int LOGINSTATUS_CHARBANNED      = 6;
	public final static int LOGINSTATUS_SERVERCLOSED    = 7;
	public final static int LOGINSTATUS_WRONGVERSION    = 8;
	public final static int LOGINSTATUS_FAILPROFILELOAD = 9;
	
	private static int mLoginStatus;
	
	private TextField tfUsername;
	private TextField tfPassword;
	private Label lbStatus;
	private int mPhase = LOGINPHASE_IDLE;
	
	private long mTimeoutTicks = 0;
	private long mLastTimeoutTick = 0;
	
	public PU_Login()
	{
		super(0, 0, PU_Engine.SCREEN_WIDTH, PU_Engine.SCREEN_HEIGHT);
		
		setupControls();
	}
	
	public static int getLoginStatus()
	{
		return mLoginStatus;
	}
	
	public static void setLoginStatus(int loginStatus)
	{
		mLoginStatus = loginStatus;
	}
	
	private void setupControls()
	{
		tfUsername = new TextField(451, 453, 160, 20);
		tfUsername.setFontColor(0, 0, 0);
		tfUsername.setOnKeyDownListener(mOnKeyDownListener);
		addChild(tfUsername);
		tfUsername.setFocus(true);
		
		tfPassword = new TextField(451, 515, 160, 20);
		tfPassword.setFontColor(0, 0, 0);
		tfPassword.setPassword(true);
		tfPassword.setOnKeyDownListener(mOnKeyDownListener);
		addChild(tfPassword);
		
		lbStatus = new Label(407, 389, "");
		lbStatus.setFontColor(0, 185, 47);
		addChild(lbStatus);
	}
	
	public OnKeyDownListener mOnKeyDownListener = new OnKeyDownListener()
	{
		@Override
		public void OnKeyDown(int button)
		{
			if(button == 13)
			{
				if(mPhase == LOGINPHASE_IDLE)
				{
					startLogin();
				}
			}
		}
	};
	
	private void startLogin()
	{
		lbStatus.setText("");
		lbStatus.setFontColor(0, 185, 47);
		
		tfUsername.setReadOnly(true);
		tfPassword.setReadOnly(true);
		
		mPhase = LOGINPHASE_CONNECT;
	}
	
	private void processLogin()
	{
		switch(mPhase)
		{
		case LOGINPHASE_CONNECT:
			lbStatus.setText("Connecting to server...");
			PUWeb.connection().connect();
			mPhase = LOGINPHASE_CONNECTING;
			break;
			
		case LOGINPHASE_CONNECTING:
			if(PUWeb.connection().getState() == PU_Connection.STATE_CONNECTED)
			{
				lbStatus.setText("Verifying username and password...");
				
				mLoginStatus = LOGINSTATUS_IDLE;
				PUWeb.connection().getProtocol().sendLogin(tfUsername.getText(), tfPassword.getText(), PUWeb.CLIENT_VERSION);
				
				mTimeoutTicks = 0;
				mLastTimeoutTick = System.currentTimeMillis();
				mPhase = LOGINPHASE_VERIFYDATA;
			}
			else if(PUWeb.connection().getState() == PU_Connection.STATE_DISCONNECTED)
			{
				loginFailed("Could not connect to server.");
			}
			break;
		
		case LOGINPHASE_VERIFYDATA:
			if(mLoginStatus  != LOGINPHASE_IDLE)
			{
				switch(mLoginStatus) 
				{
				case LOGINSTATUS_WRONGACCOUNT:
					loginFailed("Invalid username and/or password.");
					return;

				case LOGINSTATUS_SERVERERROR:
					loginFailed("The server has returned an error. Please retry.");
					return;

				case LOGINSTATUS_DATABASEERROR:
					loginFailed("The database has returned an error. Please retry.");
					return;

				case LOGINSTATUS_ALREADYLOGGEDIN:
					loginFailed("You are already logged in.");
					return;

				case LOGINSTATUS_CHARBANNED:
					loginFailed("This account is banned from the game.");
					return;

				case LOGINSTATUS_SERVERCLOSED:
					loginFailed("The server is currently closed.");
					return;

				case LOGINSTATUS_WRONGVERSION:
					loginFailed("This client version is outdated.");
					return;
					
				case LOGINSTATUS_FAILPROFILELOAD:
					loginFailed("Your profile could not be loaded. Please retry.");
					return;
				}
				
				lbStatus.setText("Loading gameworld...");
				
				mLoginStatus = LOGINSTATUS_IDLE;
				PUWeb.connection().getProtocol().sendRequestLogin();
				
				mTimeoutTicks = 0;
				mLastTimeoutTick = System.currentTimeMillis();
				mPhase = LOGINPHASE_LOADGAMEWORLD;
			}
			else
			{
				mTimeoutTicks += System.currentTimeMillis() - mLastTimeoutTick;
				mLastTimeoutTick = System.currentTimeMillis();
				if(mTimeoutTicks >= 10000) // 10 seconds
				{
					loginFailed("Timeout while verifying data. Please retry.");
					break;
				}
			}
			break;
			
		case LOGINPHASE_LOADGAMEWORLD:
			if(mLoginStatus == LOGINSTATUS_READY)
			{				
				mPhase = LOGINPHASE_IDLE;
				PUWeb.hideLogin();
				PUWeb.game().showWorldPanel();
				PUWeb.game().setState(PU_Game.GAMESTATE_WORLD);
			}
			else
			{
				mTimeoutTicks += System.currentTimeMillis() - mLastTimeoutTick;
				mLastTimeoutTick = System.currentTimeMillis();
				if(mTimeoutTicks >= 30000) // 30 seconds
				{
					loginFailed("Timeout while loading gameworld. Please retry.");
					break;
				}
			}
			break;
		}
	}
	
	private void loginFailed(String message)
	{
		PUWeb.connection().close();
		mPhase = LOGINPHASE_IDLE;
		
		lbStatus.setText(message);
		lbStatus.setFontColor(202, 0, 0);
		
		tfUsername.setReadOnly(false);
		tfPassword.setReadOnly(false);
	}
	
	@Override
	public void draw(PU_Rect drawArea)
	{
		PU_Image intro = PUWeb.resources().getGuiImage(GuiImages.IMG_GUI_INTROBG);
		if(intro != null)
		{
			intro.draw(0, 0);
		}
		
		super.draw(drawArea);
		
		// Abuse the draw loop for other stuff
		if(mPhase != LOGINPHASE_IDLE)
		{
			processLogin();
		}
	}
}
