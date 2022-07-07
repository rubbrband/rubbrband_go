package main

import (
	"fmt"

	"github.com/rubbrband/rubbrband_go"
)

func main() {
	rubbrband_go.SetApiKey("3b444016-4d89-4757-ac2c-02eb7beb7fff")

	// Store a string
	result1 := rubbrband_go.Replay("foo_2")
	fmt.Println("Replay Result,", result1)
}
