using System;

public class CPHInline
{
	public bool Execute()
	{
		if(!CPH.ObsIsConnected())
		{
			return false;
		}

		string kill = args["kills"].ToString();
		string death = args["deaths"].ToString();
		string assist = args["assists"].ToString();

		updateData(kill, death, assist);

		return true;
	}

    private void updateData(string kill, string death, string assist)
    {
		bool scene_vis = CPH.ObsIsSourceVisible("Gameplay", "csgo_data");
		
		if((CPH.ObsGetCurrentScene() == "Gameplay") && scene_vis)
		{
			CPH.ObsSetGdiText("csgo_data", "kills", kill);
			CPH.ObsSetGdiText("csgo_data", "deaths", death);
			CPH.ObsSetGdiText("csgo_data", "assists", assist);
		}
    }
}