package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/thecarlhall/gosubsonic"
)

const (
	// ALPHA represents the type 'alpha' for the input character
	ALPHA = "alpha"
	// NUMBER represents the type 'number' for the input character
	NUMBER = "number"
)

// Artists is an alias for easier typing
type Artists []gosubsonic.IndexArtist

// Input is the input we accept from a user
type Input struct {
	Index     string
	IndexType string
	Idx       int
}

func getClient(config *Config) *gosubsonic.Client {
	subsonic, err := gosubsonic.New(config.ServerURL, config.Username, config.Password)

	if err != nil {
		log.Fatal(err)
	}

	return subsonic
}

func loadIndexes(subsonic *gosubsonic.Client) map[string]Artists {
	var folderID, modifiedSince int64
	indexes, _ := subsonic.GetIndexes(folderID, modifiedSince)
	indexesByAlpha := make(map[string]Artists)

	for _, index := range indexes {
		indexesByAlpha[index.Name] = index.Artist
	}

	return indexesByAlpha
}

func promptForInput() Input {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("=>[?]  ")
	text, _ := reader.ReadString('\n')
	index := strings.ToUpper(strings.TrimSpace(text))

	if index == "X" || index == "Y" || index == "Z" {
		index = "X-Z"
	}

	indexType := ALPHA
	var idx int
	var err error
	if idx, err = strconv.Atoi(index); err == nil {
		indexType = NUMBER
	}

	return Input{
		Index:     index,
		IndexType: indexType,
		Idx:       idx,
	}
}

func main() {
	config := LoadConfig()
	subsonic := getClient(config)
	indexes := loadIndexes(subsonic)
	printer := NewPrinter(indexes, subsonic)

	fmt.Println("Start with a letter or number...")
	var inputs []Input
	for input := promptForInput(); ; input = promptForInput() {
		if input.IndexType == ALPHA {
			inputs = []Input{input}
		} else {
			inputs = append(inputs, input)
		}
		printer.PrintIndex(inputs)
	}
}
