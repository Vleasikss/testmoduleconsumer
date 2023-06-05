package main

import (
	"fmt"

	"github.com/Vleasikss/testmoduleproducer/v5/sender"
)

func main() {

	fmt.Println("Hello from Consumer")
	sender.SendMessage("Hello from Producer")
	sender.SendMessageNTimes("hello", 5)
	sender.SayBye()

}
