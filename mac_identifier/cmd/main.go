package main

import (
	"log"
	"mac_identifier/internal/adapters/vendormap"
	"mac_identifier/internal/api"
	"mac_identifier/internal/applications/mac"
	"os"
)

func main() {
	vendorMapPath := os.Getenv("VENDOR_MAP_PATH")
	if vendorMapPath == "" {
		vendorMapPath = "./mac_vendor_map"
	}

	repo, err := vendormap.NewVendorMapRepo(vendorMapPath)
	if err != nil {
		log.Fatalf("failed to load vendor map: %v", err)
	}

	macService := mac.NewService(repo)
	macHandler := api.NewMacHandler(macService)

	app := api.NewRouter(macHandler)

	log.Println("Starting server on :8000")
	if err := app.Listen(":8000"); err != nil {
		log.Fatal(err)
	}
}
