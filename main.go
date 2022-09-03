package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/LC-Gub/Go_for_begineer/api"
	"github.com/fatih/color"
)

//run: "go run ." << for development
//build: "go build ." << ready to ship code OR "go build -o yourfileName.exe ."

// cross compile
// GOOS=darwin GOARCH=arm64 go build -o boardgameatlas-darwin-arm64 .
// GOOS=windows GOARCH=amd64 go build -o boardgameatlas-windows-amd64 .

func main() {
	//pointer to query

	// bga --query "ticket to ride" --clientID abcd123 --skip 10 --limit 5
	//Defining command line arguments
	query := flag.String("query", "", "Boardgame name to search.")
	clientId := flag.String("clientId", "", "Boardgame Atlas client_id.")
	skip := flag.Uint("skip", 0, "Skip number of results provided.")
	limit := flag.Uint("limit", 0, "Limit number of results provided.")
	timeout := flag.Uint("timeout", 10, "Timeout.")

	//Parsing command line, must use this to set the values
	flag.Parse()

	//fmt.Printf("query=%s, clientId=%s, limit=%d, skip=%d\n", *query, *clientId, *limit, *skip)

	//Chcek if --query and --clientId are set.
	if isNull(*query) {
		log.Fatalln("\nPlease use --query to set the boardgame name to search.")
	}

	if isNull(*clientId) {
		log.Fatalln("\nPlease use --clientId to set the boardgame Atlas client_id")
	}

	//Instantiate a BoardGameAtlas struct
	bga := api.New(*clientId)

	//Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(*timeout*uint(time.Second)))
	defer cancel() //Delay execution, defer is a function

	time.Sleep(5 * time.Second)

	//Invocation
	result, err := bga.Search(ctx, *query, *limit, *skip)
	if nil != err {
		log.Fatalf("Cannot search for boardgame: %v", err)
	}

	//colors
	boldCyan := color.New(color.Bold).Add(color.FgHiCyan).SprintFunc()

	for _, g := range result.Games {
		fmt.Printf("%s: %s\n", boldCyan("Name"), g.Name)
		fmt.Printf("%s: %s\n", boldCyan("Description"), g.Description)
		fmt.Printf("%s: %s\n\n", boldCyan("Url"), g.Url)
	}

}

func isNull(s string) bool {
	//Remove all white spaces, then count the length. If it's less than or equals 0, then it's empty.
	return len(strings.TrimSpace(s)) <= 0
}
