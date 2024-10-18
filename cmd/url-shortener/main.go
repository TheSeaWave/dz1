package main

import (
	"example/url-shortener/internal/config"
	"fmt"
)

func main() {
	// TODO: init config: cleanenv
	cfg := config.MustLoad()

	fmt.Println(cfg)

	// TODO: init logger: slog

	// TODO: init sotage: sqlite

	// TODO: init router: chi, "chi-render"

	// TODO: run server:

}
