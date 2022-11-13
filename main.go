//go:build windows
// +build windows

package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/dank/go-csgsi"
)

const steamID string = "76561198040749114"

var (
	playerstate_change bool = false
)

/*
 * TODO: figure out how to find start and end of game
 */
func main() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("Exiting...")
		sendAction("end_scene")
		os.Exit(1)
	}()

	fmt.Println("[@] Starting CSGO Data Relay")
	me := PlayerState{0, 0, 0}
	game := csgsi.New(0)
	go sendAction("start_scene") // init streamer.bot

	go func() {
		fmt.Println("Running...")
		for state := range game.Channel {
			if state.Map == nil {
				// Not in game
				continue
			}

			if state.Map.Phase == "gameover" {
				go sendUpdateEvent(PlayerState{0, 0, 0}) // GameOver reset. gets triggered twice so ignore that for now
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

			if playerstate_change { // fired only when an update is needed
				fmt.Printf("Switch stats updated %d\n", me)
				go sendUpdateEvent(me)
			}
		}
	}()
	game.Listen("127.0.0.1:3000")
}
