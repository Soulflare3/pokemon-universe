package pu.web.client.resources.pokemon;

import com.google.gwt.core.client.GWT;
import com.google.gwt.resources.client.ClientBundle;
import com.google.gwt.resources.client.ImageResource;
import com.google.gwt.resources.client.TextResource;

public interface Pokemon extends ClientBundle
{
	public static Pokemon INSTANCE = GWT.create(Pokemon.class);

	@Source(value = { "pokemang.xml" })
	TextResource getPokemonInfo();
	
	@Source(value = { "pokemang.png" })
	ImageResource getPokemonBitmap();
}