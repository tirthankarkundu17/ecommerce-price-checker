package main

import "github.com/tirthankarkundu17/ecommerce-price-checker/handler"

func main() {

	var server = handler.Server{}
	server.Initialize()

	// Start server
	go func() {
		server.Run("localhost", "8000")
	}()
	// Block main thread to keep server running
	select {}
}
