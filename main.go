package main

import (
	"fmt"
	"log"

	"github.com/thecarlhall/gosubsonic"
)

func main() {
	config := LoadConfig()
	subsonic, err := gosubsonic.New(config.ServerURL, config.Username, config.Password)

	if err != nil {
		log.Fatal(err)
	}

	folders, err := subsonic.GetMusicFolders()

	if err != nil {
		log.Fatal(err)
	}

	for _, folder := range folders {
		indexes, _ := subsonic.GetIndexes(folder.ID, 0)

		for _, index := range indexes {
			fmt.Printf("-----[ %s ]-----\n", index.Name)

			for _, artist := range index.Artist {
				fmt.Printf("%+v\n", artist)
			}

			fmt.Println("")
		}
	}
}
