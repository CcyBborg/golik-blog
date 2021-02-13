package main

import (
	"flag"
	"log"

	"github.com/CcyBborg/golik-blog/internal/api/handler/categorieshandler"
	"github.com/CcyBborg/golik-blog/internal/api/handler/commentshandler"
	"github.com/CcyBborg/golik-blog/internal/api/handler/mypostshandler"
	"github.com/CcyBborg/golik-blog/internal/api/handler/posthandler"
	"github.com/CcyBborg/golik-blog/internal/api/handler/postshandler"

	"github.com/BurntSushi/toml"
	"github.com/CcyBborg/golik-blog/internal/app/apiserver"
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

	// Initialize handlers
	postsHandler := postshandler.New(st)
	postHandler := posthandler.New(st)
	mypostsHandler := mypostshandler.New(st)
	commentsHandler := commentshandler.New(st)
	categoriesHandler := categorieshandler.New(st)

	s := apiserver.New(config)

	// Register HTTP-handlers
	s.RegisterHTTPHandler("/posts", postsHandler.Handle)
	s.RegisterHTTPHandler("/posts/{postID}", postHandler.Handle)
	s.RegisterHTTPHandler("/myposts", mypostsHandler.Handle)
	s.RegisterHTTPHandler("/posts/{postID}/comments", commentsHandler.Handle)
	s.RegisterHTTPHandler("/categories", categoriesHandler.Handle)

	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
