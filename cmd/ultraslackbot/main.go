package main

import (
	"math/rand"
	"time"

	"github.com/disiqueira/ultraslackbot/internal/cli"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	cli.Execute()
}
