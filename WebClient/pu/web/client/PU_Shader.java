package pu.web.client;

import com.googlecode.gwtgl.binding.WebGLProgram;
import com.googlecode.gwtgl.binding.WebGLUniformLocation;

public class PU_Shader
{
	private WebGLUniformLocation mUTexture;
	private WebGLUniformLocation mUModulation;
	private WebGLUniformLocation mUProjection;
	private WebGLUniformLocation mUColor;
	
	private int mAPosition = 0;
	private int mATexCoord = 0;
	
	private WebGLProgram mProgram;
	
	public PU_Shader(WebGLProgram program)
	{
		mProgram = program;
		
		mUTexture = PUWeb.gl().getUniformLocation(program, "u_texture");
		mUModulation = PUWeb.gl().getUniformLocation(program, "u_modulation");
		mUProjection = PUWeb.gl().getUniformLocation(program, "u_projection");
		mUColor = PUWeb.gl().getUniformLocation(program, "u_color");
		
		mAPosition = PUWeb.gl().getAttribLocation(program,"a_position");
		mATexCoord = PUWeb.gl().getAttribLocation(program,"a_texCoord");
	}
	
	public WebGLProgram getProgram()
	{
		return mProgram;
	}
	
	public WebGLUniformLocation getUTexture()
	{
		return mUTexture;
	}
	
	public WebGLUniformLocation getUModulation()
	{
		return mUModulation;
	}
	
	public WebGLUniformLocation getUProjection()
	{
		return mUProjection;
	}
	
	public WebGLUniformLocation getUColor()
	{
		return mUColor;
	}
	
	public int getAPosition()
	{
		return mAPosition;
	}
	
	public int getATexCoord()
	{
		return mATexCoord;
	}
}
