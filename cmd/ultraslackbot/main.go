package main

import (
	"github.com/disiqueira/ultraslackbot/internal/cli"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	cli.Execute()
}

