package pu.web.client.resources.gui;

import pu.web.shared.ResourceBundle;

import com.google.gwt.core.client.GWT;

public interface GuiImageBundle extends ResourceBundle
{
	static GuiImageBundle INSTANCE = GWT.create(GuiImageBundle.class);
}