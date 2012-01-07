package pu.web.client.resources.tiles;

import com.google.gwt.core.client.GWT;
import com.google.gwt.resources.client.ClientBundle;
import com.google.gwt.resources.client.ImageResource;
import com.google.gwt.resources.client.TextResource;

public interface Tiles extends ClientBundle
{
	public static Tiles INSTANCE = GWT.create(Tiles.class);

	@Source(value = { "putiles.xml" })
	TextResource getTilesInfo();
	
	@Source(value = { "putiles.png" })
	ImageResource getTilesBitmap();
}