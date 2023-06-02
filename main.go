package main

import (
	"flag"
	"github.com/bevvvvv/music-league-monitor/discord_api"
	"github.com/bwmarrin/discordgo"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"os"
)

var (
	AddCommands = flag.Bool("add", false, "Whether to add commands or not")
)

func init() { flag.Parse() }

var (
	discordSession *discordgo.Session
	commands       = []*discordgo.ApplicationCommand{
		{
			Name:        "intimidate",
			Description: "Basic hello world for first tests.",
		},
	}
)

func main() {
	if *AddCommands {
		// create bot session
		discordSession, err := discordgo.New("Bot " + os.Getenv("DISCORD_TOKEN"))
		if err != nil {
			log.Fatalf("Invalid bot parameters: %v", err)
		}

		// open session printing session login info
		discordSession.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
			log.Printf("Logged in as: %v#%v\n", s.State.User.Username, s.State.User.Discriminator)
		})
		err = discordSession.Open()
		if err != nil {
			log.Fatalf("Cannot open the session: %v", err)
		}
		defer discordSession.Close()

		// add commands if requested
		for _, cmd := range commands {
			// blank guild id for global commands
			_, err := discordSession.ApplicationCommandCreate(discordSession.State.User.ID, "", cmd)
			if err != nil {
				log.Panicf("Cannot create '%v' command: %v", cmd.Name, err)
			}
		}
	}

	// create router
	r := chi.NewRouter()

	r.Post("/", discord_api.HandleChallenge)

	http.ListenAndServe(":80", r)
}
