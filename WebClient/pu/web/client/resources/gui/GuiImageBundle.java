package pu.web.client.resources.gui;

import com.google.gwt.core.client.GWT;
import com.google.gwt.resources.client.ClientBundle;
import com.google.gwt.resources.client.ImageResource;

public interface GuiImageBundle extends ClientBundle
{
	static GuiImageBundle INSTANCE = GWT.create(GuiImageBundle.class);
	
	@Source(value = { "loginBg.png" })
	ImageResource getLoginBg();
	
	@Source(value = { "chatpanel.png" })
	ImageResource getChatpanel();
	
	@Source(value = { "pokemangbar.png" })
	ImageResource getPokemonBar();
	
	@Source(value = { "pokemangslot.png" })
	ImageResource getPokemonSlot();
	
	@Source(value = { "hpbar_exp.png" })
	ImageResource getHpBarExp();
	
	@Source(value = { "hpbar_green.png" })
	ImageResource getHpBarGreen();
	
	@Source(value = { "hpbar_yellow.png" })
	ImageResource getHpBarYellow();
	
	@Source(value = { "hpbar_red.png" })
	ImageResource getHpBarRed();
	
	@Source(value = { "battle_background.png" })
	ImageResource getBattleBackground();
	
	@Source(value = { "battle_hpbar_exp.png" })
	ImageResource getBattleHpBarExp();
	
	@Source(value = { "battle_hpbar_green.png" })
	ImageResource getBattleHpBarGreen();
	
	@Source(value = { "battle_hpbar_yellow.png" })
	ImageResource getBattleHpBarYellow();
	
	@Source(value = { "battle_hpbar_red.png" })
	ImageResource getBattleHpBarRed();
	
	@Source(value = { "battle_pokeball_dead.png" })
	ImageResource getBattlePokeballDead();
	
	@Source(value = { "battle_pokeball_normal.png" })
	ImageResource getBattlePokeballNormal();
	
	@Source(value = { "battle_pokeball_empty.png" })
	ImageResource getBattlePokeballEmpty();
	
	@Source(value = { "battle_pokemon_enemy.png" })
	ImageResource getBattlePokemonEnemy();
	
	@Source(value = { "battle_pokemon_self.png" })
	ImageResource getBattlePokemonSelf();
}