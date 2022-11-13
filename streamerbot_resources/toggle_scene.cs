using System;

public class CPHInline
{
	public bool Execute()
	{
		if(!CPH.ObsIsConnected())
		{
			return false;
		}

		string status = args["status"].ToString();
		
		if(status == "show")
		{
			CPH.ObsSetSourceVisibility("Gameplay", "csgo_data", true);
		}
		else if(status == "false")
		{
			CPH.ObsSetSourceVisibility("Gameplay", "csgo_data", false);
		}

		return true;
	}
}
