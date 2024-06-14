package main

import (
	"fmt"
	"github.com/simonkarman/pastebin-client-go"
	"net/url"
	"time"
)

func main() {
	// Create Client
	host, _ := url.Parse("https://pastebin.com/")
	devKey := "<paste-your-dev-key-here>"
	userKey := "<paste-your-user-key-here>"
	client := pastebin.New(*host, devKey, userKey)

	// Create Paste
	pasteKey, err := client.CreatePaste("{\n  \"message\": \"Hello, world!\"\n}")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created: %s\n", pasteKey)

	paste, err := client.GetPaste(pasteKey)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Got: %s\n", paste)

	// sleep for 20 seconds
	fmt.Println("Deleting in 1 second...")
	time.Sleep(1 * time.Second)

	// Delete Paste
	err = client.DeletePaste(pasteKey)
	if err != nil {
		panic(err)
	}
	fmt.Println("Deleted.")
}
