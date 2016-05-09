package main

import (
	"fmt"
	"log"

	"github.com/thecarlhall/gosubsonic"
)

type Printer struct {
	indexes  map[string]Artists
	subsonic *gosubsonic.Client
}

func NewPrinter(indexes map[string]Artists, subsonic *gosubsonic.Client) *Printer {
	return &Printer{
		indexes:  indexes,
		subsonic: subsonic,
	}
}

func (p *Printer) PrintArtistsForPrefix(input Input) {
	fmt.Printf("-----[ %s ]-----\n", input.Index)

	artists := p.indexes[input.Index]
	for i, artist := range artists {
		fmt.Printf("[%d] %s\n", i, artist.Name)
	}

	fmt.Println("")
}

func (p *Printer) PrintArtist(inputs []Input) {
	if len(inputs) < 2 {
		return
	}

	fmt.Printf("-----[ %s ]-----\n", inputs[0].Index)

	artists := p.indexes[inputs[0].Index]
	artist := artists[inputs[1].Idx]
	fmt.Printf("Artist: %s\n", artist.Name)

	content, err := p.subsonic.GetMusicDirectory(artist.ID)
	if err != nil {
		log.Fatal(err)
	}

	if len(content.Directories) > 0 {
		fmt.Println("[[ Directories ]]")
		for i, dir := range content.Directories {
			fmt.Printf("=> [%d] %s\n", i, dir.Title)
		}
	}

	if len(content.Audio) > 0 {
		fmt.Println("[[ Audio ]]")
		for i, audio := range content.Audio {
			fmt.Printf("=> [%d] %s\n", i, audio.Title)
		}
	}

	fmt.Println("")
}

func (p *Printer) PrintDirectory(inputs []Input) {
	if len(inputs) < 3 {
		return
	}

	artists := p.indexes[inputs[0].Index]
	artist := artists[inputs[1].Idx]
	content, err := p.subsonic.GetMusicDirectory(artist.ID)
	if err != nil {
		log.Fatal(err)
	}

	idx := inputs[2].Idx
	if idx < len(content.Directories) {
		dir := content.Directories[idx]
		musicDir, err := p.subsonic.GetMusicDirectory(dir.ID)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("=> %+v\n", musicDir)
	} else {
		idx = idx - len(content.Directories)
		dir := content.Audio[idx]
		fmt.Printf("=> %s\n", dir.Title)
	}
}

func (p *Printer) PrintIndex(inputs []Input) {
	if len(inputs) >= 1 && inputs[0].IndexType == ALPHA {
		if len(inputs) >= 2 && inputs[1].IndexType == NUMBER {
			if len(inputs) >= 3 && inputs[2].IndexType == NUMBER {
				p.PrintDirectory(inputs)
			} else {
				p.PrintArtist(inputs)
			}
		} else {
			p.PrintArtistsForPrefix(inputs[0])
		}
	}
}
