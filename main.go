package main

import (
	"bufio"
	"encoding/json"
	"github.com/automuteus/automuteus/v7/pkg/game"
	"github.com/hesh915/go-socket.io-client"
	"log"
	"os"
	"strings"
)

func main() {
	opts := &socketio_client.Options{
		Transport: "websocket",
		Query:     make(map[string]string),
	}
	uri := "http://localhost:8123/"

	client, err := socketio_client.NewClient(uri, opts)
	if err != nil {
		log.Printf("NewClient error:%v\n", err)
		return
	}
	client.On("connection", func() {
		log.Printf("on connect\n")
	})
	client.On("disconnection", func() {
		log.Printf("on disconnect\n")
	})
	client.On("error", func() {
		log.Printf("on error\n")
	})
	client.On("message", func(msg string) {
		log.Printf("on message:%v\n", msg)
	})

	log.Println("Please provide the Room Code or capture link:")
	reader := bufio.NewReader(os.Stdin)
	data, _, _ := reader.ReadLine()
	command := string(data)
	if strings.HasPrefix(command, "aucapture://") {
		// only works for local testing
		command = strings.ReplaceAll(command, "aucapture://localhost:8123/", "")
		command = strings.ReplaceAll(command, "?insecure", "")
	} else {
		command = strings.ToUpper(command)
	}
	log.Println("Sending connect code: " + command)
	err = client.Emit("connectCode", command)
	if err != nil {
		log.Println(err)
	}

	log.Println("Ready for sending messages!\nType L for Lobby\nS for state\nP for player\nG for gameover")

	for {
		data, _, _ := reader.ReadLine()
		switch strings.ToUpper(string(data)) {
		case "L":
			l := game.Lobby{
				LobbyCode: "TESTCODE",
				Region:    game.NA,
				PlayMap:   game.SKELD,
			}
			b, _ := json.Marshal(&l)
			log.Println(string(b))
			client.Emit("lobby", string(b))
		case "G":
			client.Emit("gameover")

			// TODO implement state and player
		}
	}
}
