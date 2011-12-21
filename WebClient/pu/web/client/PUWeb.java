package pu.web.client;

import pu.web.client.gui.GUIManager;
import pu.web.client.gui.TextField;
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
	private static WebGLRenderingContext mGlContext;
	private static PU_Engine mEngine;
	private static GUIManager mGui;
	private static PU_Resources mResources;
	private static PU_Game mGame;
	private static PU_Events mEvents;

	public void onModuleLoad()
	{
		final WebGLCanvas webGLCanvas = new WebGLCanvas(PU_Engine.SCREEN_WIDTH + "px", PU_Engine.SCREEN_HEIGHT + "px");
		PUWeb.mGlContext = webGLCanvas.getGlContext();
		PUWeb.mGlContext.viewport(0, 0, PU_Engine.SCREEN_WIDTH, PU_Engine.SCREEN_HEIGHT);
		RootPanel.get("gwtGL").add(webGLCanvas);
		
		PUWeb.mResources = new PU_Resources();
		PUWeb.mGame = new PU_Game();
		
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
			}

			@Override
			public void onFailure(Throwable reason)
			{
			}
		});

		PU_Connection connection = new PU_Connection("ws://127.0.0.1:12345/echo");
		connection.connect();
	}
	
	public static void resourcesLoaded()
	{
		PUWeb.mGui = new GUIManager(0, 0, PU_Engine.SCREEN_WIDTH, PU_Engine.SCREEN_HEIGHT, mResources.getFont(Fonts.FONT_PURITAN_BOLD_14));
		PUWeb.mEvents = new PU_Events(Document.get().getElementById("gwtGL"), PUWeb.mGui);
		
		mGame.setState(PU_Game.GAMESTATE_LOGIN);
		
		// TODO: move this to the appropriate place
		TextField tfUsername = new TextField(453, 396, 160, 20);
		tfUsername.setFontColor(57, 92, 196);
		PUWeb.mGui.getRoot().addChild(tfUsername);
		PUWeb.mGui.getRoot().focusElement(tfUsername);
		
		TextField tfPassword = new TextField(453, 424, 160, 20);
		tfPassword.setFontColor(57, 92, 196);
		tfPassword.setPassword(true);
		PUWeb.mGui.getRoot().addChild(tfPassword);
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
