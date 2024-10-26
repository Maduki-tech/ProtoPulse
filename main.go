package main

import (
	"flag"
	"github.com/Maduki-tech/ProtoPulse/http_tester"
)

func main() {
	duration := flag.Int("duration", 30, "Duration of the test in seconds")
	concurrency := flag.Int("concurrency", 10, "Number of concurrent requests")
	port := flag.String("port", ":8080", "Port to listen on")

	flag.Parse()

	runHTTTP(duration, concurrency, port)

	// RUN GRPC

	// Save the results
}

func runHTTTP(duration *int, concurrency *int, port *string) {
	// RUN HTTP server
	httpServer := http_tester.New(*port)
	go httpServer.Run()

	// RUN HTTP client
	httpClient := http_tester.NewClient(*concurrency, *duration, *port)
	httpClient.Run()
}
