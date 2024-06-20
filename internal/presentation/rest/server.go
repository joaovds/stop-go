package rest

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/joaovds/stop-go/internal/infra/sqlite"
	"github.com/joaovds/stop-go/internal/presentation/rest/routes"
)

type Rest struct {
	mainMux *http.ServeMux
	MuxV1   *http.ServeMux
	MuxV1WS *http.ServeMux
	Port    string
	wg      *sync.WaitGroup
	db      *sqlite.DB
}

func NewRest(wg *sync.WaitGroup, db *sqlite.DB) *Rest {
	mainMux := http.NewServeMux()
	muxV1 := http.NewServeMux()
	muxV1WS := http.NewServeMux()

	mainMux.Handle("/api/v1/", http.StripPrefix("/api/v1", muxV1))
	mainMux.Handle("/api/v1/ws", http.StripPrefix("/api/v1/ws", muxV1WS))

	rest := &Rest{
		mainMux: mainMux,
		MuxV1:   muxV1,
		MuxV1WS: muxV1WS,
		Port:    "8080",
		wg:      wg,
		db:      db,
	}

	rest.SetupRoutes()

	return rest
}

// ----- ... -----

func (r *Rest) SetupRoutes() {
	routes.NewRoomRoutes(r.MuxV1, r.db).RegisterRoutes()
	routes.NewPlayerRoutes(r.MuxV1, r.db).RegisterRoutes()
}

// ----- ... -----

func (r *Rest) Run() {
	defer r.wg.Done()

	log.Println(fmt.Sprintf("Starting server on port %s", r.Port))
	err := http.ListenAndServe(":"+r.Port, r.mainMux)
	if err != nil {
		log.Fatalf("Error starting webserver on port %s: %v", r.Port, err)
	}
}
