package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/slack-go/slack"
)

var (
	channelID       = flag.String("channelid", "", "The ID of the channel posting to.")
	messageTemplate = "Hey <@%s>! Its time for you to break the ice.\n\n Go to %s, select a juicy question and post in this channel to get the conversation started."
	url             = flag.String("url", "https://museumhack.com/list-icebreakers-questions/", "The url to retrieve questions from.")
	token           = flag.String("token", "", "The API token for authenticating with the slack API.")
)

func main() {
	flag.Parse()

	// create the client
	api := slack.New(*token)

	// get the list of users from the channel
	users, _, err := api.GetUsersInConversation(
		&slack.GetUsersInConversationParameters{
			ChannelID: *channelID,
		},
	)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	// structs to hold filtered user list
	validUsers := make([]struct {
		Name string
		ID   string
	}, 0, len(users))

	// loop thru users and check that they are not bots
	// + not set to away then add to validUsers slice
	for _, val := range users {
		user, err := api.GetUserInfo(val)
		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}
		// check user presence
		pres, _ := api.GetUserPresence(val)
		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}

		if user.IsBot == false && pres.Presence == "active" {
			validUsers = append(validUsers, struct {
				Name string
				ID   string
			}{Name: user.Name, ID: user.ID})
		}

	}
	// if no suitable users found exit
	if len(validUsers) == 0 {
		log.Printf("No suitable users in channel %s.", *channelID)
		return
	}

	// randomly pick a user from the non-bot users
	rand.Seed(time.Now().Unix())
	randomIndex := rand.Intn(len(validUsers))
	pick := validUsers[randomIndex]

	log.Printf("Selected user %s from channel %s.", pick.Name, *channelID)

	// post message to channel mentioning the user picked
	api.PostMessage(
		*channelID,
		slack.MsgOptionUser(pick.ID),
		slack.MsgOptionText(fmt.Sprintf(messageTemplate, pick.Name, *url), false),
	)

}
