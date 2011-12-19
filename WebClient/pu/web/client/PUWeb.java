package pu.web.client;

import pu.web.client.gui.GUIManager;

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
	
	private PU_Events mEvents;
	private GUIManager mGui;

	public void onModuleLoad()
	{
		final WebGLCanvas webGLCanvas = new WebGLCanvas(PU_Engine.SCREEN_WIDTH + "px", PU_Engine.SCREEN_HEIGHT + "px");
		PUWeb.mGlContext = webGLCanvas.getGlContext();
		PUWeb.mGlContext.viewport(0, 0, PU_Engine.SCREEN_WIDTH, PU_Engine.SCREEN_HEIGHT);
		RootPanel.get("gwtGL").add(webGLCanvas);
		
		mGui = new GUIManager(0, 0, PU_Engine.SCREEN_WIDTH, PU_Engine.SCREEN_HEIGHT, null/*default font*/);
		mEvents = new PU_Events(Document.get().getElementById("gwtGL"), mGui);

		PUWeb.mEngine = new PU_Engine(PUWeb.mGlContext);
		PUWeb.mEngine.init();

		// Start the draw loop
		drawScene();

		GWT.runAsync(new RunAsyncCallback()
		{
			@Override
			public void onSuccess()
			{
				// Load our tiles, fonts etc.
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

		/* drawing logic */
		log("derp");
	}
}
