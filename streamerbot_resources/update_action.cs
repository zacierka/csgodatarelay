using System;

public class CPHInline
{
	public bool Execute()
	{
		
		if(!CPH.ObsIsConnected())
		{
			return false;
		}

		string action = args["action"].ToString();	
		string data = args["value"].ToString();

		updateData(action, data);

		return true;
	}

    private void updateData(string action, string data)
    {
		bool scene_vis = CPH.ObsIsSourceVisible("Gameplay", "csgo_data");
		
		if((CPH.ObsGetCurrentScene() == "Gameplay") && scene_vis)
		{
			if(action == "kill")
			{
				CPH.ObsSetGdiText("csgo_data", "kills", data);
			}
			else if (action = "death")
			{
				CPH.ObsSetGdiText("csgo_data", "deaths", data);
			}
			else if(action == "assist")
			{
				CPH.ObsSetGdiText("csgo_data", "assists", data);
			}
			
		}
    }
}
