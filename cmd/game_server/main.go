package main

import (
	"sync"

	"github.com/joaovds/stop-go/internal/infra/sqlite"
	"github.com/joaovds/stop-go/internal/presentation/rest"
)

func main() {
	wg := &sync.WaitGroup{}
	db := sqlite.SetupDB()
	defer db.Close()

	rest_server := rest.NewRest(wg, db)

	wg.Add(1)
	go rest_server.Run()

	wg.Wait()
}
