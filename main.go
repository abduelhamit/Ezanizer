package main

import (
	"context"
	"log"
	"os"

	diyanet "github.com/abduelhamit/DiyanetAwqatSalahAPI"
)

func main() {
	log.Println("Starting Ezanizer...")
	defer log.Println("Ezanizer stopped")

	ctx := context.Background()

	email, ok := os.LookupEnv("DIYANET_EMAIL")
	if !ok || email == "" {
		log.Fatalln("DIYANET_EMAIL is not set")
	}

	password, ok := os.LookupEnv("DIYANET_PASSWORD")
	if !ok || password == "" {
		log.Fatalln("DIYANET_PASSWORD is not set")
	}

	countryCode, ok := os.LookupEnv("COUNTRY_CODE")
	if !ok || countryCode == "" {
		log.Fatalln("COUNTRY_CODE is not set")
	}

	stateCode, ok := os.LookupEnv("STATE_CODE")
	if !ok || stateCode == "" {
		log.Fatalln("STATE_CODE is not set")
	}

	cityCode, ok := os.LookupEnv("CITY_CODE")
	if !ok || cityCode == "" {
		log.Fatalln("CITY_CODE is not set")
	}

	config := &diyanet.Config{
		Email:    email,
		Password: password,
	}

	client := config.NewClient(ctx)
	country, err := client.GetCountry(countryCode)
	if err != nil {
		log.Fatalln("Failed to get country:", err)
	}

	log.Printf("Loaded country: %+v\n", country)

	state, err := country.GetState(stateCode)
	if err != nil {
		log.Fatalln("Failed to get state:", err)
	}

	log.Printf("Loaded state: %+v\n", state)

	city, err := state.GetCity(cityCode)
	if err != nil {
		log.Fatalln("Failed to get city:", err)
	}

	log.Printf("Loaded city: %+v\n", city)

	cityDetail, err := city.GetCityDetail()
	if err != nil {
		log.Fatalln("Failed to get city detail:", err)
	}

	log.Printf("City detail: %+v\n", cityDetail)
}
