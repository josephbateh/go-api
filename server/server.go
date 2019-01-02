package server

import (
	"net/http"
	"os"
	"sync"

	"github.com/josephbateh/go-api/api"
	"github.com/josephbateh/go-api/log"
)

// Start begins execution for the server
func Start() {
	http.HandleFunc("/teams", api.Teams)
	var wg sync.WaitGroup
	wg.Add(1)
	if os.Getenv("PORT") != "" {
		go http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	} else {
		go http.ListenAndServe(":8080", nil)
	}
	log.Info("Server started")
	wg.Wait()
}
