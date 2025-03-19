package main

import (
	"flag"
	"log"

	"github.com/bwmarrin/discordgo"
)

var BotToken = flag.String("token", "", "Bot access token")
var s *discordgo.Session

func init() { flag.Parse() }

func init() {
	var err error
	s, err = discordgo.New("Bot " + *BotToken)
	if err != nil {
		log.Fatalf("Invalid bot parameters: %v", err)
	}
}

var (
	integerOptionMinValue          = 1.0
	dmPermission                   = false
	defaultMemberPermissions int64 = discordgo.PermissionManageServer

	commands = []*discordgo.ApplicationCommand{
		{
			Name:        "basic-command",
			Description: "basic test command",
		},
		{
			Name:                     "permission-overview",
			Description:              "Command for demonstrating of default command permissions",
			DefaultMemberPermissions: &defaultMemberPermissions,
			DMPermission:             &dmPermission,
		},
	}
)
