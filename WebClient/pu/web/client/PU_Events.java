package pu.web.client;

import pu.web.client.gui.GUIManager;

import com.google.gwt.core.client.JavaScriptObject;
import com.google.gwt.dom.client.Element;
import com.google.gwt.dom.client.NativeEvent;

public class PU_Events
{
	public static final int KEY_UP = 38;
	public static final int KEY_DOWN = 40;
	public static final int KEY_LEFT = 37;
	public static final int KEY_RIGHT = 39;
	
	private Element mRoot = null;
	private PU_Rect mRootRect = null;
	private GUIManager mGui = null;
	private boolean mKeyMap[] = new boolean[255];
	
	public PU_Events(Element root, GUIManager gui)
	{
		mRoot = root;
		mRootRect = new PU_Rect(0, 0, PU_Engine.SCREEN_WIDTH, PU_Engine.SCREEN_HEIGHT);
		mGui = gui;
		
		addListeners(mRoot);
	}
	
	public native void addListeners(JavaScriptObject root) /*-{
		var events = this;
		
		root.addEventListener('mousedown', function(e) {
			events.@pu.web.client.PU_Events::onMouseDown(Lcom/google/gwt/dom/client/NativeEvent;)(e);
		}, true);
		
		root.addEventListener('mouseup', function(e) {
			events.@pu.web.client.PU_Events::onMouseUp(Lcom/google/gwt/dom/client/NativeEvent;)(e);
		}, true);
		
		root.addEventListener('mousemove', function(e) {
			events.@pu.web.client.PU_Events::onMouseMove(Lcom/google/gwt/dom/client/NativeEvent;)(e);
		}, true);
		
		var scroll = 'mousewheel';
		if($wnd.navigator.userAgent.toLowerCase().indexOf('firefox') != -1)
			scroll = 'DOMMouseScroll';
		root.addEventListener(scroll, function(e) {
			events.@pu.web.client.PU_Events::onMouseScroll(Lcom/google/gwt/dom/client/NativeEvent;)(e);
		}, true);
		
		root.addEventListener('keydown', function(e) {
			events.@pu.web.client.PU_Events::onKeyDown(Lcom/google/gwt/dom/client/NativeEvent;)(e);
		}, true);
		
		root.addEventListener('keypress', function(e) {
			events.@pu.web.client.PU_Events::onKeyPress(Lcom/google/gwt/dom/client/NativeEvent;)(e);
		}, true);
		
		root.addEventListener('keyup', function(e) {
			events.@pu.web.client.PU_Events::onKeyUp(Lcom/google/gwt/dom/client/NativeEvent;)(e);
		}, true);
	}-*/;
	
	
	public final void onMouseDown(NativeEvent event)
	{
		int x =  event.getClientX() - mRoot.getAbsoluteLeft() + mRoot.getScrollLeft() + mRoot.getOwnerDocument().getScrollLeft();
		int y =  event.getClientY() - mRoot.getAbsoluteTop() + mRoot.getScrollTop() + mRoot.getOwnerDocument().getScrollTop();
		
		if(mRootRect.contains(x, y))
		{
			mGui.mouseDown(x, y);
		}
	}

	public final void onMouseUp(NativeEvent event)
	{
		int x =  event.getClientX() - mRoot.getAbsoluteLeft() + mRoot.getScrollLeft() + mRoot.getOwnerDocument().getScrollLeft();
		int y =  event.getClientY() - mRoot.getAbsoluteTop() + mRoot.getScrollTop() + mRoot.getOwnerDocument().getScrollTop();
		
		if(mRootRect.contains(x, y))
		{
			mGui.mouseUp(x, y);
		}
	}
	
	public final void onMouseMove(NativeEvent event)
	{
		int x =  event.getClientX() - mRoot.getAbsoluteLeft() + mRoot.getScrollLeft() + mRoot.getOwnerDocument().getScrollLeft();
		int y =  event.getClientY() - mRoot.getAbsoluteTop() + mRoot.getScrollTop() + mRoot.getOwnerDocument().getScrollTop();
		
		if(mRootRect.contains(x, y))
		{
			mGui.mouseMove(x, y);
		}
	}
	
	public final void onMouseScroll(NativeEvent event)
	{
		//TODO: get this stuff working
	}
	
	public final void onKeyDown(NativeEvent event)
	{
		int keycode = event.getKeyCode(); 
		if(keycode < 31)
		{
			mKeyMap[keycode] = true;
			mGui.keyDown(keycode);
			event.preventDefault();
		}
		else if(keycode >= KEY_LEFT && keycode <= KEY_DOWN)
		{
			PUWeb.game().keyDown(keycode);
			event.preventDefault();
		}
	}
	
	public final void onKeyPress(NativeEvent event)
	{
		event.preventDefault();
		 
		int charCode = event.getCharCode();
		if(charCode != 0 && charCode > 31)
		{
			mKeyMap[charCode] = true;
			mGui.textInput(charCode);
		}
	}

	
	public final void onKeyUp(NativeEvent event)
	{
		event.preventDefault();
		
		int keycode = event.getKeyCode(); 
		mKeyMap[keycode] = false;
		mGui.keyUp(keycode);
	}
	
	public boolean isKeyDown(int keycode)
	{
		return mKeyMap[keycode];
	}
}
