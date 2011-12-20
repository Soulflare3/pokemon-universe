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
	
	private PU_Events mEvents;
	private PU_Resources mResources;

	public void onModuleLoad()
	{
		final WebGLCanvas webGLCanvas = new WebGLCanvas(PU_Engine.SCREEN_WIDTH + "px", PU_Engine.SCREEN_HEIGHT + "px");
		PUWeb.mGlContext = webGLCanvas.getGlContext();
		PUWeb.mGlContext.viewport(0, 0, PU_Engine.SCREEN_WIDTH, PU_Engine.SCREEN_HEIGHT);
		RootPanel.get("gwtGL").add(webGLCanvas);
		
		mResources = new PU_Resources();
		
		PUWeb.mEngine = new PU_Engine(PUWeb.mGlContext);
		PUWeb.mEngine.init();

		// Start the draw loop
		drawScene();

		GWT.runAsync(new RunAsyncCallback()
		{
			@Override
			public void onSuccess()
			{
				mResources.loadFonts();
				
				PUWeb.mGui = new GUIManager(0, 0, PU_Engine.SCREEN_WIDTH, PU_Engine.SCREEN_HEIGHT, mResources.getFont(Fonts.FONT_PURITAN_BOLD_14));
				mEvents = new PU_Events(Document.get().getElementById("gwtGL"), PUWeb.mGui);
				
				TextField tf = new TextField(10, 10, 200, 40);
				tf.setBorderColor(0, 0, 0);
				tf.setFontColor(0, 255, 100);
				PUWeb.mGui.getRoot().addChild(tf);
				PUWeb.mGui.getRoot().focusElement(tf);
			}

			@Override
			public void onFailure(Throwable reason)
			{
			}
		});

		PU_Connection connection = new PU_Connection("ws://127.0.0.1:12345/echo");
		connection.connect();
	}

	public static WebGLRenderingContext gl()
	{
		return PUWeb.mGlContext;
	}

	public static PU_Engine engine()
	{
		return PUWeb.mEngine;
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

		PU_Font font = mResources.getFont(Fonts.FONT_PURITAN_BOLD_14);
		if(font != null)
		{			
			font.setColor(255, 255, 255);
			font.drawBorderedText("Test bordered text", 10, 200);
		}
		
		mEngine.setColor(255, 0, 0, 255);
		mEngine.renderLine(10, 100, 110, 100);
		
		mEngine.renderRect(10, 150, 50, 50);
		
		// Render the GUI
		if(mGui != null)
		{
			mGui.draw();
		}
	}
}
