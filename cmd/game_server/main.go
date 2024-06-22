package main

import (
	"sync"

	"github.com/joaovds/stop-go/internal/presentation/rest"
)

func main() {
	wg := &sync.WaitGroup{}

	rest_server := rest.NewRest(wg)

	wg.Add(1)
	go rest_server.Run()

	wg.Wait()
}
