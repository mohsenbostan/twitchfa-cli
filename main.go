package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/mohsenbostan/twitchfa-cli/acii"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

var hasError bool = false

// Stream structure
type Stream struct {
	Streamer string `json:"streamer"`
	Title    string `json:"title"`
	Url      string `json:"url"`
	Viewers  int    `json:"viewers"`
	Game     string `json:"game"`
}

func (s *Stream) Print() {
	fmt.Printf("\n%v [ %v ] : %v \n %v\n%v\n\n", s.Streamer, s.Viewers, s.Game, strings.TrimSpace(s.Title), s.Url)
}

// A global variable to save all streams
var streams []Stream

// Get All Streams From Twitchfa API
func GetAllStreams() error {
	req, err := http.Get("https://api.twitchfa.ir/streamer")
	defer req.Body.Close()
	reqBody, err := ioutil.ReadAll(req.Body)
	err = json.Unmarshal(reqBody, &streams)

	return errors.Unwrap(err)
}

// Search For A Specific Stream
func SearchForStream(streamer string) (*Stream, bool) {
	for _, stream := range streams {
		if strings.ToLower(stream.Streamer) == strings.ToLower(streamer) {
			return &stream, true
		}
	}

	return nil, false
}

func main() {
	// Get all streams
	err := GetAllStreams()
	if err != nil {
		hasError = true
		acii.Sadge()
		panic(err)
	}

	// Check if user provided a specific streamer
	if len(os.Args) > 1 && os.Args[1] != "" {
		// Search for the specific streamer
		stream, found := SearchForStream(os.Args[1])
		if found {
			// Print the stream if found
			stream.Print()
		} else {
			// Print error if the stream not found
			hasError = true
			acii.Sadge()
			fmt.Printf("[Sadge] Stream not found!\n\n")
		}

		fmt.Println("===============**********==============")
	} else {
		// Print all streams if user didn't provided a specific streamer
		for _, stream := range streams {
			stream.Print()
			fmt.Println("===============**********==============")
		}
	}

	fmt.Println("       Thanks for using Twitchfa")
	fmt.Println("          https://twitchfa.ir")
	fmt.Printf("===============**********==============")
	if !hasError {
		fmt.Println()
		acii.Laugh()
	}
}
