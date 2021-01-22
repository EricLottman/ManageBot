import (
    "fmt"

    "github.com/AJGherardi/ManageBot/utils"
    dgo "github.com/bwmarrin/discordgo"
)

// HandleStats handles a stat command
func HandleStats(i *dgo.InteractionCreate, s dgo.Session) {
    // Get members list
    members, _ := s.GuildMembers(i.GuildID, "", 100)
    // Get bot members
    bots := []dgo.Member{}
    for , member := range members {
        if member.User.Bot {
            bots = append(bots, member)
        }
    }
    // Get presence
    presences := []*dgo.Presence{}
    for , member := range members {
        userPresence,  := s.State.Presence(i.GuildID, member.User.ID)
        if userPresence != nil {
            presences = append(presences, userPresence)
        }
    }
    // Check how many are online start from one to count itself
    online := 1
    for , precence := range presences {
        if precence.Status == dgo.StatusOnline {
            online++
        }
    }
    // Get guild for PremiumSubscriptionCount
    guild, _ := s.Guild(i.GuildID)
    // Sends stats
    utils.SendResponse("There are "+fmt.Sprint(len(members))+" many members "+fmt.Sprint(online)+" of which are online", i, s)
    utils.SendResponse("There are "+fmt.Sprint(len(bots))+" many bots", i, s)
    utils.SendResponse("There are "+fmt.Sprint(guild.PremiumSubscriptionCount)+" many people boosting your server", i, s)
}