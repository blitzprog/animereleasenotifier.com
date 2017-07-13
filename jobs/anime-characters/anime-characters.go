package main

import (
	"fmt"
	"time"

	"github.com/animenotifier/arn"
	"github.com/fatih/color"
)

func main() {
	color.Yellow("Refreshing anime characters...")

	rateLimiter := time.NewTicker(500 * time.Millisecond)

	for anime := range arn.MustStreamAnime() {
		<-rateLimiter.C

		chars, err := anime.RefreshAnimeCharacters()

		if err != nil {
			color.Red(err.Error())
			continue
		}

		fmt.Printf("%s %s (%d characters)\n", anime.ID, anime.Title.Canonical, len(chars.Items))
	}

	color.Green("Finished.")
}
