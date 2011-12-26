package pu.web.client;

import pu.web.client.gui.GUIManager;
import pu.web.client.resources.fonts.Fonts;

import com.google.gwt.core.client.EntryPoint;
import com.google.gwt.core.client.GWT;
import com.google.gwt.core.client.RunAsyncCallback;
import com.google.gwt.dom.client.Document;
import com.google.gwt.user.client.ui.RootPanel;
import com.googlecode.gwtgl.binding.WebGLCanvas;
import com.googlecode.gwtgl.binding.WebGLRenderingContext;

public class PUWeb implements EntryPoint
{
	public static final int CLIENT_VERSION = 4;
			
	private static WebGLRenderingContext mGlContext;
	private static PU_Engine mEngine;
	private static GUIManager mGui;
	private static PU_Resources mResources;
	private static PU_Game mGame;
	private static PU_Map mMap;
	private static PU_Events mEvents;
	private static PU_Connection mConnection;
	
	private static long mFrameTime = 0;
	private long mLastFrameTime = System.currentTimeMillis();

	public void onModuleLoad()
	{
		final WebGLCanvas webGLCanvas = new WebGLCanvas(PU_Engine.SCREEN_WIDTH + "px", PU_Engine.SCREEN_HEIGHT + "px");
		PUWeb.mGlContext = webGLCanvas.getGlContext();
		PUWeb.mGlContext.viewport(0, 0, PU_Engine.SCREEN_WIDTH, PU_Engine.SCREEN_HEIGHT);
		RootPanel.get("gwtGL").add(webGLCanvas);
		
		PUWeb.mResources = new PU_Resources();
		PUWeb.mGame = new PU_Game();
		PUWeb.mMap = new PU_Map();
		
		PUWeb.mEngine = new PU_Engine(PUWeb.mGlContext);
		PUWeb.mEngine.init();

		// Start the draw loop
		drawScene();

		GWT.runAsync(new RunAsyncCallback()
		{
			@Override
			public void onSuccess()
			{
				// Load fonts
				mResources.loadFonts();
				
				// Load GUI images
				mResources.loadGuiImages();
				
				// Load tiles
				mResources.loadTiles();
			}

			@Override
			public void onFailure(Throwable reason)
			{
			}
		});

		mConnection = new PU_Connection("ws://127.0.0.1:6161/puserver");
		mConnection.connect();
	}
	
	static PU_Login login;
	public static void resourcesLoaded()
	{
		PUWeb.mGui = new GUIManager(0, 0, PU_Engine.SCREEN_WIDTH, PU_Engine.SCREEN_HEIGHT, mResources.getFont(Fonts.FONT_PURITAN_BOLD_14));
		PUWeb.mEvents = new PU_Events(Document.get().getElementById("gwtGL"), PUWeb.mGui);
		
		mGame.setState(PU_Game.GAMESTATE_LOGIN);
		
		// TODO: move this to the appropriate place
		login = new PU_Login();
		mGui.getRoot().addChild(login);
	}
	
	public static void hideLogin()
	{
		mGui.getRoot().removeChild(login);
	}

	public static WebGLRenderingContext gl()
	{
		return PUWeb.mGlContext;
	}

	public static PU_Engine engine()
	{
		return PUWeb.mEngine;
	}
	
	public static PU_Resources resources()
	{
		return PUWeb.mResources;
	}
	
	public static GUIManager gui()
	{
		return PUWeb.mGui;
	}
	
	public static PU_Connection connection()
	{
		return mConnection;
	}
	
	public static PU_Game game()
	{
		return mGame;
	}
	
	public static PU_Map map()
	{
		return mMap;
	}
	
	public static PU_Events events()
	{
		return mEvents;
	}
	
	public static long getFrameTime()
	{
		return mFrameTime;
	}

	private native void requestAnimationFrame() /*-{
		var puweb = this;
		var fn = function() {
			puweb.@pu.web.client.PUWeb::drawScene()();
		};
		if ($wnd.requestAnimationFrame) {
			$wnd.requestAnimationFrame(fn);
		} else if ($wnd.mozRequestAnimationFrame) {
			$wnd.mozRequestAnimationFrame(fn);
		} else if ($wnd.webkitRequestAnimationFrame) {
			$wnd.webkitRequestAnimationFrame(fn);
		} else if ($wnd.oRequestAnimationFrame) {
			$wnd.oRequestAnimationFrame(fn);
		} else if ($wnd.msRequestAnimationFrame) {
			$wnd.msRequestAnimationFrame(fn);
		} else {
			$wnd.setTimeout(fn, 16);
		}
	}-*/;

	public static native void log(String message) /*-{
		console.log(message);
	}-*/;

	private void drawScene()
	{
		mFrameTime = System.currentTimeMillis() - mLastFrameTime;
		mLastFrameTime = System.currentTimeMillis();
		
		requestAnimationFrame();
		mEngine.clear();

		// Render the game
		if(mGame != null)
		{
			mGame.draw();
		}
		
		// Render the GUI
		if(mGui != null)
		{
			mGui.draw();
		}
	}
}
