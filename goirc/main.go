package main

import (
	"crypto/tls"
	"fmt"
	irc "github.com/fluffle/goirc/client"
	"log"
	"strings"
)

func main() {

	cfg := irc.NewConfig("bot")
	cfg.SSL = true
	cfg.SSLConfig = &tls.Config{
		InsecureSkipVerify: true,
	}
	cfg.Server = "irc.ozinger.org:16667"
	cfg.NewNick = func(n string) string { return n + "^" }

	c := irc.Client(cfg)

	// Add handlers to do things here!
	// e.g. join a channel on connect.

	c.HandleFunc("connected",
		func(conn *irc.Conn, line *irc.Line) { conn.Join("#codeport") })

	// And a signal on disconnect
	quit := make(chan bool)
	c.HandleFunc("disconnected",
		func(conn *irc.Conn, line *irc.Line) { quit <- true })

	c.HandleFunc(irc.JOIN,
		func(conn *irc.Conn, line *irc.Line) {
			conn.Privmsg("#codeport", "안녕_하세_요_")
		})

	c.HandleFunc(irc.PRIVMSG,
		func(conn *irc.Conn, line *irc.Line) {
			if strings.HasPrefix(line.Text(), "bot") {
				command := strings.Split(line.Text(), " ")[1]
				switch command {
				case "quit":
					log.Println("Received command to quit")
					conn.Quit("안녕_")
				}

			}

		})

	// Tell client to connect.
	if err := c.Connect(); err != nil {
		fmt.Printf("Connection error: %s\n", err.Error())
	}

	// Wait for disconnect
	<-quit

}
