package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"time"

	itunessearch "github.com/mattn/itunes-search-api"
)

func main() {
	var country string
	var titles []string
	flag.StringVar(&country, "country", "JP", "country")

	if len(titles) == 0 {
		titles = append(titles, "マツケンサンバ")
	}

	flag.Parse()

	rand.Seed(time.Now().Unix())
	played := make([]string, 10)
	for {
		for _, title := range titles {
			resp, err := itunessearch.Search(title, country, "music")
			if err != nil {
				continue
			}
			results := resp.Results

			// shuffle
			n := len(results)
			for i := n - 1; i >= 0; i-- {
				j := rand.Intn(i + 1)
				results[i], results[j] = results[j], results[i]
			}
			newIndex := -1
			for _, result := range results {
				for i, p := range played {
					if p != result.PreviewUrl {
						newIndex = i
					}
				}
			}
			if newIndex != -1 {
				result := results[newIndex]
				fmt.Printf("%s: %s\n%s\n", result.ArtistName, result.TrackName, result.CollectionViewUrl)
				err = playURL(result.PreviewUrl)
				if err != nil {
					log.Println(err)
				}
				played, played[0] = append(played[:1], played[0:]...), result.PreviewUrl
			}
		}
	}
}
