package main

import (
	"flag"
	"log"

	"github.com/CcyBborg/golik-blog/internal/api/handler/postshandler"

	"github.com/BurntSushi/toml"
	"github.com/CcyBborg/golik-blog/internal/app/apiserver"
	"github.com/CcyBborg/golik-blog/internal/services/posts"
	"github.com/CcyBborg/golik-blog/internal/store"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "Path to toml server config file")
}

func main() {
	flag.Parse()

	config := apiserver.NewConfig()
	if _, err := toml.DecodeFile(configPath, config); err != nil {
		log.Fatal(err)
	}

	// Initialize store
	st := store.New(config.Store)
	if err := st.Open(); err != nil {
		log.Fatal(err)
	}
	defer st.Close()

	// Initialize services
	postsService := posts.New(st)

	// Initialize handlers
	postsHandler := postshandler.New(postsService)

	s := apiserver.New(config)

	// Register HTTP-handlers
	s.RegisterHTTPHandler("/api/1/posts", postsHandler.Handle)

	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
