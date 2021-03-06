package pu.web.client.gui.impl;

import pu.web.client.gui.Container;
import pu.web.client.gui.Scrollbar;

public class PU_ChatChannel
{
	public static final int CHANNEL_LOG = -2;
	public static final int CHANNEL_IRC = -1;
	public static final int CHANNEL_LOCAL = 0;
	public static final int CHANNEL_WORLD = 1;
	public static final int CHANNEL_TRADE = 2;
	public static final int CHANNEL_BATTLE = 3;
	
	private int mId;
	private String mName;
	
	private boolean mClosable = true;
	private boolean mUpdated = false;
	private boolean mNotifications = true;
	private boolean mGameChannel = false;
	
	private Scrollbar mScrollbar;
	private PU_Chatbox mChatbox;
	
	public PU_ChatChannel(int id, String name, Container container)
	{
		mId = id;
		
		switch(id)
		{
		case CHANNEL_WORLD:
		case CHANNEL_TRADE:
		case CHANNEL_BATTLE:
		case CHANNEL_IRC:
		case CHANNEL_LOG:
			mGameChannel = true;
		}
		
		mChatbox = new PU_Chatbox(13, 571, 350, 110);	
		mChatbox.setVisible(false);
		container.addChild(mChatbox);		
					
		mScrollbar = new Scrollbar(372, 573, 13, 103, Scrollbar.SCROLLBAR_VERTICAL);
		mScrollbar.setVisible(false);
		mScrollbar.getScroller().setVisible(false);
		mChatbox.setScrollbar(mScrollbar);		
		container.addChild(mScrollbar);
	}
	
	public int getId()
	{
		return mId;
	}

	public void setId(int id)
	{
		mId = id;
	}

	public String getName()
	{
		return mName;
	}

	public void setName(String name)
	{
		mName = name;
	}

	public boolean isClosable()
	{
		return mClosable;
	}

	public void setClosable(boolean closable)
	{
		mClosable = closable;
	}

	public boolean isUpdated()
	{
		return mUpdated;
	}

	public void setUpdated(boolean updated)
	{
		mUpdated = updated;
	}

	public boolean isNotifications()
	{
		return mNotifications;
	}

	public void setNotifications(boolean notifications)
	{
		mNotifications = notifications;
	}

	public boolean isGameChannel()
	{
		return mGameChannel;
	}

	public void setGameChannel(boolean gamechannel)
	{
		mGameChannel = gamechannel;
	}
	
	public void addMessage(PU_Text text)
	{
		if(mChatbox != null)
		{
			int fontHeight = mChatbox.getFont().getLineHeight();
			int boxHeight = mChatbox.getRect().height - 6;
			int visibleLines = (int)((float)boxHeight / (float)fontHeight);
			if(mChatbox.getLineCount()+1 > visibleLines && mScrollbar != null)
				mScrollbar.getScroller().setVisible(true);
			
			mChatbox.addText(text);
		}
	}
	
	public void close()
	{
		if(mChatbox != null)
		{
			Container parent = (Container)mChatbox.getParent(); 
			if(parent != null)
			{
				parent.removeChild(mChatbox);
			}
		}
		
		if(mScrollbar != null)
		{
			Container parent = (Container)mScrollbar.getParent(); 
			if(parent != null)
			{
				parent.removeChild(mScrollbar);
			}
		}
	}
	
	public void setActive(boolean active)
	{
		if(mChatbox != null && mScrollbar != null)
		{
			if(active)
			{
				mChatbox.setVisible(true);
				mScrollbar.setVisible(true);
			}
			else
			{
				mChatbox.setVisible(false);
				mScrollbar.setVisible(false);
			}
		}
	}
}
