package main

import (
	"log"
	"net"
	"os"

	"shortlink/api"
	"shortlink/server"
	"shortlink/shortcut"

	"google.golang.org/grpc"
)

// getenv gets an environment variable value given a key or returns fallback and logs decision
func getenv(key string, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		log.Printf("Environment variable \"%s\" was not set, using fallback: \"%s\"", key, fallback)
		return fallback
	}
	log.Printf("Environment variable \"%s\" was set to: \"%s\" ", key, value)
	return value
}

func dialService(addr string) *grpc.ClientConn {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithBlock())
	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		log.Fatalf("Failed to connect to Service at %s: %v", addr, err)
	}
	return conn
}

func main() {
	// Connect to ShortcutService
	log.Println("Trying to connect to ShortcutService...")
	shortcutAddr := getenv("SHORTCUT_ADDR", "shortcut_svc:9092")
	shortcutClient := shortcut.NewShortcutServiceClient(dialService(shortcutAddr))
	log.Println("Done!")

	// Server listener
	log.Println("Trying to serve APIService...")
	port := getenv("PORT", ":9090")
	lis, err := net.Listen("tcp", "0.0.0.0"+port)
	if err != nil {
		log.Fatalf("Failed to listen on port %s (%v)", port, err)
	}

	// Serve APIService server
	s := grpc.NewServer()
	api.RegisterAPIServiceServer(
		s,
		&server.Server{ShortcutClient: shortcutClient},
	)
	log.Printf("Serving on port %s ...", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve on port %s (%v)", port, err)
	}
}
