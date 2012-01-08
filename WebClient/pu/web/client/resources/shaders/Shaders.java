package pu.web.client.resources.shaders;

import com.google.gwt.core.client.GWT;
import com.google.gwt.resources.client.ClientBundle;
import com.google.gwt.resources.client.TextResource;

public interface Shaders extends ClientBundle
{
	public static Shaders INSTANCE = GWT.create(Shaders.class);

	@Source(value = { "fragment_solid.txt" })
	TextResource fragmentShaderSolid();
	
	@Source(value = { "fragment_tex.txt" })
	TextResource fragmentShaderTex();
	
	@Source(value = { "fragment_sprite.txt" })
	TextResource fragmentShaderSprite();
	
	@Source(value = { "fragment_atlas.txt" })
	TextResource fragmentShaderAtlas();

	@Source(value = { "vertex.txt" })
	TextResource vertexShader();
}