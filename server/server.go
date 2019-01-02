package server

import (
	"net/http"
	"sync"

	"github.com/josephbateh/go-api/api"
	"github.com/josephbateh/go-api/log"
)

// Start begins execution for the server
func Start() {
	http.HandleFunc("/teams", api.Teams)
	var wg sync.WaitGroup
	wg.Add(1)
	go http.ListenAndServe(":8080", nil)
	log.Info("Server started")
	wg.Wait()
}
