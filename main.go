package main

import (
	"embed"

	"github.com/HoneySinghDev/go-echo-rest-api-template/cmd"
)

//go:embed static
var static embed.FS

func main() {
	cmd.App(static)
}
