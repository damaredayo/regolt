package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/damaredayo/regolt"
)

func main() {
	s := regolt.CreateBot("TOKEN")

	err := s.Open()
	if err != nil {
		log.Fatalln(err)
	}

	s.AddHandler(OnReady)
	s.AddHandler(OnMessage)

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}

func OnReady(s *regolt.Session, r *regolt.Ready) {
	fmt.Printf("Bot ready!\nUsers:%v\nServers:%v\nChannels:%v\n",
		len(r.Users), len(r.Servers), len(r.Channels))
}

func OnMessage(s *regolt.Session, m *regolt.MessageCreate) {
	if m.Message.Content == "!ping" {
		resp, err := s.ChannelMessageSendString(m.Message.Channel, "Pong!")
		fmt.Println(string(resp), err)
	}
}
