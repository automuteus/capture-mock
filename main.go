package main

import (
	"encoding/json"
	"fmt"
	"github.com/automuteus/automuteus/v8/pkg/game"
	input2 "github.com/automuteus/capture-mock/input"
	"github.com/hesh915/go-socket.io-client"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	opts := &socketio_client.Options{
		Transport: "websocket",
		Query:     make(map[string]string),
	}
	host := os.Getenv("GALACTUS_HOST")
	if host == "" {
		host = "http://localhost:8123"
	}
	log.Println("Using " + host + " as Galactus Host")

	client, err := socketio_client.NewClient(host, opts)
	if err != nil {
		log.Printf("NewClient error:%v\n", err)
		return
	}
	client.On("connection", func() {
		log.Printf("on connect\n")
	})
	client.On("disconnection", func() {
		log.Fatal("Disconnected from Server, Exiting")
	})
	client.On("error", func() {
		log.Printf("on error\n")
	})
	client.On("message", func(msg string) {
		log.Printf("on message:%v\n", msg)
	})

	playerInfos := make([]game.PlayerInfo, 0)

	code := input2.String("Please provide the Room Code or capture link", "")
	if strings.HasPrefix(code, "aucapture://") {
		// only works for local testing
		code = strings.ReplaceAll(code, "aucapture://localhost:8123/", "")
		code = strings.ReplaceAll(code, "?insecure", "")
	} else {
		code = strings.ToUpper(code)
	}
	log.Println("Sending connect code: " + code)
	err = client.Emit("connectCode", code)
	if err != nil {
		log.Println(err)
	}

	log.Println("Ready for sending messages!")

	for {
		input := input2.String("Type of Message?\nL for Lobby\nS for State\nP for Player\nG for Gameover", "")
		switch strings.ToUpper(input) {
		case "L":
			l := game.Lobby{
				LobbyCode: input2.String("Lobby Code?", "TESTCODE"),
				Region:    input2.Region(),
				PlayMap:   input2.PlayMap(),
			}
			b, _ := json.Marshal(l)
			log.Println(string(b))
			client.Emit("lobby", string(b))
			time.Sleep(250)
			// lobby events automatically have an associated state change to Lobby
			client.Emit("state", fmt.Sprintf("%d", game.LOBBY))
		case "G":
			g := game.Gameover{
				GameOverReason: input2.GameResult(),
				PlayerInfos:    playerInfos,
			}
			b, _ := json.Marshal(g)
			log.Println(string(b))
			client.Emit("gameover", string(b))
			time.Sleep(250)
			// gameover events intrinsically have an associated state change to Lobby
			client.Emit("state", fmt.Sprintf("%d", game.LOBBY))
		case "P":
			p := game.Player{
				Action:       input2.PlayerAction(),
				Name:         input2.String("Player Name?", "Player"),
				Color:        input2.Color(),
				IsDead:       input2.Bool("isDead?", false),
				Disconnected: input2.Bool("disconnected?", false),
			}
			var found = false
			for _, v := range playerInfos {
				if v.Name == p.Name {
					found = true
					break
				}
			}
			if !found {
				playerInfos = append(playerInfos, game.PlayerInfo{
					Name:       p.Name,
					IsImpostor: input2.Bool("isImposter? (sent at GameOver)", true),
				})
			}
			b, _ := json.Marshal(p)
			log.Println(string(b))
			client.Emit("player", string(b))

		case "S":
			client.Emit("state", fmt.Sprintf("%d", input2.Phase()))
		default:
			log.Println("Sorry, I didn't understand the command \"" + input + "\"")
		}
	}
}
