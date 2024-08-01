package main

import (
	"github.com/diamondburned/arikawa/v3/discord"
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
