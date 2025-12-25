package main

import (
	"context"
	"log"
	"os"
)

func main() {
	log.Println("Starting Ezanizer...")
	defer log.Println("Ezanizer stopped")

	ctx := context.TODO()

	email, ok := os.LookupEnv("DIYANET_EMAIL")
	if !ok || email == "" {
		log.Fatalln("DIYANET_EMAIL is not set")
	}

	password, ok := os.LookupEnv("DIYANET_PASSWORD")
	if !ok || password == "" {
		log.Fatalln("DIYANET_PASSWORD is not set")
	}

	config := &Config{
		Email:    email,
		Password: password,
	}

	_, err := config.Token(ctx)
	if err != nil {
		log.Fatalf("Failed to get token: %v", err)
	}
}
