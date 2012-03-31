package pu.web.client.gui.impl;

import pu.web.client.PUWeb;
import pu.web.client.PU_Image;
import pu.web.client.PU_Player;
import pu.web.client.PU_Pokemon;
import pu.web.client.PU_Rect;
import pu.web.client.gui.Panel;
import pu.web.client.resources.gui.GuiImages;

import com.googlecode.gwtgl.binding.WebGLTexture;

public class PU_WorldPanel extends Panel
{
	public PU_WorldPanel(int x, int y, int width, int height)
	{
		super(x, y, width, height);
	}
	
	@Override
	public void draw(PU_Rect drawArea)
	{
		PU_Player self = PUWeb.game().getSelf();
		
		// Pokemon bar
		PU_Image pokemonBar = PUWeb.resources().getGuiImage(GuiImages.IMG_GUI_WORLD_POKEMONBAR);
		if(pokemonBar != null)
		{
			pokemonBar.draw(193, 11);
		}
		
		// The following routine might not seem very efficient, 
		// but since webgl is the slowest element in the program - this is the best thing to do
		PU_Image pokemonSlot = PUWeb.resources().getGuiImage(GuiImages.IMG_GUI_WORLD_POKEMONSLOT);
		if(pokemonSlot != null)
		{
			// First draw all slots (to avoid texture switching)
			for(int i = 0; i < 6; i++)
			{
				pokemonSlot.draw(201+(i*(pokemonSlot.getWidth()+2)), 14);
			}
			
			if(self != null)
			{
				// Now draw our pokemon icons in a batch
				int pokemonCount = self.getPokemonCount();
				if(pokemonCount > 0)
				{
					WebGLTexture pokemonTexture = PUWeb.resources().getPokemonTexture();
					if(pokemonTexture != null)
					{
						PUWeb.engine().beginTextureBatch(pokemonTexture, 2048, pokemonCount, 255, 255, 255, 255);
						for(int i = 0; i < 6; i++)
						{
							PU_Pokemon pokemon = self.getPokemon(i);
							if(pokemon != null)
							{
								PU_Image pokemonIcon =  PUWeb.resources().getPokemonIcon(pokemon.getSpeciesId());
								if(pokemonIcon != null)
								{
									pokemonIcon.draw(204+(i*(pokemonSlot.getWidth()+2)), 8, true);
								}
								
								PU_Image hpbarImage = null;
								float hpperc = ((float)pokemon.getHp() / (float)pokemon.getHpmax());
								if(hpperc > 0.6f)
								{
									hpbarImage = PUWeb.resources().getGuiImage(GuiImages.IMG_GUI_WORLD_HPBAR_GREEN);
								}
								else if(hpperc > 0.2f)
								{
									hpbarImage = PUWeb.resources().getGuiImage(GuiImages.IMG_GUI_WORLD_HPBAR_YELLOW);
								}
								else
								{
									hpbarImage = PUWeb.resources().getGuiImage(GuiImages.IMG_GUI_WORLD_HPBAR_RED);
								}
								if(hpbarImage != null)
								{
									hpbarImage.drawRect(new PU_Rect(256+(i*(pokemonSlot.getWidth()+2)), 19, (int)Math.ceil((float)hpperc * hpbarImage.getWidth()), hpbarImage.getHeight()));
								}
								
								PU_Image expbarImage = PUWeb.resources().getGuiImage(GuiImages.IMG_GUI_WORLD_HPBAR_EXP);
								if(expbarImage != null)
								{
									expbarImage.drawRect(new PU_Rect(256+(i*(pokemonSlot.getWidth()+2)), 29, (int)Math.ceil(((float)pokemon.getExpPerc()/100.0f) * expbarImage.getWidth()), expbarImage.getHeight()));
								}
							}
						}
						PUWeb.engine().endTextureBatch();	
					}
				}
			}
		}
		
		super.draw(drawArea);
	}
}
