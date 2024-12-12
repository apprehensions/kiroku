package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"time"

	"github.com/diamondburned/arikawa/v3/discord"
	astate "github.com/diamondburned/arikawa/v3/state"
)

var (
	token string
	cid int64
	limit uint
)

type state struct {
	*astate.State
}

func init() {
	flag.StringVar(&token, "token", "", "Discord Token")
	flag.Int64Var(&cid, "channel", 0, "Channel ID")
	flag.UintVar(&limit, "limit", 0, "messages to fetch (0 for unlimited)")
}

func main() {
	flag.Parse()

	s := state{
		State: astate.New(token),
	}

	if cid == 0 {
		fmt.Println("channel id must be given")
		os.Exit(1)
	}

	if err := s.Open(context.TODO()); err != nil {
		log.Fatal(err)
	}
	defer s.Close()

	if err := s.printChannel(discord.ChannelID(cid)); err != nil {
		log.Fatal(err)
	}
}

func (s *state) printChannel(cid discord.ChannelID) error {
	c, err := s.Channel(cid)
	if err != nil {
		return err
	}

	ms, err := s.MessagesBefore(c.ID, discord.MessageID(0), limit)
	if err != nil {
		return err
	}

	sort.Slice(ms, func(i, j int) bool {
		return ms[i].Timestamp.Time().Before(ms[j].Timestamp.Time())
	})

	for _, m := range ms {
		printMessage(&m)
	}

	return nil
}


func printMessage(m *discord.Message) {
	fmt.Printf("[%s] %s %s (%s)",
		m.ID, m.Timestamp.Format(time.DateTime), m.Author.Username, m.Author.ID)
	if m.Pinned {
		fmt.Println(" [pinned]")
	}

	switch m.Type {
	case discord.CallMessage:
		fmt.Printf("%s", " [started a call]")
	case discord.RecipientAddMessage, discord.RecipientRemoveMessage:
		if m.Type == discord.RecipientAddMessage {
			fmt.Printf("%s", "[added")
		} else {
			fmt.Printf("%s", "[removed")
		}
		for _, u := range m.Mentions {
			fmt.Printf(" %s (%s)", u.Username, u.ID)
		}
		fmt.Printf("%s", " <> group]")
	case discord.ChannelNameChangeMessage:
		fmt.Printf(" [changed group name to %s]", m.Content)
	case discord.ChannelIconChangeMessage:
		fmt.Printf(" [changed group icon]")
	case discord.ChannelPinnedMessage:
		if m.ReferencedMessage != nil {
			fmt.Printf(" [pinned message %s]", m.ReferencedMessage.ID)
		} else {
			fmt.Printf(" [pinned unknown message]")
		}
	case discord.GuildMemberJoinMessage:
		fmt.Printf(" [joined guild]")
	case discord.DefaultMessage,
		discord.InlinedReplyMessage,
		discord.ChatInputCommandMessage:
		if m.ReferencedMessage != nil {
			fmt.Printf(" [replying to message %s]", m.ReferencedMessage.ID)
		} else {
			fmt.Printf(" [replying to unknown message]")
		}

		if m.Content != "" {
			fmt.Printf(": %s", m.Content)
		}

		for _, a := range m.Attachments {
			fmt.Printf(" [attachment: %s (%s) %s %s]",
				a.Filename, a.ContentType, a.URL, a.Description)
		}
		for _, s := range m.Stickers {
			fmt.Printf(" [sticker: %s (%s)]", s.Name, s.ID)
		}
		for _, r := range m.Reactions {
			fmt.Printf(" [reaction: %d %s (%s)", r.Count, r.Emoji.Name, r.Emoji.ID)
			if r.Me {
				fmt.Printf(" reacted")
			}
			fmt.Printf("]")
		}
		for _, e := range m.Embeds {
			fmt.Printf(" [embed:")
			if e.Title != "" {
				fmt.Printf(" %s", e.Title)
			}
			if e.Description != "" {
				fmt.Printf(" %s", e.Description)
			}
			if e.Footer != nil {
				fmt.Printf(" [footer: %s]", e.Footer.Text)
			}
			if e.Image != nil {
				fmt.Printf(" [image: %s]", e.Image.URL)
			}
			if e.Thumbnail != nil {
				fmt.Printf(" [thumbnail: %s]", e.Thumbnail.URL)
			}
			if e.Video != nil {
				fmt.Printf(" [video: %s]", e.Video.URL)
			}
			for _, f := range e.Fields {
				fmt.Printf(" [field: %s = %s]", f.Name, f.Value)
			}
			fmt.Printf("]")
		}
		if m.EditedTimestamp.IsValid() {
			fmt.Printf(" [edited at %s]", m.Timestamp.Format(time.DateTime))
		}
	default:
		fmt.Printf("[unhandled message type %s]", typeNames[m.Type])
	}
	fmt.Println()
}
