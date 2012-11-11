package pu.web.client.resources.gui;

import com.google.gwt.resources.client.ImageResource;

public class GuiImages
{
	private static ImageResource[] mImages;
	
	public static final int IMG_GUI_INTROBG = 0;
	public static final int IMG_GUI_WORLD_BOTTOMBASE = 1;
	public static final int IMG_GUI_WORLD_POKEMONBAR = 2;
	public static final int IMG_GUI_WORLD_POKEMONSLOT = 3;
	public static final int IMG_GUI_WORLD_HPBAR_EXP = 4;
	public static final int IMG_GUI_WORLD_HPBAR_GREEN = 5;
	public static final int IMG_GUI_WORLD_HPBAR_YELLOW = 6;
	public static final int IMG_GUI_WORLD_HPBAR_RED = 7;
	public static final int IMG_GUI_BATTLE_BACKGROUND = 8;
	public static final int IMG_GUI_BATTLE_HPBAR_EXP = 9;
	public static final int IMG_GUI_BATTLE_HPBAR_GREEN = 10;
	public static final int IMG_GUI_BATTLE_HPBAR_YELLOW = 11;
	public static final int IMG_GUI_BATTLE_HPBAR_RED = 12;
	public static final int IMG_GUI_BATTLE_POKEBALL_DEAD = 13;
	public static final int IMG_GUI_BATTLE_POKEBALL_EMPTY = 14;
	public static final int IMG_GUI_BATTLE_POKEBALL_NORMAL = 15;
	public static final int IMG_GUI_BATTLE_POKEMON_ENEMY = 16;
	public static final int IMG_GUI_BATTLE_POKEMON_SELF = 17;
	
	static {
		mImages = new ImageResource[] {
				GuiImageBundle.INSTANCE.getLoginBg(),
				GuiImageBundle.INSTANCE.getChatpanel(),
				GuiImageBundle.INSTANCE.getPokemonBar(),
				GuiImageBundle.INSTANCE.getPokemonSlot(),
				GuiImageBundle.INSTANCE.getHpBarExp(),
				GuiImageBundle.INSTANCE.getHpBarGreen(),
				GuiImageBundle.INSTANCE.getHpBarYellow(),
				GuiImageBundle.INSTANCE.getHpBarRed(),
				GuiImageBundle.INSTANCE.getBattleBackground(),
				GuiImageBundle.INSTANCE.getBattleHpBarExp(),
				GuiImageBundle.INSTANCE.getBattleHpBarGreen(),
				GuiImageBundle.INSTANCE.getBattleHpBarYellow(),
				GuiImageBundle.INSTANCE.getBattleHpBarRed(),
				GuiImageBundle.INSTANCE.getBattlePokeballDead(),
				GuiImageBundle.INSTANCE.getBattlePokeballEmpty(),
				GuiImageBundle.INSTANCE.getBattlePokeballNormal(),
				GuiImageBundle.INSTANCE.getBattlePokemonEnemy(),
				GuiImageBundle.INSTANCE.getBattlePokemonSelf()
		};
	}
	
	public static ImageResource[] getImages()
	{
		return mImages;
	}
}
