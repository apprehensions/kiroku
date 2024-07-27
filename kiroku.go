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
	"github.com/diamondburned/arikawa/v3/state"
)

var typeNames = map[discord.MessageType]string{
	discord.DefaultMessage:                             "DefaultMessage",
	discord.RecipientAddMessage:                        "RecipientAddMessage",
	discord.RecipientRemoveMessage:                     "RecipientRemoveMessage",
	discord.CallMessage:                                "CallMessage",
	discord.ChannelNameChangeMessage:                   "ChannelNameChangeMessage",
	discord.ChannelIconChangeMessage:                   "ChannelIconChangeMessage",
	discord.ChannelPinnedMessage:                       "ChannelPinnedMessage",
	discord.GuildMemberJoinMessage:                     "GuildMemberJoinMessage",
	discord.NitroBoostMessage:                          "NitroBoostMessage",
	discord.NitroTier1Message:                          "NitroTier1Message",
	discord.NitroTier2Message:                          "NitroTier2Message",
	discord.NitroTier3Message:                          "NitroTier3Message",
	discord.ChannelFollowAddMessage:                    "ChannelFollowAddMessage",
	discord.GuildDiscoveryDisqualifiedMessage:          "GuildDiscoveryDisqualifiedMessage",
	discord.GuildDiscoveryRequalifiedMessage:           "GuildDiscoveryRequalifiedMessage",
	discord.GuildDiscoveryGracePeriodInitialWarning:    "GuildDiscoveryGracePeriodInitialWarning",
	discord.GuildDiscoveryGracePeriodFinalWarning:      "GuildDiscoveryGracePeriodFinalWarning",
	discord.ThreadCreatedMessage:                       "ThreadCreatedMessage",
	discord.InlinedReplyMessage:                        "InlinedReplyMessage",
	discord.ChatInputCommandMessage:                    "ChatInputCommandMessage",
	discord.ThreadStarterMessage:                       "ThreadStarterMessage",
	discord.GuildInviteReminderMessage:                 "GuildInviteReminderMessage",
	discord.ContextMenuCommand:                         "ContextMenuCommand",
	discord.AutoModerationActionMessage:                "AutoModerationActionMessage",
	discord.RoleSubscriptionPurchaseMessage:            "RoleSubscriptionPurchaseMessage",
	discord.InteractionPremiumUpsellMessage:            "InteractionPremiumUpsellMessage",
	discord.StageStartMessage:                          "StageStartMessage",
	discord.StageEndMessage:                            "StageEndMessage",
	discord.StageSpeakerMessage:                        "StageSpeakerMessage",
	discord.StageTopicMessage:                          "StageTopicMessage",
	discord.GuildApplicationPremiumSubscriptionMessage: "GuildApplicationPremiumSubscriptionMessage",
}

func main() {
	token := flag.String("token", "", "Discord Token")
	cid := flag.Int64("channel", 0, "Channel ID")
	limit := flag.Uint("limit", 0, "messages to fetch (0 for unlimited)")
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

	ms, err := s.MessagesBefore(c.ID, discord.MessageID(0), *limit)
	if err != nil {
		log.Fatal(err)
	}

	sort.Slice(ms, func(i, j int) bool {
		return ms[i].Timestamp.Time().Before(ms[j].Timestamp.Time())
	})

	for _, m := range ms {
		printMessage(&m)
	}

	s.Close()
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
		fmt.Printf(" [pinned message %s]", m.ReferencedMessage.ID)
	case discord.GuildMemberJoinMessage:
		fmt.Printf(" [joined guild]")
	case discord.DefaultMessage,
		discord.InlinedReplyMessage,
		discord.ChatInputCommandMessage:
		if m.ReferencedMessage != nil {
			fmt.Printf(" [replying to message %s]", m.ReferencedMessage.ID)
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
