package pu.web.client.resources.tiles;

import pu.web.shared.ResourceBundle;

import com.google.gwt.core.client.GWT;

public interface TilesBundle extends ResourceBundle
{
	static TilesBundle INSTANCE = GWT.create(TilesBundle.class);
}