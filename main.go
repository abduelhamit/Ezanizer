package main

import (
	"log"
	"os"
)

func main() {
	log.Println("Starting Ezanizer...")
	defer log.Println("Ezanizer stopped")

	email, ok := os.LookupEnv("DIYANET_EMAIL")
	if !ok || email == "" {
		log.Fatalln("DIYANET_EMAIL is not set")
	}

	password, ok := os.LookupEnv("DIYANET_PASSWORD")
	if !ok || password == "" {
		log.Fatalln("DIYANET_PASSWORD is not set")
	}
}
