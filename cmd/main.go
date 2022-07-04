package main

import (
	"log"
	"os"

	// Blank-import he function package so the init() runs
	_ "github.com/edg956/pixer/functions"
	"github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
)

func main() {
	port := "8080"
	if envPort := os.Getenv("PORT"); envPort != "" {
		port = envPort
	}
	if err := funcframework.Start(port); err != nil {
		log.Fatalf("funcframework.Start: %V\n", err)
	}
}
