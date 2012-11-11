package pu.web.client.gui.impl;

import java.util.HashMap;

import pu.web.client.PUWeb;
import pu.web.client.PU_Image;
import pu.web.client.PU_Rect;
import pu.web.client.battle.PU_Battle;
import pu.web.client.gui.OnKeyDownListener;
import pu.web.client.gui.Panel;
import pu.web.client.gui.TextField;
import pu.web.client.resources.gui.GuiImages;

public class PU_ChatPanel extends Panel
{
	private PU_Image mImageBase = null;
	
	// Const
    public static final int SPEAK_NORMAL = 1;
    public static final int SPEAK_YELL = 2;
    public static final int SPEAK_WHISPER = 3;
    public static final int SPEAK_PRIVATE = 6;
	
	// Controls
	private TextField mChatInput = null;
	
	// Members
	private HashMap<Integer, PU_ChatChannel> mChatChannels = new HashMap<Integer, PU_ChatChannel>();
	private int mActiveChannel = 0;
	
	public PU_ChatPanel(int x, int y, int width, int height)
	{
		super(x, y, width, height);
		
		mImageBase = PUWeb.resources().getGuiImage(GuiImages.IMG_GUI_WORLD_BOTTOMBASE);
		
		mChatInput = new TextField(10, 694, 349, 18);
		mChatInput.setFontColor(0, 0, 0);
		mChatInput.setOnKeyDownListener(mChatInputKeydownListener);
		PUWeb.gui().getRoot().focusElement(mChatInput);
		addChild(mChatInput);
		
		addChannel(PU_ChatChannel.CHANNEL_LOCAL, "Local", false);
		addChannel(PU_ChatChannel.CHANNEL_WORLD, "World", false);
		addChannel(PU_ChatChannel.CHANNEL_TRADE, "Trade", false);
		addChannel(PU_ChatChannel.CHANNEL_BATTLE, "Battle", false);
		addChannel(PU_ChatChannel.CHANNEL_LOG, "Log", false);
		
		setActive(PU_ChatChannel.CHANNEL_LOCAL);
	}
	
	public void setActive(int id)
	{
		for(PU_ChatChannel channel : mChatChannels.values())
		{
			if(channel.getId() == id)
			{
				channel.setActive(true);
				channel.setUpdated(false);
			}
			else
			{
				channel.setActive(false);
			}
		}
	}
	
	public void addChannel(int id, String name, boolean closable)
	{
		if(!mChatChannels.containsKey(id))
		{
			PU_ChatChannel channel = new PU_ChatChannel(id, name, this);
			channel.setClosable(closable);
			mChatChannels.put(id, channel);
		}
	}
	
	public void closeChannel(int id)
	{
		PU_ChatChannel channel = mChatChannels.get(id);
		if(channel != null)
		{
			if(id == mActiveChannel)
				mActiveChannel = 0;
			
			channel.close();
			mChatChannels.remove(id);
		}
	}
	
	public void addMessage(int id, PU_Text text)
	{
		PU_ChatChannel channel = mChatChannels.get(id);
		if(channel != null)
		{
			channel.addMessage(text);
			if(id != mActiveChannel && channel.isNotifications())
			{
				channel.setUpdated(true);
			}
		}
	}
	
	private OnKeyDownListener mChatInputKeydownListener = new OnKeyDownListener()
	{	
		@Override
		public void OnKeyDown(int button)
		{
			if(button == 13)
			{				
				String message = mChatInput.getText();
				PUWeb.connection().getProtocol().sendChat(PU_ChatChannel.CHANNEL_LOCAL, 7, message);
				
				if(message.equals("/test"))
				{
					PU_Battle battle = new PU_Battle(0);
					battle.start();
				}
				
				mChatInput.setText("");
			}
		}
	};
	
	@Override
	public void draw(PU_Rect drawArea)
	{
		// Draw the base (fading white background, chat box, chat input box)
		mImageBase.draw(0, 568);
		
		super.draw(drawArea);
	}
}
