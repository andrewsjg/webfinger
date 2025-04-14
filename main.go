package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/andrewsjg/webfinger/internal/mastodon"
)

func main() {
	username := flag.String("username", "", "Mastodon username to query")
	flag.Parse()

	if *username == "" {
		log.Fatal("Username must be provided")
	}

	client := mastodon.NewMastodonClient("https://mastodonapp.uk")
	profile, err := client.GetProfile(*username)
	if err != nil {
		log.Fatalf("Error fetching profile: %v", err)
	}

	fmt.Printf("Username: %s\n", profile.Username)
	fmt.Printf("Display Name: %s\n", profile.DisplayName)
	fmt.Printf("Bio: %s\n", profile.Bio)

}
