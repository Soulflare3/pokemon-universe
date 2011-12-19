package pu.web.shared;

import com.google.gwt.resources.client.ResourcePrototype;

public interface ResourceBundle
{
	ResourcePrototype getResource(String name);
	
	ResourcePrototype[] getResources();
}