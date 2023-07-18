package main

import (
	"fmt"
	"petProject/internal/config"
)

func main() {
	cfg := config.MustLoad()

	fmt.Println(cfg)

	// TODO: inti logger: slog

	// TODO: inti storage: sqlite

	// TODO: inti router: chi, "chi reder"

	// TODO: run server
}
