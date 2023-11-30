package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/diamondburned/arikawa/v3/state"
)

func main() {
	const messagesFetchLimit = 1024

	token := flag.String("token", "", "Discord Token")
	cid := flag.Int64("channel", 0, "Channel ID")
	flag.Parse()

	s := state.New(*token)

	if *cid == 0 {
		fmt.Println("channel id must be given")
		os.Exit(1)
	}

	if err := s.Open(context.TODO()); err != nil {
		log.Fatal(err)
	}

	c, err := s.Channel(discord.ChannelID(*cid))
	if err != nil {
		log.Fatal(err)
	}

	var lastMessageID discord.MessageID

	for {
		ms, err := s.MessagesBefore(c.ID, lastMessageID, messagesFetchLimit)
		if err != nil {
			log.Fatal(err)
		}

		for _, m := range ms {
			printMessage(&m)
		}

		if len(ms) > 0 {
			break
		}

		lastMessageID = ms[len(ms)-1].ID
	}

	s.Close()
}

func printMessage(m *discord.Message) {
	name := m.Author.DisplayName
	if name == "" {
		name = m.Author.Username
	}

	fmt.Printf("[%s] %s %s (%s): ", m.ID, m.Timestamp.Format(time.DateTime), name, m.Author.ID)
	switch m.Type {
	case discord.CallMessage:
		fmt.Printf("%s", "[started a call]")
	case discord.DefaultMessage, discord.InlinedReplyMessage:
		if m.ReferencedMessage != nil {
			fmt.Printf("[replying to message %s] ", m.ReferencedMessage.ID)
		}

		fmt.Printf("%s", m.Content)

		for _, a := range m.Attachments {
			fmt.Printf(" [attachment: %s %s %s]", a.Filename, a.URL, a.Description)
		}
		for _, s := range m.Stickers {
			fmt.Printf(" [sticker: %s %s]", s.Name, s.ID)
		}
		for _, r := range m.Reactions {
			fmt.Printf(" [reaction: %d %s]", r.Count, r.Emoji.Name)
		}
		for _, e := range m.Embeds {
			fmt.Printf(" [embed: %s %s]", e.Title, e.Description)
		}
		if m.EditedTimestamp.IsValid() {
			fmt.Printf(" [edited at %s]", m.Timestamp.Format(time.DateTime))
		}
	default:
		fmt.Printf("[unhandled message type %d]", m.Type)
	}
	fmt.Println()
}
