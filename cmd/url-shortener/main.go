package main

import (
	"fmt"
	"urlShortener/internal/config"
)

func main() {

	// todo init config: cleanenv
	cfg := config.MustLoad()
	fmt.Printf("%+v\n", cfg)

	// todo init logger: log/slog

	// todo init db: sqlite

	// todo init router: chi, render

	// todo run server

}
