package main

import (
	"fmt"

	"github.com/dank/go-csgsi"
)

const steamID string = "76561198040749114"

var (
	playerstate_change bool = false
	gameState          bool = false // live, over, freezetime
	// need var to count rounds
)

func main() {
	me := PlayerState{0, 0, 0}
	game := csgsi.New(0)

	go func() {
		for state := range game.Channel {
			if ingame := state.Player.State; ingame == nil { // Not in game
				gameState = false
				continue
			}

			if state.Player.State != nil && gameState !=  {
				gameState = true
			}

			playerstate_change = false
			if state.Player.SteamId != steamID {
				continue
			}

			kills := state.Player.Match_stats.Kills
			deaths := state.Player.Match_stats.Deaths
			assists := state.Player.Match_stats.Assists

			if kills != me.kills {
				me.kills = kills
				playerstate_change = true
			}

			if deaths != me.deaths {
				me.deaths = deaths
				playerstate_change = true
			}

			if assists != me.assists {
				me.assists = assists
				playerstate_change = true
			}

			if playerstate_change {
				// send to streamer.bot
				fmt.Printf("Switch stats updated %d\n", me)
			}
		}
	}()

	game.Listen(":3000")
}
