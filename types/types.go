package types

import (
	dgo "github.com/bwmarrin/discordgo"
)

// Handler holds a refrence to the handler function for a application command
type Handler struct {
	Name     string
	Callback func(i *dgo.InteractionCreate, s *dgo.Session)
}

// Subcommand holds a refrence to the handler function for a application subcommand
type Subcommand struct {
	Name     string
	Callback func(parms SubcommandParms)
}

// SubcommandParms holds all neccesary parms to invoke a subcommand
type SubcommandParms struct {
	Interaction *dgo.InteractionCreate
	Session     *dgo.Session
	Option      *dgo.ApplicationCommandInteractionDataOption
}